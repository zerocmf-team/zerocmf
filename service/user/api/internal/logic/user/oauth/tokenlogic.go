package oauth

import (
	"context"
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/apisix"
	"zerocmf/common/bootstrap/apisix/plugins/authentication"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/user/model"

	"strconv"
	"time"

	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TokenLogic {
	return &TokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenLogic) Token(req *types.TokenReq) (resp types.Response) {

	username := req.Usermame
	password := req.Password

	if username == "" {
		resp.Error("用户名不能为空", nil)
		return
	}

	if password == "" {
		resp.Error("密码不能为空", nil)
		return
	}

	c := l.svcCtx
	r := c.Request
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))
	conf := c.Config

	// 验证用户账号密码
	user := model.User{}
	tx := db.Where("user_login = ?", username).First(&user)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			resp.Error("查询用户失败", nil)
			return
		}
		resp.Error("数据库出错", tx.Error.Error())
		return
	}

	//验证密码
	if util.GetMd5(password) != user.UserPass {
		resp.Error("账号密码不正确！", nil)
		return
	}

	// 更新最后登录记录
	u := model.User{
		LastLoginIp: r.RemoteAddr,
		LastLoginAt: time.Now().Unix(),
	}
	db.Where("id = ?", user.Id).Updates(&u)

	var exp int64 = 86400
	userId := strconv.Itoa(user.Id)
	apisix.NewConsumer(conf.Apisix.ApiKey, conf.Apisix.Host).Add(userId, apisix.WithJwtAuth(authentication.JwtAuth{Key: userId, Exp: exp}))

	token, tokenErr := apisix.NewJwt(conf.Apisix.ApiKey, conf.Apisix.Host).GetAuthorizeToken(userId)
	if tokenErr != nil {
		resp.Error(tokenErr.Error(), nil)
		return
	}

	resp.Success("获取成功！", token)
	return
}
