package application

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	log_helper "github.com/krafton-hq/golib/log-helper"
	"github.com/krafton-hq/red-fox/apis/app_lifecycle"
	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/namespaces"
	"github.com/krafton-hq/red-fox/server/application/configs"
	"github.com/krafton-hq/red-fox/server/controllers/app_lifecycle_con"
	"github.com/krafton-hq/red-fox/server/controllers/document_con"
	"github.com/krafton-hq/red-fox/server/controllers/namespace_con"
	"github.com/krafton-hq/red-fox/server/repositories/apiobject_repository"
	"github.com/krafton-hq/red-fox/server/services/namespace_service"
	"github.com/krafton-hq/red-fox/server/services/natip_service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Application struct {
	config *configs.RedFoxConfig

	grpcServer *grpc.Server

	nsController    *namespace_con.Controller
	natIpController *document_con.NatIpController
	appController   *app_lifecycle_con.GrpcController
}

func NewApplication(config *configs.RedFoxConfig) *Application {
	return &Application{config: config}
}

func (a *Application) Init() {
	a.initInternal()

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			append(log_helper.GetUnaryServerInterceptors(),
				grpc_recovery.UnaryServerInterceptor(),
			)...,
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			append(log_helper.GetStreamServerInterceptors(),
				grpc_recovery.StreamServerInterceptor(),
			)...,
		)))
	reflection.Register(grpcServer)

	app_lifecycle.RegisterApplicationLifecycleServer(grpcServer, a.appController)
	namespaces.RegisterNamespaceServerServer(grpcServer, a.nsController)
	documents.RegisterNatIpServerServer(grpcServer, a.natIpController)

	for name := range grpcServer.GetServiceInfo() {
		zap.S().Infow("Registered gRpc Service", "name", name)
	}

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
	app_lifecycle_con.NewAppLifecycleHttp(a.appController).Register(appLifecycle)

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

func (a *Application) initInternal() error {
	rawNsRepo := apiobject_repository.NewInMemoryNamespacedRepository[*namespaces.Namespace](&namespaces.GroupVersionKind{
		Group:   "core",
		Version: "v1",
		Kind:    "Namespace",
		Enabled: true,
	})
	nsRepo := apiobject_repository.NewSimpleClusterRepository[*namespaces.Namespace](rawNsRepo)

	natIpRepo := apiobject_repository.NewInMemoryNamespacedRepository[*documents.NatIp](&namespaces.GroupVersionKind{
		Group:   "red-fox.sbx-central.io",
		Version: "v1alpha1",
		Kind:    "NatIp",
		Enabled: true,
	})

	endpointRepo := apiobject_repository.NewInMemoryNamespacedRepository[*documents.Endpoint](&namespaces.GroupVersionKind{
		Group:   "red-fox.sbx-central.io",
		Version: "v1alpha1",
		Kind:    "Endpoint",
		Enabled: true,
	})
	var nsRepos = []apiobject_repository.NamespacedRepositoryMetadata{natIpRepo, endpointRepo}

	natIpService := natip_service.NewService(natIpRepo)
	a.natIpController = document_con.NewNatIpDocController(natIpService)

	nsService := namespace_service.NewService(nsRepo, nsRepos)
	a.nsController = namespace_con.NewController(nsService)

	a.appController = app_lifecycle_con.NewAppLifecycle()
	return nil
}
