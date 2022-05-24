# gRPCSum
Simple unary grpc to calculate Sum of two numbers, its simple to follow and this way its easy to understand the fundamentals of it.

## Make info
```shell
❯ make
all                            Generate Pbs and build
sum                            Generate Pbs and build for Sum project
test                           Test all 
clean                          clean all generated resources
clean_pbs                      clean pbs generated files
rebuild                        rebuild whole project .. clean and then build all
about                          Display ingo related to the platform Go version and OS
help                           show this help message
```

## Build info
```shell 
❯ make about
OS: macos 21.3.0 x86_64
Shell: bash 3.2.57(1)-release
Protoc version: libprotoc 3.19.4
Go version: go version go1.18.1 darwin/amd64
Go package: github.com/nXnUs25/gRPCSum
Openssl version: LibreSSL 2.8.3
```

## Build all
```shell
❯ make all
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto
go build -o bin/server ./server
go build -o bin/client ./client
```

## Run 
```shell
❯ bin/server
2022/05/24 22:41:22 listening on 0.0.0.0:50051
2022/05/24 22:41:32 Greetings func called a:6 b:5
```

```shell
❯ bin/client
2022/05/24 22:41:32 11
```
