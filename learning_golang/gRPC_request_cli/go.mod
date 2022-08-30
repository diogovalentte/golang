module github.com/diogovalentte/golang/gRPC_request_cli

go 1.18

require (
	github.com/spf13/cobra v1.5.0
	github.com/spf13/pflag v1.0.5
	google.golang.org/grpc v1.48.0
)

require github.com/diogovalentte/golang/gRPC_server v0.0.0

replace github.com/diogovalentte/golang/gRPC_server v0.0.0 => ../gRPC_server

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	golang.org/x/net v0.0.0-20220812174116-3211cb980234 // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220812140447-cec7f5303424 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
