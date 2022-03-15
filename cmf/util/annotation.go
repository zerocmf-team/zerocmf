/**
** @创建时间: 2020/7/16 4:38 下午
** @作者　　: return
 */
package util

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"gorm.io/gorm"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)
var db *gorm.DB
/**
** @创建时间: 2020/7/16 4:49 下午
** @作者　　: return
** @描述　　: 解析注解生成api到数据库
 */
func AnParseDir(path string,d ...*gorm.DB) {
	if d != nil {
		db = d[0]
	}
	list, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	for _, d := range list {
		dName := d.Name()
		path = strings.TrimRight(path, "/") + "/"
		fullPath := path + dName
		if d.IsDir() {
			AnParseDir(fullPath)
		} else {
			if strings.HasSuffix(fullPath, ".go") {
				filename := filepath.Join(path, dName)
				astInspect(filename)
			}
		}
	}
}

func astInspect(filename string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("err", err.Error())
	}

	var s string
	// 检查AST并打印所有标识符和文字。
	ast.Inspect(f, func(n ast.Node) bool {
		astc, ok := n.(*ast.Comment)
		if !ok {
			return true
		}
		s = astc.Text
		if s != "" {
			annotation(s)
		}
		return true
	})
}

func annotation(contents string) {

	rm := regexp.MustCompile(`@restApi\(([\s\S]*)\)`)
	contents = rm.FindString(contents)
	if contents != "" {
		contents = regExp(`(\*/)+|(/\*)+|(\*)+|[ ,'"\t]+`, contents, "")
		resultContents := regExp(`(\s)+`, contents, ",")
		arr := strings.Split(resultContents, ",")
		type apiList struct {
			Id     int    `json:"id"`
			Name   string `gorm:"type:varchar(30);comment:'名称'" json:"name"`
			Desc   string `gorm:"type:varchar(255);comment:'描述'" json:"desc"`
			Url    string `gorm:"type:varchar(255);comment:'接口url'" json:"url"`
			Param  string `gorm:"type:varchar(30);comment:'参数'" json:"param"`
			Method string `gorm:"type:varchar(30);comment:'请求类型'" json:"method"`
			Status int    `gorm:"type:tinyint(3);comment:'状态';default:1" json:"status"`
		}

		api := apiList{}
		for _, v := range arr {
			if v != "" {
				itemList := strings.Split(v, "=>")
				switch itemList[0] {
				case "name":
					api.Name = itemList[1]
					break
				case "desc":
					api.Desc = itemList[1]
					break
				case "url":
					api.Url = itemList[1]
					break
				case "param":
					api.Param = itemList[1]
					break
				case "method":
					api.Method = itemList[1]
					break
				case "status":
					status,_ := strconv.Atoi(itemList[1])
					api.Status = status
					break
				default:
					break

				}
			}
		}

		apiModel := apiList{
			Name:   api.Name,
			Desc:   api.Desc,
			Url:    api.Url,
			Param:  api.Param,
			Method: api.Method,
			Status: api.Status,
		}
		db.Where(apiList{Name: api.Name,Method: api.Method}).FirstOrCreate(&apiModel)

	}
}

func regExp(regExpStr string, content string, repl string) string {
	reg := regexp.MustCompile(regExpStr)
	contents := reg.ReplaceAllString(content, repl)
	return contents
}

