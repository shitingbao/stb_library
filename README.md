# 一个以 Kratos 为基础的项目例子

官方的文档地址 : <https://go-kratos.dev/docs/>

## 功能包含

### 1. 多服务之间的架构分类  

  多个服务项目目录以及对应 proto 结构划分

  --api（所有服务的 protoc）  
    --项目 a 对应 protoc  
    --项目 b 对应 protoc  
  --app（项目目录）  
    --项目 a  
    --项目 b  
  --lib（功能目录）

### 2. http ，rpc ，路由，websocket

#### 2.1 protoc api路由

    central 项目，路由定义在 api/central/v1/central.proto 中，使用 google api 定义了 SayHello 和 Healthy

#### 2.2 gin 路由兼容

    storage/internal/sgin 定义了 gin 核心，类似（service），这里为了方便重命名一个文件夹，也可以放在 service 中；

#### 2.3 websocket

    在 http 加载的时候写入。storage/internal/server/http.go 中加入了 websocket 路由，具体的方法封装在 lib/ws 中；

### 3. 服务发现

  1.kratos 项目的 grpc 调用 和 2.其他项目的服务调用，也就是说调用不使用 kratos 的项目的 rpc 通讯。

#### 3.1 调用 kratos 对应的服务

    kratos使用了 consul 来进行注册和发现，需要先下载开启，并且在 config 中写入地址。 app/storage/internal/data.go 中定义了服务的注册和发现基本组件 NewDiscovery 和 NewRegistrar，对应了服务的注册和发现。
  建立一个 grpc 连接，相同 kratos 中的服务 NewCentralGrpcClient ，其中 WithEndpoint 写入对应的规则（与 main 中定义的 Name 变量对应）
  
#### 3.2 调用第三方的 grpc

    进行连接，也就是不是使用 kratos 写的，同时也不注册在这些服务列表内。
  app/storage/internal/slog.go 中的 NewSlogServiceClient 连接了一个外部的服务，只需要将 grpc.WithEndpoint 中的内容改为 ip+port 即可；

### 4.自定义log

kratos 中带的 logger ，只要实现了对应的 Log(level Level, keyvals ...interface{}) error 方法即可；  
这里本项目中不使用本地收集的方法（因为官方例子中都有），我这里用了把日志发送到第三方服务的方式来收集；

## 打包及体积压缩（）

1. 构建使用 -ldflags 参数

example
```
go build -ldflags "-w -s"
```

2. 使用 upx 将生成的执行文件压缩打包

example
```
upx --brute test.exe
```

// 注意，如果不是对体积有特殊的要求，upx 会使你的内存使用率增高
原因：工作原理.它们压缩磁盘文件,而不是可执行代码.要使压缩文件再次可执行,需要将其解压缩,并将未压缩的数据存储在内存中.使用普通的非压缩EXE文件,操作系统将仅加载此刻所需的文件部分.其余的可以留在磁盘上.由于整个未压缩的应用程序都在内存中,因此您的内存使用率会更高

// 注意2，如果对接 c ，那可能会损坏某些 dll