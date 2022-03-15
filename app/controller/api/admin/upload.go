package admin

import (
	"encoding/json"
	"fmt"
	"gincmf/app/model"
	"gincmf/app/util"
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	"github.com/gincmf/cmf/controller"
	"strconv"
)

// AssetsController 图片资源控制器，定义了资源文件增删改查接口
type UploadController struct {
	rc controller.RestController
}

func (rest *UploadController) Get(c *gin.Context) {
	uploadSetting := util.UploadSetting(c)
	rest.rc.Success(c, "获取成功！", uploadSetting)
}

func (rest *UploadController) Show(c *gin.Context) {
	var rewrite struct {
		id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	rest.rc.Success(c, "操作成功show", nil)
}

func (rest *UploadController) Edit(c *gin.Context) {
	rest.rc.Success(c, "操作成功Edit", nil)
}

func (rest *UploadController) Store(c *gin.Context) {

	maxFiles := c.PostForm("max_files")
	chunkSize := c.PostForm("chunk_size")
	imageMaxFileSize := c.PostForm("file_types[image][upload_max_file_size]")
	imageExtensions := c.PostForm("file_types[image][extensions]")

	videoMaxFileSize := c.PostForm("file_types[video][upload_max_file_size]")
	videoExtensions := c.PostForm("file_types[video][extensions]")

	audioMaxFileSize := c.PostForm("file_types[audio][upload_max_file_size]")
	audioExtensions := c.PostForm("file_types[audio][extensions]")

	fileMaxFileSize := c.PostForm("file_types[file][upload_max_file_size]")
	fileExtensions := c.PostForm("file_types[file][extensions]")

	maxFilesInt ,err := strconv.Atoi(maxFiles)
	if err != nil {
		rest.rc.Error(c,"最大同时上传文件数必须为数字",nil)
		return
	}
	chunkSizeInt, err := strconv.Atoi(chunkSize)
	if err != nil {
		rest.rc.Error(c,"文件分块上传分块大小必须为数字",nil)
		return
	}
	imageMaxFileSizeInt ,err := strconv.Atoi(imageMaxFileSize)
	if err != nil {
		rest.rc.Error(c,"允许图片上传大小必须为数字",nil)
		return
	}
	videoMaxFileSizeInt ,err := strconv.Atoi(videoMaxFileSize)
	if err != nil {
		rest.rc.Error(c,"允许视频上传大小必须为数字",nil)
		return
	}
	audioMaxFileSizeInt ,err := strconv.Atoi(audioMaxFileSize)
	if err != nil {
		rest.rc.Error(c,"允许音频上传大小必须为数字",nil)
		return
	}
	fileMaxFileSizeInt ,err := strconv.Atoi(fileMaxFileSize)
	if err != nil {
		rest.rc.Error(c,"允许音频上传大小必须为数字",nil)
		return
	}

	uploadSetting := &model.UploadSetting{
		MaxFiles:  maxFilesInt,
		ChunkSize: chunkSizeInt,
		FileTypes: model.FileTypes{
			Image: model.TypeValues{
				UploadMaxFileSize: imageMaxFileSizeInt,
				Extensions:        imageExtensions,
			},
			Video: model.TypeValues{
				UploadMaxFileSize: videoMaxFileSizeInt,
				Extensions:        videoExtensions,
			},
			Audio: model.TypeValues{
				UploadMaxFileSize: audioMaxFileSizeInt,
				Extensions:        audioExtensions,
			},
			File: model.TypeValues{
				UploadMaxFileSize: fileMaxFileSizeInt,
				Extensions:        fileExtensions,
			},
		},
	}
	uploadSettingValue, _ := json.Marshal(uploadSetting)
	fmt.Println("uploadSettingValue", string(uploadSettingValue))

	cmf.Db.Model(&model.Option{}).Where("option_name = ?","upload_setting").Update("option_value", string(uploadSettingValue))

	rest.rc.Success(c, "修改成功",uploadSetting)
}

func (rest *UploadController) Delete(c *gin.Context) {
	rest.rc.Success(c, "操作成功Delete", nil)
}
