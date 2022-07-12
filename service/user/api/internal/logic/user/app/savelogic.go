package app

import (
	"context"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"
	"zerocmf/service/user/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveLogic {
	return &SaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 保存用户信息
 * @Date 2022/3/22 10:28:40
 * @Param
 * @return
 **/

func (l *SaveLogic) Save(req *types.AppSaveReq) (resp types.Response) {
	c := l.svcCtx
	userId, _ := c.Get("userId")
	db := c.Db
	user := model.User{}
	err := user.Show(db, "id = ?", []interface{}{userId})
	if err != nil {
		msg := "系统错误：" + err.Error()
		if err == gorm.ErrRecordNotFound {
			msg = "该管理员账号不存在"
		}
		resp.Error(msg, nil)
		return
	}

	copier.Copy(&user, &req)

	if req.BirthdayTime != "" {
		times, _ := time.Parse("2006-01-02", req.BirthdayTime)
		birthday := times.Unix()
		user.Birthday = birthday
	}

	tx := db.Where("id = ?", userId).Save(&user)

	if tx.Error != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	user.AvatarPrev = util.FileUrl(user.Avatar)
	resp.Success("操作成功！", user)
	return
}
