/**
** @创建时间: 2020/7/15 10:41 下午
** @作者　　: return
 */
package admin

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"gincmf/app/model"
	"gincmf/app/util"
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	"github.com/gincmf/cmf/controller"
	cmfUtil "github.com/gincmf/cmf/util"
	uuid "github.com/nu7hatch/gouuid"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type AssetsController struct {
	rc controller.RestController
}

/*
 * @restApi(
 *     'name'  		=> '资源文件列表',
 *	   'desc'  		=> '资源文件列表'
 *     'url'   		=> 'api/admin/asset',
 *	   'param' 		=>  '',
 *	   'method'		=> 'get',
 *	   'list_order' => '10000',
 *	   'status'		=> 1
 * )
 */

func (rest *AssetsController) Get(c *gin.Context) {

	var assets []model.Asset
	query := "status = ?"
	queryArgs := []interface{}{"1"}

	paramType := c.DefaultQuery("type", "0")

	query += " AND type = ?"
	queryArgs = append(queryArgs, paramType)

	current := c.DefaultQuery("current", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	intCurrent, _ := strconv.Atoi(current)
	intPageSize, _ := strconv.Atoi(pageSize)

	if intCurrent <= 0 {
		rest.rc.Error(c, "当前页码需大于0！", nil)
		return
	}

	if intPageSize <= 0 {
		rest.rc.Error(c, "每页数需大于0！", nil)
		return
	}

	var total int64 = 0
	cmf.Db.Where(query, queryArgs...).Find(&assets).Count(&total)
	result := cmf.Db.Where(query, queryArgs...).Limit(intPageSize).Offset((intCurrent - 1) * intPageSize).Order("id desc").Find(&assets)

	if result.RowsAffected == 0 {
		rest.rc.Error(c, "该页码内容不存在！", nil)
		return
	}

	var tempAssets []model.TempAsset

	for _, v := range assets {
		prevPath := util.GetFileUrl(v.FilePath)
		tempAssets = append(tempAssets, model.TempAsset{Asset: v, PrevPath: prevPath})
	}

	paginationData := &model.Paginate{Data: tempAssets, Current: current, PageSize: pageSize, Total: total}
	if len(tempAssets) == 0 {
		paginationData.Data = make([]string, 0)
	}

	rest.rc.Success(c, "获取成功！", paginationData)
}

/*
 * @restApi(
 *     'name'  		=> '获取单个资源文件',
 *	   'desc'  		=> '获取单个资源文件'
 *     'url'   		=> 'api/admin/asset/:id',
 *	   'param' 		=>  '',
 *	   'method'		=> 'get',
 *	   'list_order' => '10000',
 *	   'status'		=> 1
 * )
 */
func (rest *AssetsController) Show(c *gin.Context) {
	var rewrite struct {
		id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	rest.rc.Success(c, "操作成功show", nil)
}

func (rest *AssetsController) Edit(c *gin.Context) {
	rest.rc.Success(c, "操作成功Edit", nil)
}

/*
 * @restApi(
 *     'name'  		=> '上传资源文件',
 *	   'desc'  		=> '上传单个或多个资源文件'
 *     'url'   		=> 'api/admin/asset',
 *	   'param' 		=>  '',
 *	   'method'		=> 'post',
 *	   'list_order' => '10000',
 *	   'status'		=> 1
 * )
 */
func (rest *AssetsController) Store(c *gin.Context) {

	form, _ := c.MultipartForm()
	files := form.File["file[]"]

	fileType := c.DefaultPostForm("type", "0")

	if len(files) <= 0 {
		rest.rc.Error(c, "图片不能为空！", nil)
		return
	}



	var fileList map[string]string
	var err error

	type tempAssets struct {
		FileName string `json:"file_name"`
		FilePath string `json:"file_path"`
		PrevPath string `json:"prev_path"`
	}
	var result []tempAssets

	for _, fileItem := range files {
		fileList, err = handleUpload(c, fileItem, fileType)
		if err != nil {
			rest.rc.Error(c, err.Error(), nil)
			return
		}
		result = append(result,tempAssets{FileName: fileList["fileName"],FilePath: fileList["filePath"],PrevPath: fileList["prevPath"]})
	}

	rest.rc.Success(c, "上传成功", result)
}

/*
 * @restApi(
 *     'name'  		=> '删除资源文件',
 *	   'desc'  		=> '删除单个或多个资源文件'
 *     'url'   		=> 'api/admin/asset',
 *	   'param' 		=>  '',
 *	   'method'		=> 'delete',
 *	   'list_order' => '10000',
 *	   'status'		=> 1
 * )
 */
func (rest *AssetsController) Delete(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}

	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	ids := c.QueryArray("ids")

	fmt.Println("first_ids", ids)
	asset := &model.Asset{}

	if len(ids) == 0 {
		if err := c.ShouldBindUri(&rewrite); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		fmt.Println("Id", rewrite.Id)

		result := cmf.Db.First(&asset, rewrite.Id)
		if result.RowsAffected == 0 {
			rest.rc.Error(c, "该内容不存在！", nil)
			return
		}

		asset.Id = rewrite.Id
		asset.Status = 0

		if err := cmf.Db.Save(asset).Error; err != nil {
			rest.rc.Error(c, "删除失败！", nil)
			return
		}
	} else {
		fmt.Println("ids", ids)
		if err := cmf.Db.Model(&asset).Where("id IN (?)", ids).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
			rest.rc.Error(c, "删除失败！", nil)
			return
		}
	}

	rest.rc.Success(c, "删除成功！", nil)
}

