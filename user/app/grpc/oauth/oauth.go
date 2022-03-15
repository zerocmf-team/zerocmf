/**
** @创建时间: 2021/12/17 14:37
** @作者　　: return
** @描述　　:
 */

package oauth

import (
	"context"
	"gincmf/app/controller/api/common"
)

type Oauth struct {
	UnimplementedOauthServer
}

func (s *Oauth) ValidationBearerToken(ctx context.Context, in *OauthRequest) (*OauthReply, error) {

	token := in.GetToken()
	oauth := new(common.Oauth).NewServer("")

	srv := oauth.Srv
	ti ,err := srv.Manager.LoadAccessToken(context.Background(), token)
	or := OauthReply{}

	if err != nil {
		or.Code = 0
		or.Message = err.Error()
	} else {
		or.Code = 1
		or.Message = "获取成功！"
		or.UserId = ti.GetUserID()
	}


	return &or,nil

}
