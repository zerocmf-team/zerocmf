/**
** @创建时间: 2022/3/13 13:33
** @作者　　: return
** @描述　　:
 */

package database

import (
	"database/sql"
	"github.com/jinzhu/copier"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"strconv"
	"time"
)

type Database struct {
	Name     string
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

type GormDB struct {
	Database Database
	Db       *gorm.DB
}

func NewGormDb(db Database) (gormDb *gorm.DB) {
	dbName := db.Database
	gormDb = db.newConn(dbName)
	gormDb.Set("siteId", "")
	return
}

func (db *Database) NewConf(siteId string) (conf Database) {
	copier.Copy(&conf, &db)
	if siteId != "" {
		conf.Database = "site_" + siteId + "_" + db.Name
	}
	//db.Db = db.newConn(conf.Database)
	return
}

func (db *Database) ManualDb(siteId string) (outDb *gorm.DB) {
	dbName := "site_" + siteId + "_" + db.Name
	// 未指定则默认为主库
	if siteId == "" {
		dbName = db.Database
	}
	outDb = db.newConn(dbName)
	outDb.Set("siteId", siteId)
	return
}

func (db *Database) newConn(database string) *gorm.DB {

	if database == "" {
		panic("Database cannot empty")
	}

	db.CreateTable(database)

	username := db.Username
	pwd := db.Password
	host := db.Host

	HOST := os.Getenv("MYSQL_HOST")
	if HOST != "" {
		host = HOST
	}

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
		panic("failed to connect Database：" + err.Error())
	}

	gDb.Set("prefix", prefix)

	sqlDB, err := gDb.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(500)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(2000)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(90 * time.Minute)
	return gDb

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 不存在数据库则新建
 * @Date 2021/11/23 21:12:56
 * @Param
 * @return
 **/

func (db *Database) CreateTable(dbName string) {

	database := os.Getenv("MYSQL_DATABASE")
	if database != "" {
		db.Database = database
	}

	host := os.Getenv("MYSQL_HOST")
	if host != "" {
		db.Host = host
	}

	username := os.Getenv("MYSQL_USER")
	if username != "" {
		db.Username = username
	}
	password := os.Getenv("MYSQL_PASS")
	if password != "" {
		db.Password = password
	}

	port := strconv.Itoa(db.Port)

	typ := db.Type
	user := db.Username
	pwd := db.Password
	host = db.Host

	dataSource := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/"

	sqlDb, tempErr := sql.Open(typ, dataSource)
	if tempErr != nil {
		panic(new(error))
	}

	_, err := sqlDb.Exec("CREATE DATABASE IF NOT EXISTS `" + dbName + "` CHARACTER set utf8mb4 COLLATE utf8mb4_general_ci")
	if err != nil {
		panic(err)
	}
	sqlDb.Close()
}
