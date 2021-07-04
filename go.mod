module github.com/yoshida-lab/avalon

go 1.16

replace github.com/yoshida-lab/avalon => ../avalon

require (
	github.com/gin-gonic/examples v0.0.0-20210522065734-9fd0db1d6a7c
	github.com/gin-gonic/gin v1.7.2
	github.com/golang/protobuf v1.4.3
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	google.golang.org/grpc v1.39.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
	google.golang.org/protobuf v1.25.0
)
