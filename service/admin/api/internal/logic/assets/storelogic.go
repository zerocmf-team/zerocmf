package assets

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/sha1"
	"errors"
	"fmt"
	"gincmf/common/bootstrap/util"
	"gincmf/service/admin/model"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"

	"gincmf/service/admin/api/internal/svc"
	"gincmf/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type assetsResult struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	PrevPath string `json:"prev_path"`
}

func NewStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) StoreLogic {
	return StoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreLogic) Store(req types.AssetsRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	resp = new(types.Response)

	c := l.svcCtx
	db := c.Db
	r := c.Request
	r.ParseMultipartForm(32 << 20)
	files := r.MultipartForm.File["file"]

	if len(files) <= 0 {
		result := c.Error("文件不能为空！", nil)
		copier.Copy(&resp, &result)
		return
	}

	fileType := req.Type
	var assets []assetsResult
	var fileList map[string]string

	for _, fileItem := range files {
		fileList, err = handleUpload(c, db, fileItem, fileType)
		if err != nil {
			result := c.Error("文件不能为空！", nil)
			copier.Copy(&resp, &result)
		}
		assets = append(assets, assetsResult{FileName: fileList["fileName"], FilePath: fileList["filePath"], PrevPath: fileList["prevPath"]})
	}

	result := c.Success("上传成功！", assets)
	copier.Copy(&resp, &result)
	return
}

// 根据文件处理上传逻辑
// 1.判断上传类型，验证后缀合理性 type [0 => "图片" 1 => "视频" 2 => "文件"]
func handleUpload(c *svc.ServiceContext, db *gorm.DB, file *multipart.FileHeader, fileType string) (result map[string]string, err error) {
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

	//fileMd5 := hex.EncodeToString(md5h.Sum([]byte("")))

	sha1h := sha1.New()
	sha1h.Write(buf.Bytes())

	// fileSha1 := hex.EncodeToString(sha1h.Sum([]byte("")))

	// 上传文件至指定目录

	saveUploadedFile(file, realpath)

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

	result = make(map[string]string, 0)
	result["fileName"] = fileNameSuffix
	result["filePath"] = filePath
	result["prevPath"] = util.FileUrl(filePath)

	return
}

func saveUploadedFile(file *multipart.FileHeader, dst string) (err error) {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
