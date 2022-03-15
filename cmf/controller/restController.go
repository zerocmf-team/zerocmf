package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/model"
	"net/http"
)

type RestInterface interface {
	Get(c *gin.Context)
	Show(c *gin.Context)
	Edit(c *gin.Context)
	Store(c *gin.Context)
	Delete(c *gin.Context)
}

type Rest struct{}

func (r Rest) Forbidden(c *gin.Context) {
	c.String(http.StatusNotFound, "页面不存在！")
}

func (r Rest) Success(c *gin.Context, msg string, data interface{}) {
	var result model.ReturnData
	result = model.ReturnData{Code: 1, Msg: msg, Data: data}
	c.JSON(http.StatusOK, result)
}

func (r Rest) Error(c *gin.Context, msg string, data interface{}) {
	var result model.ReturnData
	result = model.ReturnData{Code: 0,Msg: msg, Data: data}
	c.JSON(http.StatusOK, result)
}
func (r Rest) ErrorCode(c *gin.Context,code int, msg string, data interface{}) {
	var result model.ReturnData
	result = model.ReturnData{Code: code,Msg: msg, Data: data}
	c.JSON(http.StatusOK, result)
}

func (r Rest) JsonSuccess(msg string, data interface{}) string {

	var result model.ReturnData
	result = model.ReturnData{Code:1,Msg: msg, Data: data}
	bytes, _ := json.Marshal(result)
	return string(bytes)

}

func (r Rest) JsonError(msg string, data interface{}) string {

	var result model.ReturnData
	result = model.ReturnData{Code:0,Msg: msg, Data: data}
	bytes, _ := json.Marshal(result)
	return string(bytes)

}
