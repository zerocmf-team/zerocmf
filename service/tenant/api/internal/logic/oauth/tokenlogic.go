package oauth

import (
	"context"
	"gorm.io/gorm"
	"strconv"
	"time"
	"zerocmf/common/bootstrap/apisix"
	"zerocmf/common/bootstrap/apisix/plugins/authentication"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/tenant/model"
	"zerocmf/service/user/rpc/userclient"

	"zerocmf/service/tenant/api/internal/svc"
	"zerocmf/service/tenant/api/internal/types"

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

	c := l.svcCtx
	r := c.Request
	conf := c.Config

	loginType := req.LoginType

	// 如果是手机登录
	if loginType == "ram" {

		siteId := req.SiteId

		if siteId == "" {
			resp.Error("主站点不能为空", nil)
			return
		}

		userLogin := req.UserLogin
		password := req.Password

		if userLogin == "" {
			resp.Error("用户名不能为空", nil)
			return
		}
		if password == "" {
			resp.Error("密码不能为空", nil)
			return
		}

		userRpc := c.UserRpc
		userReply, err := userRpc.RamLogin(l.ctx, &userclient.LoginReq{
			SiteId:    siteId,
			UserLogin: userLogin,
			UserPass:  util.GetMd5(password),
		})

		if err != nil {
			resp.Error("查询用户失败", err.Error())
			return
		}

		if userReply.ErrorMsg != "" {
			resp.Error(userReply.ErrorMsg, err.Error())
			return
		}

		var exp int64 = 86400
		userId := strconv.FormatInt(userReply.Id, 10)
		key := siteId + "_" + userId
		apisix.NewConsumer(conf.Apisix.ApiKey, conf.Apisix.Host).Add(key, apisix.WithJwtAuth(authentication.JwtAuth{Key: key, Exp: exp}))

		token, tokenErr := apisix.NewJwt(conf.Apisix.Host).GetAuthorizeToken(key, map[string]string{
			"type":   "ram",
			"siteId": siteId,
			"userId": userId,
		})

		if tokenErr != nil {
			resp.Error(tokenErr.Error(), nil)
			return
		}

		resp.Success("获取成功！", token)

	} else {
		mobile := req.UserLogin
		password := req.Password

		if mobile == "" {
			resp.Error("用户名不能为空", nil)
			return
		}

		if password == "" {
			resp.Error("密码不能为空", nil)
			return
		}

		db := c.Db

		// 验证用户账号密码
		user := model.User{}
		tx := db.Where("mobile = ?", mobile).First(&user)
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
		userId := strconv.FormatInt(user.Uid, 10)
		apisix.NewConsumer(conf.Apisix.ApiKey, conf.Apisix.Host).Add(userId, apisix.WithJwtAuth(authentication.JwtAuth{Key: userId, Exp: exp}))

		token, tokenErr := apisix.NewJwt(conf.Apisix.Host).GetAuthorizeToken(userId, nil)
		if tokenErr != nil {
			resp.Error(tokenErr.Error(), nil)
			return
		}

		resp.Success("获取成功！", token)
	}
	return
}
