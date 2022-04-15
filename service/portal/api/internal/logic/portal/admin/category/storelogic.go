package category

import (
	"context"
	comModel "gincmf/common/bootstrap/model"
	"gincmf/service/portal/model"
	"github.com/jinzhu/copier"
	"strconv"
	"strings"
	"unicode/utf8"

	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line
	c := l.svcCtx
	resp = Save(c, req)
	return
}

func Save(c *svc.ServiceContext, req *types.CateSaveReq) (resp types.Response) {

	db := c.Db
	name := req.Name
	parentId := req.ParentId
	editId := req.Id
	portalCategory := model.PortalCategory{
		ParentId: parentId,
		Name:     name,
	}

	msg := "新增成功！"
	if editId == 0 {
		portalCategory.Id = editId
	}else {
		msg = "更新成功！"
		pCate := model.PortalCategory{}
		tx := db.Where("id = ?",editId).First(&pCate)

		if tx.Error != nil {
			resp.Error(tx.Error.Error(),nil)
			return
		}

		// 新的父级不能小于原父级
		if pCate.ParentId < parentId {
			resp.Error("非法父级",nil)
			return
		}

	}
	copier.Copy(&portalCategory, &req)
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
