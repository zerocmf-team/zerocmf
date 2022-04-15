package app

import (
	"context"
	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"
	"gincmf/service/user/model"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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

func (l *CurrentUserLogic) CurrentUser() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	c := l.svcCtx
	userId, _ := c.Get("userId")

	db := c.Db

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

	resp.Success( "获取成功！", user)
	return
}
