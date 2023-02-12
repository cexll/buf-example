# buf.build

buf 之前的 proto 管理槽点

1. makefile win支持不友好
2. protoc/protoc-gen-xxx版本不一致
3. proto文件依赖难以管理

buf 优势

1. 跨平台 支持 linux mac win
2. plugin 中心化管理
3. 编译速度更快
4. proto依赖管理 类似 go mod
5. 可配置lint break change检查

buf 文件/目录介绍

```
.
├── README.md
├── api
│   ├── buf.lock
│   ├── buf.yaml
│   ├── demo
│   │   └── v1
│   │       ├── demo.pb.go
│   │       ├── demo.pb.validate.go
│   │       ├── demo.proto
│   │       ├── demo_grpc.pb.go
│   │       └── demo_http.pb.go
│   └── openapi.yaml
├── buf.gen.yaml
└── buf.work.yaml
```

```bash
buf.yaml 定义存放 proto api 文件的根目录 管理proto依赖 lint break change等
buf.lock buf build 之后根据buf.yaml自动生成的文件，用来记录和锁定依赖
buf.work.yaml 构建工作目录下，定义了当前工作目录下需要编译的proto api文件目录
buf.gen.yaml 构建工作目录下，定义了编译proto文件时的具体规则和插件信息
```

buf 初始化

```bash
buf mod init
```

编写proto文件

在buf.yaml写入需要引入的proto

```yaml
version: v1
breaking:
  use:
    - FILE
deps:
  - buf.build/googleapis/googleapis
  - buf.build/envoyproxy/protoc-gen-validate
  - buf.build/kratos/apis
lint:
  use:
    - DEFAULT
```

然后执行 `buf mod update` 拉取文件

执行编译 `buf build`

返回生成目录 创建 `buf.gen.yaml` and `buf.work.yaml`

buf.gen.yaml 编写所有是用的插件 版本 输出目录
```yaml
version: v1
plugins:
  # Use protoc-gen-go at v1.28.1
  - plugin: buf.build/protocolbuffers/go:v1.28.1
    out: api
    opt: paths=source_relative
  # Use the latest version of protoc-gen-go-grpc
  - plugin: buf.build/grpc/go
    out: api
    opt:
      - paths=source_relative
      - require_unimplemented_servers=false
  - plugin: go-http
    out: api
    opt:
      - paths=source_relative
  - plugin: openapi
    out: api
    opt:
      - paths=source_relative
  # Use the latest version of protoc-gen-validate
  - plugin:  buf.build/bufbuild/validate-go
    out: api
    opt:
      - paths=source_relative
  - plugin: go-errors
    out: api
    opt:
      - paths=source_relative
```
buf.work.yaml 编写 我们的工作目录

```yaml
version: v1
directories:
  - api
```

然后执行 buf generate 成功生成内容如下

```bash
.
├── buf.lock
├── buf.yaml
├── demo
│   └── v1
│       ├── demo.pb.go
│       ├── demo.pb.validate.go
│       ├── demo.proto
│       ├── demo_grpc.pb.go
│       └── demo_http.pb.go
└── openapi.yaml
```