// 根据文件处理上传逻辑
// 1.判断上传类型，验证后缀合理性 type [0 => "图片" 1 => "视频" 2 => "文件"]
func handleUpload(c *gin.Context, file *multipart.FileHeader, fileType string)  (map[string]string,error) {
	tempFile, tempErr := file.Open()
	defer tempFile.Close()

	if tempErr != nil {
		fmt.Println("tempErr", tempErr)
	}

	var fileSize int64 = 0

	type Size interface {
		Size() int64
	}

	if sizeInterface, ok := tempFile.(Size); ok {
		fileSize = sizeInterface.Size()
	}

	fmt.Println("fileSize", fileSize)

	suffixArr := strings.Split(file.Filename, ".")

	suffix := suffixArr[len(suffixArr)-1]

	fmt.Println("suffix", suffix)

	//获取后缀列表
	uploadSetting := util.UploadSetting(c)

	fmt.Println("fileType", fileType)

	switch fileType {
	case "0":
		iExtensionArr := strings.Split(uploadSetting.Image.Extensions, ",")
		iResult := util.ToLowerInArray(suffix, iExtensionArr)
		fmt.Println("iResult", iResult)
		if !iResult {
			return nil,errors.New("【" + suffix + "】不是合法的图片后缀！")
		}
	case "1":
		aExtensionArr := strings.Split(uploadSetting.Audio.Extensions, ",")
		if !util.ToLowerInArray(suffix, aExtensionArr) {
			return nil,errors.New("【" + suffix + "】不是合法的音频后缀！")
		}
	case "2":
		vExtensionArr := strings.Split(uploadSetting.Video.Extensions, ",")
		if !util.ToLowerInArray(suffix, vExtensionArr) {
			return nil,errors.New("【" + suffix + "】不是合法的音频后缀！")
		}
	case "3":
		fExtensionArr := strings.Split(uploadSetting.File.Extensions, ",")
		if !util.ToLowerInArray(suffix, fExtensionArr) {
			return nil,errors.New("【" + suffix + "】不是合法的附件后缀！")
		}

	default:
		return nil,errors.New("非法访问")
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
	fileName := cmfUtil.GetMd5(fileUuid.String() + suffixArr[0])
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
	if _, err := io.Copy(buf, tempFile); err != nil {
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

	userId, _ := c.Get("user_id")
	fmt.Println("userId", userId)
	if userId == nil {
		userId = "0"
	}
	userIdInt, _ := strconv.Atoi(userId.(string))

	fileTypeInt, _ := strconv.Atoi(fileType)
	//保存到数据库
	cmf.Db.Create(&model.Asset{
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

	tempMap := make(map[string]string,0)
	tempMap["fileName"] = fileNameSuffix
	tempMap["filePath"] = filePath
	tempMap["prevPath"] = util.GetFileUrl(filePath)

	return tempMap,nil
}
