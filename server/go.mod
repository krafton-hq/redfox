module github.com/krafton-hq/red-fox/server

go 1.18

require (
	github.com/spf13/cobra v1.4.0
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/gofiber/utils v0.1.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/huandu/xstrings v1.3.1 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mitchellh/copystructure v1.1.1 // indirect
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/rs/cors v1.7.0 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.37.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/goleak v1.1.12 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20220427172511-eb4f295cb31f // indirect
	golang.org/x/exp v0.0.0-20220303212507-bbda1eaf7a17 // indirect
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220324131243-acbaeb5b85eb // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/klog/v2 v2.60.1 // indirect
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9 // indirect
	nhooyr.io/websocket v1.8.6 // indirect
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
)

require (
	github.com/Masterminds/sprig/v3 v3.2.2
	github.com/doug-martin/goqu/v9 v9.18.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gofiber/adaptor/v2 v2.1.24
	github.com/gofiber/fiber/v2 v2.34.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/improbable-eng/grpc-web v0.15.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/krafton-hq/golib v0.0.2
	github.com/krafton-hq/red-fox/apis v0.0.0
	github.com/samber/lo v1.21.0
	github.com/xo/dburl v0.11.0
	go.uber.org/zap v1.21.0
	google.golang.org/grpc v1.47.0
	k8s.io/apimachinery v0.24.2
	sigs.k8s.io/external-dns v0.12.0
	sigs.k8s.io/yaml v1.3.0
)

replace github.com/krafton-hq/red-fox/apis => ../apis
