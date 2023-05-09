package oauth

import (
	"context"
	"gorm.io/gorm"
	"zerocmf/service/tenant/model"

	"zerocmf/service/tenant/api/internal/svc"
	"zerocmf/service/tenant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurrentUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurrentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurrentUserLogic {
	return &CurrentUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CurrentUserLogic) CurrentUser() (resp types.Response) {
	c := l.svcCtx
	userId, _ := c.Get("userId")
	db := c.Db
	user := model.User{}
	tx := db.Where("id = ?", userId).First(&user)
	if tx.Error != nil {
		msg := "系统错误：" + tx.Error.Error()
		if tx.Error == gorm.ErrRecordNotFound {
			msg = "该管理员账号不存在"
		}
		resp.Error(msg, nil)
		return
	}
	resp.Success("获取成功！", user)
	return
}
