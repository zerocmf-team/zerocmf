/**
** @创建时间: 2020/10/4 9:14 下午
** @作者　　: return
** @描述　　:
 */
package data

import (
	"github.com/gin-gonic/gin"
)

// 定义ws结构体
type WsRouterStruct struct {
	RelativePath string
	Handlers      []gin.HandlerFunc
}

//定义路由结构体
type RouterMapStruct struct {
	RelativePath string
	Handlers     []gin.HandlerFunc
	Method       string
}

// 路由组
type GroupMapStruct struct {
	RelativePath string
	Handlers     []gin.HandlerFunc
}

//定义Template结构体
type TemplateMapStruct struct {
	Theme     string `json:"theme"`
	ThemePath string `json:"themePath"`
	Glob      string `json:"glob"`
	Static    string `json:"static"`
}
