package main

import (
	"fmt"
	"syscall"

	"github.com/cuijxin/go-gin-example/models"
	"github.com/cuijxin/go-gin-example/pkg/gredis"
	"github.com/cuijxin/go-gin-example/pkg/logging"
	"github.com/cuijxin/go-gin-example/pkg/setting"
	"github.com/cuijxin/go-gin-example/pkg/util"
	"github.com/cuijxin/go-gin-example/routers"
	"github.com/fvbock/endless"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		logging.Info("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		logging.Info("Server err: %v", err)
	}
	// ==========================================

	// router := routers.InitRouter()

	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	// 	Handler:        router,
	// 	ReadTimeout:    setting.ReadTimeout,
	// 	WriteTimeout:   setting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// s.ListenAndServe()

	// ===========================================

	// router := routers.InitRouter()

	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	// 	Handler:        router,
	// 	ReadTimeout:    setting.ReadTimeout,
	// 	WriteTimeout:   setting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// go func() {
	// 	if err := s.ListenAndServe(); err != nil {
	// 		logging.Info("Listen: %s\n", err)
	// 	}
	// }()

	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	// <-quit

	// logging.Info("shutdown server ...")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := s.Shutdown(ctx); err != nil {
	// 	logging.Fatal("server shutdown:", err)
	// }

	// logging.Info("server exiting")
}
