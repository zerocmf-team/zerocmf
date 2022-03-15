/**
** @创建时间: 2021/11/23 09:49
** @作者　　: return
** @描述　　:
 */

package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
)

type Index struct {
	controller.Rest
}

func (rest *Index) Get(c *gin.Context) {
	rest.Success(c, "获取成功！", gin.H{
		"version":   "v1",
		"create_at": "2021/11/24 18:59:00",
		"author":    "gincmf-team",
	})
}
