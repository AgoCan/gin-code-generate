# gin-code-generate
基于gin和gorm的mvc框架目录结构构建

https://github.com/urfave/cli 命令行工具

## 运行方式

```bash
./
```

## 代码目录结构

```bash
.
├── Dockerfile          # 构建镜像，直接跑代码
├── README.md           # 阐述项目
├── client              # 目录
│   └── client.go       # 
├── config              # 目录，配置文件使用yaml
│   ├── config.go       # 解析配置文件 
│   ├── config.yaml     # yaml的配置文件
│   ├── config_test.go  # 测试代码
│   └── configstruct.go # 把yaml配置文件解析成config.go使用
├── docs                # 目录，文档
│   ├── README.md       # 文档入口
│   ├── apis.md         # 接口文档
│   └── deployment.md   # 部署文档
├── go.mod              # 使用go mod
├── go.sum
├── handlers             # 控制层
│   └── example
│       └── example.go
├── main.go             # 入口函数
├── log                 # 日志打印默认位置
├── middleware          # 中间件
│   ├── log.go
├── model               # 模型层
│   ├── model.go
├── proto               # protobuf dil
├── routers             # 路由层
│   └── router.go
├── templates           # 渲染
│   └── base.tmpl
└── utils               # 工具类
    ├── base.go         # 基础工具类
    ├── logging         # 手动打印的日志
    │   ├── log.go
    │   └── log_test.go
    └── response        # 固定返回格式
        └── response.go
```
