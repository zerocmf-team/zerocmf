/**
** @创建时间: 2021/11/23 12:21
** @作者　　: return
** @描述　　:
 */

package db

import (
	"database/sql"
	"github.com/gincmf/bootstrap/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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

func ManualDb(tenantId string) *gorm.DB {

	database := "tenant_" + tenantId

	// 未指定则默认为主库
	if tenantId == "" {
		conf := config.Config()
		database = conf.Database.Database
	}

	if db[database] == nil {
		db[database] = newConn(database)
	}
	return db[database]
}

func newConn(database string) *gorm.DB {

	if database == "" {
		panic("database cannot empty")
	}

	CreateTable(database)

	conf := config.Config()
	username := conf.Database.Username
	pwd := conf.Database.Password
	host := conf.Database.Host
	port := conf.Database.Port
	prefix := conf.Database.Prefix
	dsn := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: prefix,   // 表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
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
