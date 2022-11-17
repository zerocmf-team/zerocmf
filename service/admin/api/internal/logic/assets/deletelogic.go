package assets

import (
	"context"
	"fmt"
	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"
	"zerocmf/service/admin/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DeleteReq) (resp *types.Response) {
	resp = new(types.Response)

	c := l.svcCtx
	db := c.Db

	id := req.Id
	ids := req.Ids

	assets := model.Assets{}

	if len(ids) == 0 {

		tx := db.Where("id", id).First(&assets)
		if tx.Error != nil {
			resp.Error("服务器错误："+tx.Error.Error(), nil)
			return
		}

		assets.Id = id
		assets.Status = 0

		if err := db.Save(assets).Error; err != nil {
			resp.Error("删除失败！", nil)
			return
		}
	} else {
		fmt.Println("ids", ids)
		if err := db.Model(&assets).Where("id IN (?)", ids).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
			resp.Error("删除失败！", nil)
			return
		}
	}

	resp.Success("删除成功！", assets)

	return
}
