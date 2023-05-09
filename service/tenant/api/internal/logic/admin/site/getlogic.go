package site

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/tenant/api/internal/svc"
	"zerocmf/service/tenant/api/internal/types"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.SiteGetReq) (resp types.Response) {

	c := l.svcCtx
	r := c.Request
	db := c.Db
	userId, _ := c.Get("userId")

	prefix := c.Config.Database.Prefix

	var result []struct {
		SiteId     int64   `gorm:"type:bigint(20);comment;站点唯一编号" json:"siteId"`
		Name       string  `gorm:"type:varchar(32);comment:站点名称" json:"name"`
		Desc       string  `gorm:"type:varchar(255);comment:站点描述" json:"desc"`
		Status     int     `gorm:"type:tinyint(3);default:1;comment:文件状态" json:"status"`
		Oid        int64   `gorm:"type:bigint(20);comment:真实站点用户id;not null" json:"oid"`
		IsOwner    int     `gorm:"type:tinyint(3);comment:是否为站点拥有者;not null" json:"isOwner"`
		CreateAt   int64   `gorm:"type:bigint(20);NOT NULL" json:"createAt"`
		CreateTime string  `gorm:"-" json:"createTime"`
		ListOrder  float64 `gorm:"type:float;default:10000;comment:排序（越大越靠前）" json:"listOrder" label:"排序"`
	}

	current, pageSize, err := data.NewPaginate(r).Default()

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	var total int64 = 0

	tx := db.Select("s.site_id,s.name,s.desc,s.status,s.create_at,su.oid,su.is_owner").Table(prefix+"site s").Joins("left join "+prefix+"site_user su on s.site_id = su.site_id").
		Joins("inner join "+prefix+"user u on u.uid = su.uid").
		Where("u.uid = ? AND s.delete_at = ?", userId, 0).Scan(&result).Count(&total)

	if tx.Error != nil {
		err = tx.Error
		return
	}

	//获取当用户信息
	tx = db.Select("s.site_id,s.name,s.desc,s.status,s.create_at,su.oid,su.is_owner,su.list_order").Table(prefix+"site s").Joins("left join "+prefix+"site_user su on s.site_id = su.site_id").
		Joins("inner join "+prefix+"user u on u.uid = su.uid").
		Where("u.uid = ? AND s.delete_at = ?", userId, 0).Offset((current - 1) * pageSize).Order("su.list_order desc").Scan(&result)

	if tx.Error != nil {
		resp.Error("操作失败！", tx.Error.Error())
		return
	}
	pageData := data.Paginate{Data: result, Current: current, PageSize: pageSize, Total: total}
	resp.Success("操作成功！", pageData)
	return
}
