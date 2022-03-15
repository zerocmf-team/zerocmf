/**
** @创建时间: 2021/12/6 22:32
** @作者　　: return
** @描述　　:
 */

package service

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/paginate"
	"github.com/gincmf/bootstrap/util"
	"github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type Assets struct{}

/**
 * @Author return <1140444693@qq.com>
 * @Description
 * @Date 2021/12/6 23:25:18
 * @Param
 * @return
 **/

func (service *Assets) Get(c *gin.Context, db *gorm.DB, query string, queryArgs []interface{}) (paginateData paginate.Paginate, err error) {
	current, pageSize, err := new(paginate.Paginate).Default(c)
	if err != nil {
		return paginateData, errors.New("服务器错误！")
	}
	paginateData, err = new(model.Assets).Get(db, current, pageSize, query, queryArgs)
	return paginateData, err
}

type assetsResult struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	PrevPath string `json:"prev_path"`
}

func (service *Assets) Store(c *gin.Context, db *gorm.DB) (result []assetsResult, err error) {

	form, _ := c.MultipartForm()
	files := form.File["file[]"]

	fileType := c.DefaultPostForm("type", "0")

	if len(files) <= 0 {
		return result, errors.New("图片不能为空！")
	}

	var fileList map[string]string

	for _, fileItem := range files {
		fileList, err = handleUpload(c, db, fileItem, fileType)
		if err != nil {
			return result, err
		}
		result = append(result, assetsResult{FileName: fileList["fileName"], FilePath: fileList["filePath"], PrevPath: fileList["prevPath"]})
	}
	return result, err
}

// 根据文件处理上传逻辑
// 1.判断上传类型，验证后缀合理性 type [0 => "图片" 1 => "视频" 2 => "文件"]
func handleUpload(c *gin.Context, db *gorm.DB, file *multipart.FileHeader, fileType string) (result map[string]string, err error) {
	mulFile, mulErr := file.Open()
	defer mulFile.Close()

	if mulErr != nil {
		return result, mulErr
	}

	var fileSize int64 = 0

	type Size interface {
		Size() int64
	}

	if sizeInterface, ok := mulFile.(Size); ok {
		fileSize = sizeInterface.Size()
	}

	suffixArr := strings.Split(file.Filename, ".")

	suffix := suffixArr[len(suffixArr)-1]

	uploadSetting, err := model.UploadSettings(db)

	if err != nil {
		return result, err
	}

	//获取后缀列表

	switch fileType {
	case "0":
		iExtensionArr := strings.Split(uploadSetting.Image.Extensions, ",")
		iResult := util.ToLowerInArray(suffix, iExtensionArr)
		fmt.Println("iResult", iResult)
		if !iResult {
			return nil, errors.New("【" + suffix + "】不是合法的图片后缀！")
		}
	case "1":
		aExtensionArr := strings.Split(uploadSetting.Audio.Extensions, ",")
		if !util.ToLowerInArray(suffix, aExtensionArr) {
			return nil, errors.New("【" + suffix + "】不是合法的音频后缀！")
		}
	case "2":
		vExtensionArr := strings.Split(uploadSetting.Video.Extensions, ",")
		if !util.ToLowerInArray(suffix, vExtensionArr) {
			return nil, errors.New("【" + suffix + "】不是合法的音频后缀！")
		}
	case "3":
		fExtensionArr := strings.Split(uploadSetting.File.Extensions, ",")
		if !util.ToLowerInArray(suffix, fExtensionArr) {
			return nil, errors.New("【" + suffix + "】不是合法的附件后缀！")
		}

	default:
		return nil, errors.New("非法访问")
		c.Abort()
	}

	path := "public/uploads"
	t := time.Now()
	timeArr := []int{t.Year(), int(t.Month()), t.Day()}

	var timeDir string
	for key, timeInt := range timeArr {

		current := strconv.Itoa(timeInt)
		if key > 0 {
			if len(current) <= 1 {
				current = "0" + current
			}
		}
		// tempStr := "/" + current
		timeDir += current
	}

	temPath := "default"

	fileUuid, err := uuid.NewV4()

	remarkName := file.Filename
	fileName := util.GetMd5(fileUuid.String() + suffixArr[0])
	fileNameSuffix := fileName + "." + suffix

	uploadPath := temPath + "/" + timeDir + "/"
	filePath := uploadPath + fileNameSuffix
	realpath := path + "/" + filePath

	fmt.Println("上传路径:", path+"/"+uploadPath)
	_, err = os.Stat(path + "/" + uploadPath)
	if err != nil {
		os.MkdirAll(path+"/"+uploadPath, os.ModePerm)
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, mulFile); err != nil {
		fmt.Println(err)
	}

	md5h := md5.New()
	md5h.Write(buf.Bytes())

	fileMd5 := hex.EncodeToString(md5h.Sum([]byte("")))
	log.Println("md5", fileMd5)

	sha1h := sha1.New()
	sha1h.Write(buf.Bytes())

	fileSha1 := hex.EncodeToString(sha1h.Sum([]byte("")))
	log.Println("sha1", fileSha1)
	// 上传文件至指定目录

	c.SaveUploadedFile(file, realpath)

	userId, _ := c.Get("userId")
	if userId == nil {
		userId = "0"
	}
	userIdInt, _ := strconv.Atoi(userId.(string))

	fileTypeInt, _ := strconv.Atoi(fileType)
	//保存到数据库
	db.Create(&model.Assets{
		UserId:     userIdInt,
		FileSize:   fileSize,
		CreateAt:   time.Now().Unix(),
		FileKey:    fileUuid.String(),
		RemarkName: remarkName,
		FileName:   fileNameSuffix,
		FilePath:   filePath,
		Suffix:     suffix,
		AssetType:  fileTypeInt,
	})

	tempMap := make(map[string]string, 0)
	tempMap["fileName"] = fileNameSuffix
	tempMap["filePath"] = filePath
	tempMap["prevPath"] = util.FileUrl(filePath)

	return tempMap, nil
}
