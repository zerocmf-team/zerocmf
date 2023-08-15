/**
** @创建时间: 2022/3/13 13:33
** @作者　　: return
** @描述　　:
 */

package database

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	Name     string `json:"Name,optional"`
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

func NewGormDb(db Database, siteIds ...string) (gormDb *gorm.DB) {
	dbName := db.Database
	siteId := ""
	if len(siteId) > 0 {
		siteId = siteIds[0]
		dbName = "site_" + siteId + "_" + db.Name
	}
	gormDb = db.newConn(dbName)
	gormDb.Set("siteId", siteId)
	return
}

func CreateGormDb(db Database, siteIds ...string) (gormDb *gorm.DB) {
	dbName := db.Database
	siteId := ""
	if len(siteIds) > 0 {
		siteId = siteIds[0]
		dbName = "site_" + siteId + "_" + db.Name
	}

	if dbName == "" {
		panic("Database cannot empty")
	}
	db.createTable(dbName)
	gormDb = db.newConn(dbName)
	gormDb.Set("siteId", siteId)
	return
}

func (db *Database) NewConf(siteId int64) (conf Database) {
	copier.Copy(&conf, &db)
	if siteId > 0 {
		siteIdStr := strconv.FormatInt(siteId, 10)
		conf.Database = "site_" + siteIdStr + "_" + db.Name
	}
	//db.Db = db.newConn(conf.Database)
	return
}

func (db *Database) ManualDb(siteIdInt int64) (outDb *gorm.DB) {
	siteId := strconv.FormatInt(siteIdInt, 10)
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

	//db.createTable(database)

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

func (db *Database) createTable(dbName string) {

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
	createSql := "CREATE DATABASE IF NOT EXISTS `" + dbName + "` CHARACTER set utf8mb4 COLLATE utf8mb4_general_ci"
	_, err := sqlDb.Exec(createSql)
	if err != nil {
		panic(err)
	}
	sqlDb.Close()
}

/**
Desc: 获取获取连接串
Author: daifuyang
Contact: github.com/daifuyang
Date: Date: 2023-07-17 14:17:54
*/

func (db *Database) Dsn() (dsn string) {
	username := db.Username
	pwd := db.Password
	host := db.Host

	HOST := os.Getenv("MYSQL_HOST")
	if HOST != "" {
		host = HOST
	}

	port := strconv.Itoa(db.Port)
	database := db.Database

	dsn = username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	return
}

func (db *Database) Migrate(tables []string) {
	dbName := db.Database
	db.createTable(dbName)
	open, err := sql.Open("mysql", db.Dsn())
	if err != nil {
		return
	}
	defer open.Close()

	var wd string
	wd, err = os.Getwd()
	if err != nil {
		fmt.Println("无法获取当前工作目录：", err)
		return
	}
	// 获取上级目录的路径
	parentDir := filepath.Dir(wd)

	if parentDir == "/" {
		parentDir = "/app"
	}

	var file *os.File
	for _, v := range tables {
		filename := parentDir + "/model/mysql/" + v + ".sql"
		file, err = os.Open(filename)
		if err != nil {
			fmt.Println("无法打开文件:", err)
			return
		}
		var content []byte
		content, err = io.ReadAll(file)
		if err != nil {
			fmt.Println("读取文件失败:", err)
			return
		}
		sqlStr := string(content)
		_, err = open.Exec(sqlStr)
		if err != nil {
			panic(err)
		}
	}
	defer file.Close()
}
