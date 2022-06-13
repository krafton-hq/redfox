module github.com/krafton-hq/red-fox/server

go 1.18

require (
	github.com/spf13/cobra v1.4.0
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/gofiber/utils v0.1.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/rs/cors v1.7.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.37.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210126160654-44e461bb6506 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	nhooyr.io/websocket v1.8.6 // indirect
)

require (
	github.com/gofiber/adaptor/v2 v2.1.24
	github.com/gofiber/fiber/v2 v2.34.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/improbable-eng/grpc-web v0.15.0
	github.com/krafton-hq/golib v0.0.1
	github.com/krafton-hq/red-fox/sdk v0.0.0
	go.uber.org/zap v1.21.0
	google.golang.org/grpc v1.47.0
	sigs.k8s.io/yaml v1.3.0
)

replace github.com/krafton-hq/red-fox/sdk => ../sdk
