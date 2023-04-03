package theme

import (
	"context"
	"encoding/json"
	"time"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitLogic {
	return &InitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitLogic) Init(req *types.InitReq) (resp types.Response) {

	c := l.svcCtx
	db := c.Db

	theme := model.Theme{
		Name:      req.Theme,
		Version:   req.Version,
		Thumbnail: req.Thumbnail,
		CreateAt:  time.Now().Unix(),
	}

	queryTheme := model.Theme{}
	tx := db.Where("name = ?", req.Theme).First(&queryTheme)
	if util.IsDbErr(tx) != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}
	id := queryTheme.Id
	if id == 0 {
		db.Create(&theme)
	} else {
		theme.Id = id
		theme.UpdateAt = time.Now().Unix()
		db.Save(&theme)
	}
	var addThemeFile []model.ThemeFile
	for _, v := range req.ThemeFile {

		var file struct {
			Name     string `json:"name"`
			File     string `json:"file"`
			Type     string `json:"type"`
			Desc     string `json:"desc"`
			IsPublic int    `json:"is_public"`
		}

		json.Unmarshal([]byte(v), &file)

		queryFile := model.ThemeFile{}
		tx = db.Where("theme = ? AND file = ?", req.Theme, file.File).First(&queryFile)
		if util.IsDbErr(tx) != nil {
			resp.Error(tx.Error.Error(), nil)
			return
		}

		// 合并json
		var m1, m2 map[string]interface{}
		json.Unmarshal([]byte(v), &m1)              // 源文件
		json.Unmarshal([]byte(queryFile.More), &m2) // 数据库
		merged := util.JsonMerge(m1, m2)
		bytes, err := json.Marshal(merged)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}
		more := string(bytes)
		themeFile := model.ThemeFile{
			Theme:       req.Theme,
			IsPublic:    file.IsPublic,
			Name:        file.Name,
			File:        file.File,
			Type:        file.Type,
			Description: file.Desc,
			More:        more,
			ConfigMore:  v,
			CreateAt:    time.Now().Unix(),
		}

		id = queryFile.Id
		if id == 0 {
			addThemeFile = append(addThemeFile, themeFile)
		} else {
			themeFile.Id = id
			themeFile.UpdateAt = time.Now().Unix()
			tx = db.Save(&themeFile)
			if tx.Error != nil {
				resp.Error(tx.Error.Error(), nil)
				return
			}
		}
	}
	if len(addThemeFile) > 0 {
		tx = db.Create(&addThemeFile)
		if util.IsDbErr(tx) != nil {
			resp.Error(tx.Error.Error(), nil)
			return
		}
	}

	option := model.Option{
		OptionName:  "theme",
		OptionValue: req.Theme,
	}
	queryOption := model.Option{}
	tx = db.Where("option_name", "theme").First(&queryOption)
	if util.IsDbErr(tx) != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	if queryOption.Id == 0 {
		tx = db.Create(&option)
	} else {
		option.Id = queryOption.Id
		tx = db.Save(&option)
	}

	if tx.Error != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	resp.Success("操作成功！", req)

	return
}
