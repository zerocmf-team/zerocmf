package assets

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	_ "image/jpeg" // 支持 JPEG 格式
	_ "image/png"  // 支持 PNG 格式
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/admin/model"

	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"

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

func (l *StoreLogic) Store(req *types.AssetsReq) (resp *types.Response) {

	resp = new(types.Response)
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	r := c.Request
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		resp.Error("上传失败!", err.Error())
		return
	}
	files := r.MultipartForm.File["file"]

	if len(files) <= 0 {
		resp.Error("文件不能为空！", nil)
		return
	}

	fileType := req.Type
	var assets []assetsResult
	var fileList map[string]string

	for _, fileHeader := range files {
		
		fileList, err = handleUpload(c, db, fileHeader, fileType)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}
		assets = append(assets, assetsResult{FileName: fileList["fileName"], FilePath: fileList["filePath"], PrevPath: fileList["prevPath"]})
	}

	resp.Success("上传成功！", assets)
	return
}

// 根据文件处理上传逻辑
// 1.判断上传类型，验证后缀合理性 type [0 => "图片" 1 => "视频" 2 => "文件"]
func handleUpload(c *svc.ServiceContext, db *gorm.DB, fileHeader *multipart.FileHeader, fileType string) (result map[string]string, err error) {
	var file multipart.File
	file, err = fileHeader.Open()
	if err != nil {
		return
	}

	var (
		fileSize   int64 = 0
		resolution string
	)

	type Size interface {
		Size() int64
	}

	if sizeInterface, ok := file.(Size); ok {
		fileSize = sizeInterface.Size()
	}

	suffixArr := strings.Split(fileHeader.Filename, ".")

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
		if !iResult {
			return nil, errors.New("【" + suffix + "】不是合法的图片后缀！")
		}
		//img, _, iErr := image.DecodeConfig(file)
		//if iErr != nil {
		//	return nil, iErr
		//}
		//
		//resolution = fmt.Sprintf("%d x %d", img.Width, img.Height)

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
	fileUuid, _ := uuid.NewV4()

	remarkName := fileHeader.Filename
	fileName := util.GetMd5(fileUuid.String() + suffixArr[0])

	fileNameSuffix := fileName + "." + suffix

	uploadPath := temPath + "/" + timeDir + "/"
	filePath := uploadPath + fileNameSuffix
	realpath := path + "/" + filePath

	_, err = os.Stat(path + "/" + uploadPath)

	if err != nil {
		os.MkdirAll(path+"/"+uploadPath, os.ModePerm)
	}

	md5h := md5.New()
	if _, err = io.Copy(md5h, file); err != nil {
		return
	}
	fileMd5 := hex.EncodeToString(md5h.Sum(nil))

	fmt.Println("fileMd5", fileMd5, fileHeader.Filename)

	sha1h := sha1.New()
	if _, err = io.Copy(sha1h, file); err != nil {
		return
	}
	fileSha1 := hex.EncodeToString(sha1h.Sum(nil))

	assets := model.Assets{}
	tx := db.Where("file_md5 = ?", fileMd5).First(&assets)
	if tx.Error != nil {
		if tx.Error != gorm.ErrRecordNotFound {
			err = tx.Error
			return
		}
	}

	if assets.Id > 0 {
		result = make(map[string]string, 0)
		result["fileName"] = assets.FileName
		result["filePath"] = assets.FilePath
		result["prevPath"] = util.FileUrl(assets.FilePath)
		return
	}

	// 上传文件至指定目录
	err = saveUploadedFile(fileHeader, realpath)
	if err != nil {
		return nil, err
	}

	userId, _ := c.Get("userId")
	if userId == nil {
		userId = "0"
	}
	userIdInt, _ := strconv.Atoi(userId.(string))

	fileTypeInt, _ := strconv.Atoi(fileType)
	//保存到数据库
	db.Debug().Create(&model.Assets{
		UserId:     userIdInt,
		FileSize:   fileSize,
		CreateAt:   time.Now().Unix(),
		FileKey:    fileUuid.String(),
		RemarkName: remarkName,
		FileName:   fileNameSuffix,
		FilePath:   filePath,
		Resolution: resolution,
		FileMd5:    fileMd5,
		FileSha1:   fileSha1,
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
