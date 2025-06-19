## go web的基础框架

├── cmd/                     # 可执行程序入口
│   ├── server/              # 服务主程序（可自定义名称如 "api"）
│   │    └── main.go          # 入口文件
├── ├── server2/              # 服务主程序（可自定义名称如 "api"）
│   │    └── main.go          # 入口文件
│   
├── api/                     # API 协议定义（替代原 protos）
│   ├── hello.v1.proto       # gRPC 协议文件
│   └── swagger.json         # OpenAPI 文档
├── configs/                 # 配置文件
│   └── app.yaml             # 主配置文件
├── internal/                # 私有应用代码（禁止外部导入）
│   ├── server/              # 服务核心实现
│   │   ├── handler/         # 请求处理器
│   │   ├── middleware/      # 中间件
│   │   └── service/         # 业务逻辑层
│   ├── server2/            # 服务核心实现
│   │   ├── handler/         # 请求处理器
│   │   ├── middleware/      # 中间件
│   │   └── service/         # 业务逻辑层
│   └── dao/                 # 数据访问层
├── pkg/                     # 公共库代码（允许外部导入）
│   ├── utils/               # 工具函数
│   └── constants/           # 全局常量
├── routes/                  # HTTP 路由管理
│   ├── v1/                  # API 版本 v1
│   └── register.go          # 路由注册逻辑
├── go.mod                   # 模块定义（必须在根目录）
├── go.sum
├── Makefile                 # 构建脚本
└── README.md

```
1. 启动服务: go run .\cmd\server\main.go  || go build -o myserver.exe cmd/server/main.go && myserver.exe
```