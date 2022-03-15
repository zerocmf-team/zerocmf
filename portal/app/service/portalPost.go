/**
** @创建时间: 2021/12/9 13:48
** @作者　　: return
** @描述　　:
 */

package service

import (
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/paginate"
	"github.com/gincmf/bootstrap/util"
)

type PortalPost struct {
}

func (service *PortalPost) IndexByCategory(c *gin.Context, query string, queryArgs []interface{}) (data paginate.Paginate, err error) {
	db := util.GetDb(c)
	current, pageSize, err := new(paginate.Paginate).Default(c)
	if err != nil {

	}
	data, err = new(model.PortalPost).IndexByCategory(db, current, pageSize, query, queryArgs,nil)
	return
}
