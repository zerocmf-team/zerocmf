package account

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"zerocmf/service/tenant/rpc/tenantclient"
	"zerocmf/service/user/model"

	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"

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

func (l *DeleteLogic) Delete(req *types.OneReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")

	siteIdInt, _ := strconv.ParseInt(siteId.(string), 10, 64)

	db := c.Config.Database.ManualDb(siteId.(string))

	tenantRpc := c.TenantRpc

	id := req.Id

	if id == "1" {
		resp.Error("root用户不能被删除！", nil)
		return
	}

	user := model.User{}
	tx := db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			resp.Error("该用户不存在或已被删除！", nil)
			return
		}
		resp.Error("系统错误", nil)
		return
	}

	tx = db.Where("id = ?", id).Delete(&model.User{})
	if tx.Error != nil {
		resp.Error("系统错误", nil)
		return
	}

	_, err := tenantRpc.RemoveSiteUser(l.ctx, &tenantclient.RemoveSiteUserReq{
		SiteId: siteIdInt,
		Mobile: user.UserLogin,
	})
	if err != nil {
		resp.Error("系统错误", err.Error())
		return
	}

	resp.Success("删除成功！", user)
	return
}
