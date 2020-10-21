# book_store

- go-zero example
- 遵循go-zero规范

**文件目录**

- gateway

```
api
├── bookstore.api                  // api定义
├── bookstore.go                   // main入口定义
├── etc
│   └── bookstore-api.yaml         // 配置文件
└── internal
    ├── config
    │   └── config.go              // 定义配置
    ├── handler
    │   ├── addhandler.go          // 实现addHandler
    │   ├── checkhandler.go        // 实现checkHandler
    │   └── routes.go              // 定义路由处理
    ├── logic
    │   ├── addlogic.go            // 实现AddLogic
    │   └── checklogic.go          // 实现CheckLogic
    ├── svc
    │   └── servicecontext.go      // 定义ServiceContext
    └── types
        └── types.go               // 定义请求、返回结构体
```

- rpc

```
rpc/add
├── add.go                      // rpc服务main函数
├── add.proto                   // rpc接口定义
├── adder
│   ├── adder.go                // 提供了外部调用方法，无需修改
│   ├── adder_mock.go           // mock方法，测试用
│   └── types.go                // request/response结构体定义
├── etc
│   └── add.yaml                // 配置文件
├── internal
│   ├── config
│   │   └── config.go           // 配置定义
│   ├── logic
│   │   └── addlogic.go         // add业务逻辑在这里实现
│   ├── server
│   │   └── adderserver.go      // 调用入口, 不需要修改
│   └── svc
│       └── servicecontext.go   // 定义ServiceContext，传递依赖
└── pb
    └── add.pb.go

rpc/check
├── check.go                    // rpc服务main函数
├── check.proto                 // rpc接口定义
├── checker
│   ├── checker.go              // 提供了外部调用方法，无需修改
│   ├── checker_mock.go         // mock方法，测试用
│   └── types.go                // request/response结构体定义
├── etc
│   └── check.yaml              // 配置文件
├── internal
│   ├── config
│   │   └── config.go           // 配置定义
│   ├── logic
│   │   └── checklogic.go       // check业务逻辑在这里实现
│   ├── server
│   │   └── checkerserver.go    // 调用入口, 不需要修改
│   └── svc
│       └── servicecontext.go   // 定义ServiceContext，传递依赖
└── pb
    └── check.pb.go
```
