package middleware

import (
	"fmt"
	"gincmf/app/controller/common"
	"gincmf/app/util"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/controller"
	"gopkg.in/oauth2.v3"
	"net/http"
)

var Token oauth2.TokenInfo

//验证oauth token值
func ValidationBearerToken(c *gin.Context) {
	s := common.Srv
	t, err := s.ValidationBearerToken(c.Request)
	Token = t
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code":400,"msg":err.Error()})
		fmt.Println("err", err.Error())
		c.Abort()
		return
	}
	c.Set("user_id", t.GetUserID())
	c.Next()
}

//验证是否为管理员
func ValidationAdmin(c *gin.Context) {
	currentUser := util.CurrentUser(c)
	userType := currentUser.UserType
	if userType != 1 {
		controller.RestController{}.Error(c, "您不是管理员，无权访问！",nil)
		c.Abort()
		return
	}
	c.Next()
}
