### FastGO 框架
####简介:一个简单的go http框架, 支持命令行自定义cmd操作.基于gin web框架，gorm db操作.
####特性
1. 支持access 与 panic recovery中间件,内含完整的日志记录。
2. 基于logrus封装有log类库， 日志等级分级与自动切割，支持自定义日志保存和分割周期。
3. 支持秒级cron定时任务配置，内含demo。
4. 基于gorm，框架内含db初始化工具。
5. 基于toml配置，含有完整解析类库
6. 内含http请求工具类库。
7. 使用go mod包管理工具，不依赖GOPATH的设置
8. 内含Dockerfile与Makefile支持一键编译docker镜像，内涵docker-compose.yml示例支持一键容器运行