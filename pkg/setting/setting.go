// 配置统管
// 在Go中，当存在多个init函数时，执行顺序为：
// 1. 相同包下的init函数：按照源文件编译顺序决定执行顺序（默认按文件名排序）
// 2. 不同包下的init函数：按照包导入的依赖关系决定先后顺序
// 所以要避免多init的情况，尽量由程序包括初始化的先后顺序
package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	// ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	QrCodeSavePath string
	FontSavePath   string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

// var (
// 	Cfg *ini.File

// 	RunMode string

// 	HTTPPort int

// 	ReadTimeout  time.Duration
// 	WriteTimeout time.Duration

// 	PageSize  int
// 	JwtSecret string
// )

// func init() {
// 	var err error
// 	Cfg, err = ini.Load("conf/app.ini")
// 	if err != nil {
// 		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
// 	}

// 	LoadBase()
// 	LoadServer()
// 	LoadApp()
// }

// func LoadBase() {
// 	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
// }

// func LoadServer() {
// 	sec, err := Cfg.GetSection("server")
// 	if err != nil {
// 		log.Fatalf("Fail to get section 'server': %v", err)
// 	}

// 	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
// 	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
// 	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
// }

// func LoadApp() {
// 	sec, err := Cfg.GetSection("app")
// 	if err != nil {
// 		log.Fatalf("Fail to get sectionn 'app': %v", err)
// 	}

// 	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
// 	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
// }
