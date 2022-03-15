package view

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Template struct {
	c    *gin.Context
	Name string
	Obj  map[string]interface{}
}

func (t Template) Assign(k string, i interface{}) Template {
	if t.Obj == nil {
		t.Obj = make(map[string]interface{})
	}
	t.Obj[k] = i
	return t
}

func (t Template) GetView(c *gin.Context) Template {
	cmfView, _ := c.Get("template")
	iView := Template{}
	switch cmfView.(type) {
	case Template:
		iView = cmfView.(Template)
	}
	iView.c = c
	return iView
}

//渲染方法
func (t *Template) Fetch(name string) {
	t.c.HTML(http.StatusOK, name, t.Obj)
}

func (t *Template) Error(error string) {
	t.Obj["error"] = error
	t.c.HTML(http.StatusOK, "error.html", t.Obj)
}
