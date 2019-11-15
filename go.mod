module github.com/cuijxin/go-gin-example

go 1.13

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.51.0
	github.com/jinzhu/gorm v1.9.11
	github.com/unknwon/com v1.0.1
	google.golang.org/appengine v1.6.5 // indirect
)

replace (
	github.com/cuijxin/go-gin-example/models => ./go-gin-example/models
	github.com/cuijxin/go-gin-example/pkg/setting => ./go-gin-example/pkg/setting
	github.com/cuijxin/go-gin-example/routers => ./go-gin-example/routers
)
