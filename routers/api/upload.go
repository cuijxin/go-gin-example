package api

// 在这一大段的业务逻辑中，我们做了如下事情：
// c.Request.FormFile：获取上传的图片（返回提供的表单键的第一个文件）
// CheckImageExt、CheckImageSize检查图片大小，检查图片后缀
// CheckImage：检查上传图片所需（权限、文件夹）
// SaveUploadedFile：保存图片
// 总的来说，就是 入参 -> 检查 -》 保存 的应用流程

import (
	"net/http"

	"github.com/cuijxin/go-gin-example/pkg/app"
	"github.com/cuijxin/go-gin-example/pkg/e"
	"github.com/cuijxin/go-gin-example/pkg/logging"
	"github.com/cuijxin/go-gin-example/pkg/upload"
	"github.com/gin-gonic/gin"
)

// @Summary Import Image
// @Produce json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]
func UploadImage(c *gin.Context) {
	appG := app.Gin{C: c}
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}
