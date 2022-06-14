GO_GRPC_OUT = "./"
TS_GRPC_OUT = "./frontend/src/generated"

build-proto-go:
	rm -
	protoc --go_out=$(GO_GRPC_OUT) --go_opt=paths=source_relative --go-grpc_out=$(GO_GRPC_OUT) --go-grpc_opt=paths=source_relative \
	--proto_path=. */*
