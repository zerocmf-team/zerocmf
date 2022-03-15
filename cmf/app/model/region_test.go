/**
** @创建时间: 2021/11/27 08:53
** @作者　　: return
** @描述　　:
 */

package model

import (
	"fmt"
	"github.com/gincmf/bootstrap/config"
	"github.com/gincmf/bootstrap/db"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_Region(t *testing.T) {
	db := db.Db()
	f, err := os.Open("data/region.sql")
	if err != nil {
		fmt.Println("err", err)
	}
	bytes, _ := ioutil.ReadAll(f)
	result := string(bytes)
	conf := config.Config()
	prefix := conf.Database.Prefix
	result = strings.ReplaceAll(result, "{prefix}", prefix)
	// fmt.Println(result)
	sqlArr := strings.Split(result, ";")
	for _, sql := range sqlArr {
		db.Debug().Exec(sql)
	}
}
