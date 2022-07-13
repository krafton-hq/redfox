package fox_grpc

import (
	"crypto/x509"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/krafton-hq/red-fox/apis/api_resources"
	"github.com/krafton-hq/red-fox/apis/app_lifecycle"
	"github.com/krafton-hq/red-fox/apis/crds"
	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/namespaces"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type RedFoxClientConfig struct {
	GrpcEndpoint string

	DialOptions        []grpc.DialOption
	ClientInterceptors []grpc.UnaryClientInterceptor

	WithTls bool

	// 0 means disable
	TimeoutSeconds int
	// 0 means disable
	MaxRetry int
}

func DefaultConfig() *RedFoxClientConfig {
	grpcEndpoint, withTls, _ := ResolveGrpcEndpointFromEnv()

	return &RedFoxClientConfig{
		//domain and port should be define in endpoint.
		//example: localhost:7080, myapi.mydomain.com:443
		GrpcEndpoint:   grpcEndpoint,
		WithTls:        withTls,
		TimeoutSeconds: 20,
		MaxRetry:       3,
	}
}

func ResolveGrpcEndpointFromEnv() (grpcEndpoint string, withTls bool, err error) {
	if grpcAddr, exists := os.LookupEnv("REDFOX_GRPC_ADDR"); exists {
		grpcEndpoint = grpcAddr
	}
	if eWithTls, exists := os.LookupEnv("REDFOX_GRPC_WITH_TLS"); exists {
		withTls, _ = strconv.ParseBool(eWithTls)
		return
	}

	if restAddr, exists := os.LookupEnv("REDFOX_ADDR"); exists {
		grpcEndpoint, withTls, err = ConvertRestToGrpcEndpoint(restAddr)
		if err != nil {
			return "", false, err
		}
		return
	}
	return
}

// ConvertRestToGrpcEndpoint input: https://fox.foo.example.com/api/v1 -> fox-grpc.foo.example.com:443, withTls=true
func ConvertRestToGrpcEndpoint(restEndpoint string) (grpcEndpoint string, withTls bool, err error) {
	r, _ := regexp.Compile(`(http(s)?:\/\/)([a-z0-9\w-:]+\.*)+`)
	if !r.MatchString(restEndpoint) {
		err = fmt.Errorf("requested RestEndpoint Does Not Match Regex")
		return
	}

	restUrl, err := url.Parse(restEndpoint)
	if err != nil {
		return
	}

	if strings.EqualFold(restUrl.Scheme, "http") {
		withTls = false
	} else if strings.EqualFold(restUrl.Scheme, "https") {
		withTls = true
	} else {
		err = fmt.Errorf("requested RestEndpoint Protocol Does Not Match HTTP Or HTTPS: %s", restUrl.Scheme)
		return
	}

	subDomains := strings.SplitN(restUrl.Hostname(), ".", 2)
	if len(subDomains) <= 1 {
		err = fmt.Errorf("requested RestEndpoint Host Does Not Have SubDomains")
		return
	}
	grpcHost := fmt.Sprintf("%s-grpc.%s", subDomains[0], subDomains[1])

	var grpcPort string
	if restUrl.Port() == "" {
		if withTls {
			grpcPort = "443"
		} else {
			grpcPort = "80"
		}
	} else {
		grpcPort = restUrl.Port()
	}

	grpcEndpoint = fmt.Sprintf("%s:%s", grpcHost, grpcPort)
	return
}

func newClient(config *RedFoxClientConfig) (*RedFoxClient, error) {
	if config == nil {
		return nil, fmt.Errorf("fox client config is nil")
	}

	retryInt := grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithPerRetryTimeout(time.Duration(config.TimeoutSeconds)*time.Second),
		grpc_retry.WithMax(uint(config.MaxRetry)))

	options := append(config.DialOptions, grpc.WithChainUnaryInterceptor(append(config.ClientInterceptors, retryInt)...))
	if config.WithTls {
		pool, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		options = append(options, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(pool, "")))
	} else {
		options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(config.GrpcEndpoint, options...)
	if err != nil {
		return nil, err
	}

	fClient := &RedFoxClient{
		NatIpServerClient:                    documents.NewNatIpServerClient(conn),
		EndpointServerClient:                 documents.NewEndpointServerClient(conn),
		CustomDocumentServerClient:           documents.NewCustomDocumentServerClient(conn),
		NamespaceServerClient:                namespaces.NewNamespaceServerClient(conn),
		CustomDocumentDefinitionServerClient: crds.NewCustomDocumentDefinitionServerClient(conn),
		ApiResourcesServerClient:             api_resources.NewApiResourcesServerClient(conn),
		ApplicationLifecycleClient:           app_lifecycle.NewApplicationLifecycleClient(conn),
	}
	return fClient, nil
}

// NewClient is return new fox client.
// RedFoxClientConfig is required.
func NewClient(config *RedFoxClientConfig) (*RedFoxClient, error) {
	return newClient(config)
}
