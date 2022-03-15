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
	"os"
	"strconv"
)

/**
 * @Author return <1140444693@qq.com>
 * @Description  app结构体
 * @Date 2021/11/23 13:59:43
 * @Param
 * @return
 **/
type app struct {
	Name   string `yaml:"name"`   // 项目名称
	Type   string `yaml:"type"`   // 项目类型 single：单体应用，grpc：微服务应用
	Port   int `yaml:"port"`   // 项目端口号
	Debug  bool   `yaml:"debug"`  // debug模式
	Domain string `yaml:"domain"` // 项目域名 不填默认走上下文
}

type grpc struct {
	Port string `yaml:"port"`
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
	Grpc     grpc     `yaml:"grpc"`
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

	// 获取docker等linux系统环境变量
	appPort := os.Getenv("APP_PORT")
	if appPort != "" {
		appPortInt,_ := strconv.Atoi(appPort)
		conf.App.Port = appPortInt
	}
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort != "" {
		conf.Grpc.Port = grpcPort
	}
	database := os.Getenv("MYSQL_DATABASE")
	if database != "" {
		conf.Database.Database = database
	}

	host := os.Getenv("MYSQL_HOST")
	if host != "" {
		conf.Database.Host = host
	}

	user := os.Getenv("MYSQL_USER")
	if user != "" {
		conf.Database.Username = user
	}
	password := os.Getenv("MYSQL_PASS")
	if password != "" {
		conf.Database.Password = password
	}
	return conf
}

func SetDomain(value string) {
	conf.App.Domain = value
}
