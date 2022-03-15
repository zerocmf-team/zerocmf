/**
** @创建时间: 2021/12/26 21:56
** @作者　　: return
** @描述　　:
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/casbin"
	"github.com/gincmf/bootstrap/controller"
	"github.com/gincmf/bootstrap/model"
	"github.com/gincmf/bootstrap/util"
	pathToRegexp "github.com/soongo/path-to-regexp"
	"net/http"
	"strings"
)

/**
 * @Author return <1140444693@qq.com>
 * @Description 验证rbac中间件
 * @Date 2021/12/26 21:56:49
 * @Param
 * @return
 **/

func Rbac(c *gin.Context) {

	userId, _ := c.Get("userId")
	requestURI := c.Request.RequestURI
	var menusApi []model.AdminMenuApi
	db := util.GetDb(c)
	tx := db.Find(&menusApi)
	if util.IsDbErr(tx) != nil {
		new(controller.Rest).Error(c, tx.Error.Error(), nil)
		c.Abort()
		return
	}

	path := strings.Split(requestURI, "?")[0]

	var object string
	for _, v := range menusApi {
		match := pathToRegexp.MustMatch(v.Path, &pathToRegexp.Options{Decode: func(str string, token interface{}) (string, error) {
			return pathToRegexp.DecodeURIComponent(str)
		}})
		matchRes, err := match(path)
		if err != nil {
			panic(err.Error())
		}
		if matchRes != nil {
			object = v.Object
			break
		}
	}

	if object != "" {
		// 判断当前接口是否有权限
		e, err := casbin.NewEnforcer("")
		if err != nil {
			new(controller.Rest).Error(c, err.Error(), nil)
			c.Abort()
			return
		}

		access ,err := e.Enforce(userId,object,"*")

		if err != nil {
			new(controller.Rest).Error(c, err.Error(), nil)
			c.Abort()
			return
		}

		rolePolicies := e.GetFilteredPolicy(0, userId.(string))
		if userId == "1" || len(rolePolicies) == 0 {
			access = true
		}

		if !access {
			c.JSON(http.StatusOK,gin.H{
				"code":"403",
				"message":"该接口无权访问",
			})
			c.Abort()
			return
		}
	}

	c.Next()
}
