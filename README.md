# gin-code-generate
基于gin制作的mvc代码目录生成器
- gin       web框架
- gorm      orm工具
- zap       日志工具
- viper     配置工具
- urfave    命令行工具

使用方式

```
./gin-code-generate --path 目录 --project-name 项目名称 --mod
```



## 代码目录结构

```bash
.
├── Dockerfile          # 构建镜像，直接跑代码
├── README.md           # 阐述项目
├── config              # 配置文件目录
│   ├── config.go       # 配置文件
│   ├── config.yaml     # 配置文件yaml文件，可以使用ini
│   └── option.go       # 运行时的参数
├── go.mod              # go mod 文件
├── handlers            # 控制层目录
├── log                 # 日志打印目录，可以在配置文件配置
├── main.go             # 入口函数
├── middleware          # 中间件
│   └── log.go          # 日志中间件，使用zap工具
├── models              # 模型层目录
│   └── model.go        # 模型层初始化文件
├── proto               # protobuf dil
├── routers             # 路由层目录
│   └── router.go       # 路由初始化文件
├── templates           # 模版层
└── utils               # 工具
    └── response        # 回调使用

```


## 补充
### 工具

命令行工具  https://github.com/urfave/cli 可选项 https://github.com/spf13/cobra

```
# 生成linux-adm64的二进制命令
docker run -it --rm -e GOPROXY=https://goproxy.io -v /Users/hanke/go/src/github.com/agocan/gin-code-generate:/app golang:1.14-alpine3.11 'cd /app; go build'
```