package site

import (
	"context"
	"github.com/jinzhu/copier"
	"time"
	"zerocmf/common/bootstrap/redis"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/admin/rpc/adminclient"
	"zerocmf/service/portal/rpc/portalclient"
	"zerocmf/service/tenant/model"
	"zerocmf/service/user/rpc/userclient"

	"zerocmf/service/tenant/api/internal/svc"
	"zerocmf/service/tenant/api/internal/types"

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

func (l *StoreLogic) Store(req *types.SiteSaveReq) (resp types.Response) {
	c := l.svcCtx
	context := l.ctx
	resp = save(c, context, req)
	return
}

func save(c *svc.ServiceContext, context context.Context, req *types.SiteSaveReq) (resp types.Response) {
	db := c.Db
	userId, _ := c.Get("userId")

	user := model.User{}
	tx := db.Where("id", userId).First(&user)
	if tx.Error != nil {
		resp.Error("操作失败！", tx.Error.Error())
		return
	}

	site := model.Site{
		Status: 1,
	}

	if req.SiteId == 0 {
		copier.Copy(&site, &req)
		site.CreateAt = time.Now().Unix()
		site.UpdateAt = time.Now().Unix()

		key := "tenant:site"

		siteId, err := redis.EncryptUid(key, 0)
		if err != nil {
			resp.Error("操作失败！", err.Error())
			return
		}

		site.SiteId = siteId
		tx = db.Create(&site)
		if tx.Error != nil {
			resp.Error("操作失败！", tx.Error.Error())
			return
		}

		//	为站点关联创建者身份
		siteUser := model.SiteUser{
			TenantId: 0,
			SiteId:   siteId,
			Uid:      user.Uid,
			Oid:      1, // 初始账号为1的超级管理员
			IsOwner:  1,
			Status:   1,
		}
		tx = db.Create(&siteUser)
		if tx.Error != nil {
			resp.Error("操作失败！", tx.Error.Error())
			return
		}

		/* 为站点生成数据库
		 ** rpc调用
		 ** 生成规则
		 *** 站点数据库唯一，根据站点id来命名，一个站点对应一个数据库集合
		 */

		// todo 捕获数据库迁移异常
		c.AdminRpc.AutoMigrate(context, &adminclient.SiteReq{
			SiteId: siteId,
		})

		c.UserRpc.AutoMigrate(context, &userclient.SiteReq{
			SiteId: siteId,
		})

		c.PortalRpc.AutoMigrate(context, &portalclient.SiteReq{
			SiteId: siteId,
		})

	} else {
		tx = db.Where("site_id", req.SiteId).First(&site)
		if util.IsDbErr(tx) != nil {
			resp.Error("操作失败！", tx.Error.Error())
			return
		}
		if tx.RowsAffected == 0 {
			resp.Error("操作失败！该站点不存在", tx.Error.Error())
			return
		}
		copier.Copy(&site, &req)
		site.UpdateAt = time.Now().Unix()
		tx = db.Where("id", site.Id).Save(&site)
		if tx.Error != nil {
			resp.Error("操作失败！", tx.Error.Error())
			return
		}
	}
	resp.Success("操作成功！", site)
	return
}
