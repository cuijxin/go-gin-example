package util

import "github.com/cuijxin/go-gin-example/pkg/setting"

func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
