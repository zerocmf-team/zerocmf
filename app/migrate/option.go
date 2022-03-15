package migrate

import (
	"encoding/json"
	"fmt"
	"gincmf/app/model"
	cmf "github.com/gincmf/cmf/bootstrap"
)

type option struct {
	Migrate
}

func (m *option) AutoMigrate() {
	cmf.Db.AutoMigrate(&model.Option{})
	siteResult := cmf.Db.First(&model.Option{}, "option_name = ?", "site_info") // 查询
	if siteResult.Error != nil {
		//初始化默认json
		siteInfo := &model.SiteInfo{}
		siteInfoValue, _ := json.Marshal(siteInfo)
		cmf.Db.Create(&model.Option{AutoLoad: 1, OptionName: "site_info", OptionValue: string(siteInfoValue)})
	}

	uploadResult := cmf.Db.First(&model.Option{}, "option_name = ?", "upload_setting") // 查询
	if uploadResult.Error != nil {
		//初始化默认json
		uploadSetting := &model.UploadSetting{
			MaxFiles:  20,
			ChunkSize: 512,
			FileTypes: model.FileTypes{
				Image: model.TypeValues{
					UploadMaxFileSize: 10240,
					Extensions:        "jpg,jpeg,png,gif,bmp4,svg",
				},
				Video: model.TypeValues{
					UploadMaxFileSize: 102400,
					Extensions:        "mp4,avi,wmv,rm,rmvb,mkv",
				},
				Audio: model.TypeValues{
					UploadMaxFileSize: 10240,
					Extensions:        "mp3,wma,wav",
				},
				File: model.TypeValues{
					UploadMaxFileSize: 10240,
					Extensions:        "txt,pdf,doc,docx,xls,xlsx,ppt,pptx,zip,rar",
				},
			},
		}
		uploadSettingValue, _ := json.Marshal(uploadSetting)
		fmt.Println("uploadSettingValue", string(uploadSettingValue))
		cmf.Db.Create(&model.Option{AutoLoad: 1, OptionName: "upload_setting", OptionValue: string(uploadSettingValue)})
	}
}
