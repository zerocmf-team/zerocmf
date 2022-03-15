/**
** @创建时间: 2020/10/3 1:38 下午
** @作者　　: return
** @描述　　:
 */
package model

import (
	"database/sql"
	"errors"
	"github.com/gincmf/cmf/data"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func CreateTable(dbName string, conf data.ConfigDefault) {
	//创建不存在的数据库
	dbUser := conf.Database.User
	dbPwd := conf.Database.Pwd
	dbHost := conf.Database.Host
	dbPort := conf.Database.Port
	dbType := conf.Database.Type
	dataSource := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/"
	tempDb, tempErr := sql.Open(dbType, dataSource)
	if tempErr != nil {
		panic(new(error))
	}
	_, err := tempDb.Exec("CREATE DATABASE IF NOT EXISTS " + dbName + " CHARACTER set utf8mb4 COLLATE utf8mb4_general_ci")
	if err != nil {
		panic(err)
	}
	tempDb.Close()
}

func NewDsn(dbName string, conf data.ConfigDefault) string {
	dataSource := conf.Database.User + ":" + conf.Database.Pwd + "@tcp(" + conf.Database.Host + ":" + conf.Database.Port + ")/"
	dsn := dataSource + dbName + "?charset=" + conf.Database.Charset + "&parseTime=True&loc=Local"
	return dsn
}

func NewDb(dsn string, prefix string) *gorm.DB {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)

	gorm.ErrRecordNotFound = errors.New("内容不存在！")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,   // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	})

	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(50)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(50)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Minute)

	if err != nil {
		panic(err)
	}

	return db
}
