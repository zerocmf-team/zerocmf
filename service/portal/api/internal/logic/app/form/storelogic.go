package form

import (
	"context"
	"github.com/jinzhu/copier"
	"time"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

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

func (l *StoreLogic) Store(req *types.FormSubmitReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))

	form := model.Form{}
	tx := db.Where("id", req.FormId).First(&form)
	if util.IsDbErr(tx) != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	if form.Id == 0 {
		resp.Error("表单不存在", nil)
		return
	}

	item := model.FormItem{}
	copier.Copy(&item, &req)
	item.CreateAt = time.Now().Unix()
	item.UpdateAt = time.Now().Unix()
	item.Status = 1
	db.Create(&item)
	resp.Success("操作成功！", item)
	return
}
