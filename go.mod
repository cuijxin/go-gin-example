module github.com/cuijxin/go-gin-example

go 1.13

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.0
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.51.0
	github.com/go-openapi/jsonreference v0.19.3 // indirect
	github.com/go-openapi/spec v0.19.4 // indirect
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/jinzhu/gorm v1.9.11
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/mailru/easyjson v0.7.0 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/robfig/cron v1.2.0
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.3
	github.com/ugorji/go v1.1.7 // indirect
	github.com/unknwon/com v1.0.1
	github.com/urfave/cli v1.22.1 // indirect
	golang.org/x/net v0.0.0-20191118183410-d06c31c94cae // indirect
	golang.org/x/sys v0.0.0-20191118133127-cf1e2d577169 // indirect
	golang.org/x/tools v0.0.0-20191118222007-07fc4c7f2b98 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
)

replace (
	github.com/cuijxin/go-gin-example/docs => ./go-gin-example/docs
	github.com/cuijxin/go-gin-example/middleware/jwt => ./go-gin-example/middleware/jwt
	github.com/cuijxin/go-gin-example/models => ./go-gin-example/models
	github.com/cuijxin/go-gin-example/pkg/e => ./go-gin-example/pkg/e
	github.com/cuijxin/go-gin-example/pkg/file => ./go-gin-example/pkg/file
	github.com/cuijxin/go-gin-example/pkg/logging => ./go-gin-example/pkg/logging
	github.com/cuijxin/go-gin-example/pkg/setting => ./go-gin-example/pkg/setting
	github.com/cuijxin/go-gin-example/pkg/upload => ./go-gin-example/pkg/upload
	github.com/cuijxin/go-gin-example/pkg/util => ./go-gin-example/pkg/util
	github.com/cuijxin/go-gin-example/routers => ./go-gin-example/routers

)
