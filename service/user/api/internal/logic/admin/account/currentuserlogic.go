package account

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"
	"zerocmf/service/user/model"
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

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取当前用户信息
 * @Date 2022/3/22 9:47:25
 * @Param
 * @return
 **/

func (l *CurrentUserLogic) CurrentUser() (resp types.Response) {
	c := l.svcCtx
	userId, _ := c.Get("userId")
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))

	user := model.User{}
	tx := db.Where("id = ? and user_type = 1", userId).First(&user)

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
