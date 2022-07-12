package option

import (
	"context"
	"encoding/json"
	"fmt"
	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"
	"zerocmf/service/admin/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadStoreLogic {
	return &UploadStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadStoreLogic) UploadStore(req *types.UploadReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	fmt.Println("req",req)

	resp = new(types.Response)

	c := l.svcCtx
	db := c.Db

	maxFiles := req.MaxFiles
	chunkSize := req.ChunkSize
	imageMaxFileSize := req.FileTypes.Image.UploadMaxFileSize
	imageExtensions := req.FileTypes.Image.Extensions

	videoMaxFileSize := req.FileTypes.Video.UploadMaxFileSize
	videoExtensions := req.FileTypes.Video.Extensions

	audioMaxFileSize := req.FileTypes.Audio.UploadMaxFileSize
	audioExtensions := req.FileTypes.Audio.Extensions

	fileMaxFileSize := req.FileTypes.File.UploadMaxFileSize
	fileExtensions := req.FileTypes.File.Extensions

	uploadSetting := model.UploadSetting{
		MaxFiles:  maxFiles,
		ChunkSize: chunkSize,
		FileTypes: model.FileTypes{
			Image: model.TypeValues{
				UploadMaxFileSize: imageMaxFileSize,
				Extensions:        imageExtensions,
			},
			Video: model.TypeValues{
				UploadMaxFileSize: videoMaxFileSize,
				Extensions:        videoExtensions,
			},
			Audio: model.TypeValues{
				UploadMaxFileSize: audioMaxFileSize,
				Extensions:        audioExtensions,
			},
			File: model.TypeValues{
				UploadMaxFileSize: fileMaxFileSize,
				Extensions:        fileExtensions,
			},
		},
	}

	uploadSettingValue, _ := json.Marshal(uploadSetting)
	db.Model(&model.Option{}).Where("option_name = ?", "upload_setting").Update("option_value", string(uploadSettingValue))
	resp.Success("修改成功！", uploadSetting)
	return
}
