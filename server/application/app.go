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
	"github.com/krafton-hq/red-fox/server/pkg/database_helper"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
	"github.com/krafton-hq/red-fox/server/pkg/errors"
	"github.com/krafton-hq/red-fox/server/pkg/transactional"
	"github.com/krafton-hq/red-fox/server/repositories/apiobject_repository"
	"github.com/krafton-hq/red-fox/server/services/namespace_service"
	"github.com/krafton-hq/red-fox/server/services/natip_service"
	"github.com/krafton-hq/red-fox/server/services/service_helper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
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

func (a *Application) Init() error {
	err := a.initInternal()
	if err != nil {
		return err
	}

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

	return nil
}

func (a *Application) initInternal() error {
	var tr transactional.Transactional

	switch a.config.Database.Type {
	case configs.DatabaseType_Inmemory:
		tr = transactional.NewNoop()
	case configs.DatabaseType_Mysql:
		db, err := database_helper.NewDatabase(a.config.Database.Url, configs.ParseStringRef(a.config.Database.UsernameRef), configs.ParseStringRef(a.config.Database.PasswordRef))
		if err != nil {
			zap.S().Errorw("Create Database Connection Failed", "error", err)
			return err
		}
		tr, err = transactional.NewSqlTransactional(db, transactional.DialectMysql, nil)
		if err != nil {
			zap.S().Errorw("Create Database Transaction Helper Failed", "error", err)
			return err
		}
	default:
		return errors.NewErrorf("Unknown Database Type: %s", a.config.Database.Type.String())
	}

	var nsRepo apiobject_repository.ClusterRepository[*namespaces.Namespace]
	var natIpRepo apiobject_repository.NamespacedRepository[*documents.NatIp]
	var endpointRepo apiobject_repository.NamespacedRepository[*documents.Endpoint]

	switch a.config.Database.Type {
	case configs.DatabaseType_Inmemory:
		nsRepo = apiobject_repository.NewInMemoryClusterRepository[*namespaces.Namespace](domain_helper.NamespaceGvk, apiobject_repository.DefaultSystemNamespace)
		natIpRepo = apiobject_repository.NewGenericNamespacedRepository[*documents.NatIp](domain_helper.NatIpGvk, apiobject_repository.NewInmemoryClusterRepositoryFactory[*documents.NatIp]())
		endpointRepo = apiobject_repository.NewGenericNamespacedRepository[*documents.Endpoint](domain_helper.EndpointGvk, apiobject_repository.NewInmemoryClusterRepositoryFactory[*documents.Endpoint]())
		break
	case configs.DatabaseType_Mysql:
		nsRepo = apiobject_repository.NewMysqlClusterRepository[*namespaces.Namespace](domain_helper.NamespaceGvk, apiobject_repository.DefaultSystemNamespace, domain_helper.NewNamespaceFactory(), tr)
		natIpRepo = apiobject_repository.NewGenericNamespacedRepository[*documents.NatIp](domain_helper.NatIpGvk, apiobject_repository.NewMysqlClusterRepositoryFactory[*documents.NatIp](domain_helper.NewNatIpFactory(), tr))
		endpointRepo = apiobject_repository.NewGenericNamespacedRepository[*documents.Endpoint](domain_helper.EndpointGvk, apiobject_repository.NewMysqlClusterRepositoryFactory[*documents.Endpoint](domain_helper.NewEndpointFactory(), tr))
		break
	default:
		return errors.NewErrorf("Unknown Database Type: %s", a.config.Database.Type.String())
	}

	var namespacedRepos = []apiobject_repository.NamespacedRepositoryMetadata{natIpRepo, endpointRepo}
	var natIpService service_helper.NamespacedService[*documents.NatIp] = natip_service.NewService(natIpRepo)
	var nsService service_helper.ClusterService[*namespaces.Namespace] = namespace_service.NewService(nsRepo, namespacedRepos)

	nsService = service_helper.NewTransactionalClusterService[*namespaces.Namespace](nsService, tr)
	natIpService = service_helper.NewTransactionalNamespacedService[*documents.NatIp](natIpService, tr)

	a.nsController = namespace_con.NewController(nsService)
	a.natIpController = document_con.NewNatIpDocController(natIpService)
	a.appController = app_lifecycle_con.NewAppLifecycle()
	return nil
}
