# go-gin-example

```
go-gin-example/
|--- conf
|--- middleware
|--- models
|--- pkg
|--- routers
|--- runtime
```

 * conf: 用于存储配置文件
 * middleware: 应用中间件
 * models: 应用数据库模型
 * pkg: 第三方包
 * routers: 路由逻辑处理
 * runtime: 应用运行时数据

app.ini
 ```
[app]
PageSize = 10
JwtSecret = 233

RuntimeRootPath = runtime/

ImagePrefixUrl = http://127.0.0.1:8000
ImageSavePath = upload/images/

# MB
ImageMaxSize = 5
ImageAllowExts = .jpg,.jpeg,.png

LogSavePath = logs/
LogSaveName = log
LogFileExt = log
TimeFormat = 20060102

[server]
#debug or release
RunMode = debug
HttpPort = 8000
ReadTimeout = 60
WriteTimeout = 60

[database]
Type = mysql
User = 数据库用户名
Password = 数据库密码
Host = 数据库地址及端口号
Name = blog
TablePrefix = blog_
 ```
