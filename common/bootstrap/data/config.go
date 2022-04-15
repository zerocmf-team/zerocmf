/**
** @创建时间: 2022/3/13 16:55
** @作者　　: return
** @描述　　:
 */

package data

import (
	"gincmf/common/bootstrap/db"
	"github.com/jinzhu/copier"
)

var (
	domain string
	salts  string
)

type config struct {
	Database struct {
		Type     string
		Host     string
		Database string
		Username string
		Password string
		Port     int
		Charset  string
		Prefix   string
		AuthCode string
	}
}

var conf *config

func InitConfig(c *db.Database) {
	if conf == nil {
		conf = new(config)
	}
	copier.Copy(&conf.Database,&c)
}

func Config() *config {
	return conf
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 设置域名
 * @Date 2022/3/13 16:57:18
 * @Param
 * @return
 **/

func SetDomain(value string) {
	domain = value
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取域名
 * @Date 2022/3/13 16:57:27
 * @Param
 * @return
 **/

func Domain() string {
	return domain
}

func SetSalts(value string) {
	salts = value
}

func Salts() string {
	return salts
}
