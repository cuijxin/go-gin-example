package routers

import (
	"net/http"

	"github.com/cuijxin/go-gin-example/middleware/jwt"
	"github.com/cuijxin/go-gin-example/pkg/export"
	"github.com/cuijxin/go-gin-example/pkg/qrcode"
	"github.com/cuijxin/go-gin-example/pkg/setting"
	"github.com/cuijxin/go-gin-example/pkg/upload"
	"github.com/cuijxin/go-gin-example/routers/api"
	v1 "github.com/cuijxin/go-gin-example/routers/api/v1"

	"github.com/gin-gonic/gin"

	_ "github.com/cuijxin/go-gin-example/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	// 实现http.FileServer
	// 为了让前端访问到图片需要实现http.FileServer
	// 一般如下：
	// 1. CDN (公司使用CDN或自建分布式文件系统)
	// 2. http.FileSystem （测试用，本地搭建）
	// 本地搭建的话，Go本身对此就有很好的支持，而Gin更是再封装了一层，只需要在路由增加一行代码即可
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// 新建标签
		apiv1.POST("/tags", v1.AddTag)
		// 更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//  导出标签
		apiv1.POST("/tags/export", v1.ExportTag)
		//导入标签
		apiv1.POST("/tags/import", v1.ImportTag)

		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		// 新建文章
		apiv1.POST("/articles", v1.AddArticle)
		// 更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		// 删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
