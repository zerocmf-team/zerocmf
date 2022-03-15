/**
** @创建时间: 2021/11/23 12:55
** @作者　　: return
** @描述　　: 解析配置项yml
 */

package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

/**
 * @Author return <1140444693@qq.com>
 * @Description  app结构体
 * @Date 2021/11/23 13:59:43
 * @Param
 * @return
 **/
type app struct {
	Name  string `yaml:"name"`
	Port  string `yaml:"port"`
	Debug bool   `yaml:"debug"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description  数据库结构体
 * @Date 2021/11/23 13:59:43
 * @Param
 * @return
 **/
type database struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Charset  string `yaml:"charset"`
	Prefix   string `yaml:"prefix"`
	AuthCode string `yaml:"authcode"`
}

type config struct {
	App      app      `yaml:"app"`
	Database database `yaml:"database"`
}

var conf *config

func Config() *config {

	if conf == nil {
		c := new(config)
		file, err := ioutil.ReadFile("config/config.yml")
		if err != nil {
			fmt.Println("readfile err", err.Error())
		}
		err = yaml.Unmarshal(file, c)
		if err != nil {
			fmt.Println("err yaml unmarshal", err.Error())
		}

		conf = c
	}

	return conf
}
