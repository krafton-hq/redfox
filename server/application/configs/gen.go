package configs

import (
	"os"

	"go.uber.org/zap"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=. config.proto

func ParseStringRef(ref *StringRef) string {
	if ref.Type == RefType_Static {
		return ref.Value
	} else if ref.Type == RefType_Env {
		return os.Getenv(ref.Value)
	} else {
		zap.S().Fatalw("Unknown ref value type", "ref", ref)
		return ""
	}
}
