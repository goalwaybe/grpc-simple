

## gRPC Simple Example (Go)

这是一个使用 Go 语言编写的简单 gRPC 示例项目，展示了如何定义 `.proto` 服务、生成 Go 代码，并实现 gRPC 服务端与客户端。

## 📁 项目结构

```markdown
grpc-simple/
├── README.md
├── go.mod
├── go.sum
├── proto/
│   └── greet.proto
├── pb/
│   └── greet.pb.go
├── server/
│   └── main.go
└── client/
└── main.go
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
```

确保 `$GOPATH/bin` 已加入系统 `PATH`。

## 🔧 生成 Protocol Buffer 代码

在项目根目录运行：

```bash
protoc --go_out=. --go-grpc_out=. --proto_path=./proto ./proto/hello.proto
```

这会生成 `pb/greet.pb.go` 和 `pb/greet_grpc.pb.go`（或类似路径，取决于你的 `option go_package` 设置）。

> 💡 提示：本项目已包含生成好的文件，此步骤仅用于修改 `.proto` 后重新生成。

## ▶️ 运行服务端

```bash
cd server
go run main.go
```

默认监听 `localhost:50051`。

## 📡 运行客户端

在另一个终端中：

```bash
cd client
go run main.go
```

你将看到类似输出：
```
Response from server: Hello, Alice
```

## 📦 依赖管理

使用 Go Modules：

```bash
go mod tidy
```

主要依赖：
- `google.golang.org/grpc`
- `google.golang.org/protobuf`

## 📚 参考

- [gRPC Go 官方文档](https://grpc.io/docs/languages/go/)
- [Protocol Buffers 文档](https://developers.google.com/protocol-buffers)

## 📝 许可证

本项目采用 [MIT 许可证](./LICENSE)。