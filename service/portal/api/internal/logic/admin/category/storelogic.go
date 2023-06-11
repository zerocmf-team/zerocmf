package category

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"strings"
	"unicode/utf8"
	comModel "zerocmf/common/bootstrap/model"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"
)

type StoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreLogic {
	return &StoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreLogic) Store(req *types.CateSaveReq) (resp types.Response) {
	c := l.svcCtx
	resp = Save(c, req)
	return
}

func Save(c *svc.ServiceContext, req *types.CateSaveReq) (resp types.Response) {

	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))
	editId := req.Id
	portalCategory := model.PortalCategory{}

	msg := "新增成功！"
	if editId == 0 {
		portalCategory.Id = editId
	} else {
		msg = "更新成功！"

		tx := db.Where("id = ?", editId).First(&portalCategory)

		if tx.Error != nil {
			resp.Error(tx.Error.Error(), nil)
			return
		}

		// 新的父级不能等于自己
		parentId := req.ParentId
		if portalCategory.Id == parentId {
			resp.Error("非法父级", nil)
			return
		}
	}
	copier.Copy(&portalCategory, req)
	portalCategory.Status = req.Status

	data, err := portalCategory.Save(db)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	alias := portalCategory.Alias
	if alias != "" {

		if strings.HasPrefix(alias, "/") {
			aliasRune := []rune(alias)
			sLen := utf8.RuneCountInString(alias)
			alias = string(aliasRune[1:sLen])
		}
		fullUrl := "list/" + strconv.Itoa(portalCategory.Id)
		url := alias
		// 插入别名
		route := comModel.Route{
			Type:    1,
			FullUrl: fullUrl,
			Url:     url,
		}
		err = route.Set(db)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}
	}
	resp.Success(msg, data)
	return
}
