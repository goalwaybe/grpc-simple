# gRPC Simple Example (Go)


这是一个使用 Go 语言编写的简单 gRPC 示例项目，展示了如何定义 `.proto` 服务、生成 Go 代码，并实现 gRPC 服务端与客户端。

## 📁 项目结构

```
grpc-simple/
├── README.md
├── go.mod
├── go.sum
├── proto/
│   └── greet.proto # gRPC 服务定义
├── pb/ # 自动生成的 Protocol Buffer Go 代码
│   └── greet.pb.go
├── server/
│   └── main.go # gRPC 服务端
└── client/
    └── main.go # gRPC 客户端
```


## ⚙️ 前置要求

- Go 1.20+
- [Protocol Buffers Compiler (`protoc`)](https://github.com/protocolbuffers/protobuf/releases)
- Go gRPC 插件: `google.golang.org/protobuf/cmd/protoc-gen-go`
- Go gRPC Gateway 插件（可选）: `google.golang.org/grpc/cmd/protoc-gen-go-grpc`

安装插件（如未安装）：
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest