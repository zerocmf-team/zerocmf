/**
** @创建时间: 2022/3/13 13:33
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
	"strconv"
	"sync"
)

type database struct {
	mu       sync.RWMutex
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

var curDb *database

func Database() *database {
	if curDb == nil {
		curDb = new(database)
	}
	return curDb
}

func (db *database) Db() (gdb *gorm.DB) {
	dbName := db.Database
	db.mu.Lock()
	gdb = db.newConn(dbName)
	db.mu.Unlock()
	return
}

func (db *database) ManualDb(tenantId string) (gdb *gorm.DB) {
	dbName := "tenant_" + tenantId
	// 未指定则默认为主库
	if tenantId == "" {
		conf := config.Config()
		dbName = conf.Database.Database
	}
	db.mu.Lock()
	gdb = db.newConn(dbName)
	db.mu.Unlock()
	return
}

func (db *database) newConn(database string) *gorm.DB {

	if database == "" {
		panic("database cannot empty")
	}

	db.CreateTable(database)

	username := db.Username
	pwd := db.Password
	host := db.Host
	port := strconv.Itoa(db.Port)
	prefix := db.Prefix
	dsn := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	gDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix, // 表名前缀，`User`表为`t_users`
			SingularTable: true,   // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		panic("failed to connect database：" + err.Error())
	}
	return gDb

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 不存在数据库则新建
 * @Date 2021/11/23 21:12:56
 * @Param
 * @return
 **/

func (db *database) CreateTable(dbName string) {
	typ := db.Type
	user := db.Username
	pwd := db.Password
	host := db.Host
	port := strconv.Itoa(db.Port)

	dataSource := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/"
	sqlDb, tempErr := sql.Open(typ, dataSource)
	if tempErr != nil {
		panic(new(error))
	}
	_, err := sqlDb.Exec("CREATE DATABASE IF NOT EXISTS " + dbName + " CHARACTER set utf8mb4 COLLATE utf8mb4_general_ci")
	if err != nil {
		panic(err)
	}
	sqlDb.Close()
}
