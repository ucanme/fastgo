### FastGO 框架
#### 简介
一个简单实用的go http框架, 支持命令行自定义cmd操作.基于gin web框架，gorm db操作，用起来比较顺手的轮子。
#### 特性
1. 支持access 与 panic recovery中间件,内含完整的日志记录。
2. 基于logrus封装有log类库， 日志等级分级与自动切割，支持自定义日志保存和分割周期。
3. 支持秒级cron定时任务配置，内含demo。
4. 基于gorm，框架内含db初始化工具。
5. 基于toml配置文件，含有完整解析类库
6. 内含http请求工具类库。
7. 使用go mod包管理工具，不依赖GOPATH的设置，需要go1.13以上版本
8. 内含Dockerfile与Makefile支持一键编译docker镜像，内含docker-compose.yml示例支持一键容器运行

#### 基本使用

```
#基于命令行实用
 go run main.go server #web服务运行
 go run main.go init-db #db初始化
#基于docker实用
 make docker #编译生成镜像
 docker-compose up #docker-compose运行,需提前替换docker-compose.yml中的镜像
```

#### 如何获取
github仓库地址: https://github.com/ucanme/fastgo.git
欢迎批评指正，轮子会不断升级维护。使用交流QQ群: 15895722

#### 优惠福利
阿里云服务福利疫情最后几天活动 2核8g内存40G磁盘5m带宽三年1399，0.6折价格可做开发机，学习机，业务机，技术在于折腾。购买地址：https://www.aliyun.com/minisite/goods?userCode=b2d0no2s  


#### 最近更新:
新增Trace分支,支持requestId的track功能。

通过修改源代码获取goroutine id,可以尝试直接修改源代码src/runtime/runtime2.go中添加Goid函数，将goid暴露给应用层：
```
func Goid() int64 {
    _g_ := getg()
    return _g_.goid
}
```

