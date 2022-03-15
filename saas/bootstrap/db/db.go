/**
** @创建时间: 2021/11/23 12:21
** @作者　　: return
** @描述　　:
 */

package db

import (
	"database/sql"
	"gincmf/bootstrap/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db = make(map[string]*gorm.DB)

func Db() *gorm.DB {
	conf := config.Config()
	dbName := conf.Database.Database
	if db[dbName] == nil {
		db[dbName] = newConn(dbName)
	}
	return db[dbName]
}

func ManualDb(tenant string) *gorm.DB {

	dbName := tenant

	// 未指定则默认为主库
	if tenant == "" {
		conf := config.Config()
		dbName = conf.Database.Database
	}

	if db[dbName] == nil {
		db[dbName] = newConn(dbName)
	}
	return db[dbName]
}

func newConn(dbName string) *gorm.DB {

	if dbName == "" {
		panic("database cannot empty")
	}

	CreateTable(dbName)

	conf := config.Config()
	username := conf.Database.Username
	pwd := conf.Database.Password
	host := conf.Database.Host
	port := conf.Database.Port
	dsn := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database：" + err.Error())
	}
	return db

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 不存在数据库则新建
 * @Date 2021/11/23 21:12:56
 * @Param
 * @return
 **/

func CreateTable(dbName string) {
	conf := config.Config()
	user := conf.Database.Username
	pwd := conf.Database.Password
	host := conf.Database.Host
	port := conf.Database.Port
	typ := conf.Database.Type
	dataSource := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/"
	tempDb, tempErr := sql.Open(typ, dataSource)
	if tempErr != nil {
		panic(new(error))
	}
	_, err := tempDb.Exec("CREATE DATABASE IF NOT EXISTS " + dbName + " CHARACTER set utf8mb4 COLLATE utf8mb4_general_ci")
	if err != nil {
		panic(err)
	}
	tempDb.Close()
}
