package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"zerocmf/service/user/common"
	"github.com/jinzhu/copier"

	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TokenRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TokenRequestLogic {
	return &TokenRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenRequestLogic) TokenRequest() (resp types.Response) {
	c := l.svcCtx
	r := c.Request
	w := c.ResponseWriter

	r.ParseForm()

	f,_ := json.Marshal(r.Form)

	fmt.Println("form", string(f) )

	conf := c.Config
	inConf := common.Config{}
	copier.Copy(&inConf, &conf)

	oauth := common.NewServer(inConf, "")
	defer oauth.Store.Close()
	srv := oauth.Srv

	err := srv.HandleTokenRequest(w, r)
	if err != nil {
		resp.Error(" token error ",err.Error())
	}
	return
}
