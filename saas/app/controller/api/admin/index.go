/**
** @创建时间: 2021/11/23 09:49
** @作者　　: return
** @描述　　:
 */

package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Index struct{}

func (rest *Index) Get(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code":    "1",
		"message": "获取成功！",
		"data": gin.H{
			"hello": "world",
		},
	})
}
