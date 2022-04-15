package oauth

import (
	"context"
	"gincmf/common/bootstrap/util"
	"gincmf/service/user/common"
	"gincmf/service/user/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"strconv"
	"time"

	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/go-oauth2/oauth2/v4/server"
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

func (l *TokenLogic) Token(req *types.TokenReq) (resp *types.Response, err error) {

	resp = new(types.Response)

	username := req.Usermame
	password := req.Password

	c := l.svcCtx
	r := c.Request
	db := c.Db

	// 验证用户账号密码
	user := model.User{}
	tx := db.Where("user_login = ?", username).First(&user)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		resp.Error("查询用户失败", nil)
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

	tokenExp := "168"

	exp, err := strconv.Atoi(tokenExp)

	if err != nil {
		resp.Error("失效时间应该是整数，单位为小时！", nil)
		return
	}

	conf := c.Config
	inConf := common.Config{}
	copier.Copy(&inConf, &conf)

	oauth := common.NewServer(inConf, "")
	defer oauth.Store.Close()
	srv := oauth.Srv
	oauthConfig := oauth.Config

	duration := time.Duration(exp) * time.Hour

	authReq := &server.AuthorizeRequest{
		RedirectURI:    oauthConfig.RedirectURL,
		ResponseType:   "code",
		ClientID:       oauthConfig.ClientID,
		State:          "jwt",
		Scope:          "all",
		UserID:         strconv.Itoa(user.Id),
		AccessTokenExp: duration,
		Request:        c.Request,
	}

	ti, err := srv.GetAuthorizeToken(l.ctx, authReq)

	if err != nil {
		resp.Error("系统错误："+err.Error(), nil)
		return
	}

	code := ti.GetCode()

	token, err := oauthConfig.Exchange(context.Background(), code)

	if err != nil {
		resp.Error("获取失败："+err.Error(), nil)
		return
	}

	resp.Success("获取成功！", token)
	return
}
