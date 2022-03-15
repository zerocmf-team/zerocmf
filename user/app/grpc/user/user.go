/**
** @创建时间: 2022/2/24 13:07
** @作者　　: return
** @描述　　:
 */

package user

import (
	"context"
	"gincmf/app/model"
	"github.com/gincmf/bootstrap/util"
	"github.com/jinzhu/copier"
)

type User struct {
	UnimplementedUserServer
}

func (s *User) Get(ctx context.Context, in *UserRequest) (userReply *UserReply, err error) {

	userId := in.GetUserId()
	tenantId := in.GetTenant()

	db := util.ManualDb(tenantId)

	// 执行services逻辑 根据id查询当前用户信息
	user := model.User{}
	var query = "id = ? AND delete_at = 0"
	var queryArgs = []interface{}{userId}

	tx := db.Where(query, queryArgs...).First(&user)

	if util.IsDbErr(tx) != nil {
		userReply = &UserReply{
			Code:    0,
			Message: "请求失败！",
		}
		err = tx.Error
		return
	}

	if user.Id == 0 {

		userReply = &UserReply{
			Code:    0,
			Message: "该用户不存在！",
		}
		return

	}

	data := Data{}
	copier.Copy(&data, &user)

	userReply = &UserReply{
		Code:    1,
		Message: "获取成功！",
		Data: &data,
	}

	return

}
