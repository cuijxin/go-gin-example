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
 #debug or release
RUN_MODE = debug

[app]
PAGE_SIZE = 10
JWT_SECRET = 23347$040412

[server]
HTTP_PORT = 8000
READ_TIMEOUT = 60
WRITE_TIMEOUT = 60

[database]
TYPE = mysql
USER = 数据库账号
PASSWORD = 数据库密码
HOST = 数据库IP:数据库端口号
NAME = blog
TABLE_PREFIX = blog_
 ```
