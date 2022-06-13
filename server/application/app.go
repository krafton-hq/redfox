package application

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	log_helper "github.com/krafton-hq/golib/log-helper"
	"github.com/krafton-hq/red-fox/sdk/app_life"
	"github.com/krafton-hq/red-fox/server/application/configs"
	"github.com/krafton-hq/red-fox/server/controllers/app_lifecycle_con"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Application struct {
	config *configs.RedFoxConfig
}

func NewApplication(config *configs.RedFoxConfig) *Application {
	return &Application{config: config}
}

func GetStreamServerInterceptors() []grpc.StreamServerInterceptor {
	list := []grpc.StreamServerInterceptor{
		grpc_zap.StreamServerInterceptor(zap.L()),
	}

	if ce := zap.L().Check(zap.DebugLevel, "Lorem Ipsum"); ce != nil {
		list = append(list, grpc_zap.PayloadStreamServerInterceptor(zap.L(), func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
			return true
		}))
	}
	return list
}

func (a *Application) Init() {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(log_helper.GetUnaryServerInterceptors()...)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(GetStreamServerInterceptors()...)))
	reflection.Register(grpcServer)

	appLifecycleController := app_lifecycle_con.NewAppLifecycle()
	app_life.RegisterApplicationLifecycleServer(grpcServer, appLifecycleController)

	wrappedGrpc := grpcweb.WrapServer(grpcServer)
	grpcWebHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if wrappedGrpc.IsGrpcWebRequest(req) {
			if strings.HasPrefix(req.URL.Path, "/grpc_web") {
				req.URL.Path = strings.Replace(req.URL.Path, "/grpc_web", "", 1)
			}
			wrappedGrpc.ServeHTTP(res, req)
			return
		}
		zap.S().Info("Unknown request received")
		res.WriteHeader(http.StatusNotFound)
	})
	httpServer := fiber.New()
	httpServer.Use(logger.New())
	httpServer.Group("/grpc_web", adaptor.HTTPHandlerFunc(grpcWebHandler))

	appLifecycle := httpServer.Group("/api/v1/app")
	app_lifecycle_con.NewAppLifecycleHttp(appLifecycleController).Register(appLifecycle)

	go func(port int32) {
		listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		fmt.Printf("Grpc Server listen http://localhost:%d\n", port)
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}(a.config.Listeners.GrpcPort)

	go func(port int32) {
		fmt.Printf("Grpc Web Server listen http://localhost:%d\n", port)
		err := httpServer.Listen(fmt.Sprintf(":%d", port))
		if err != nil {
			log.Fatalf("failed to start grpc-web server: %v", err)
		}
	}(a.config.Listeners.RestPort)

}
