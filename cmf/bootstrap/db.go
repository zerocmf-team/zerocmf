/**
** @创建时间: 2020/11/3 8:35 下午
** @作者　　: return
** @描述　　:
 */
package bootstrap

import (
	"errors"
	"fmt"
	"github.com/gincmf/cmf/data"
	"github.com/gincmf/cmf/model"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	commDb  *gorm.DB
	redisDb *redis.Client
)

// 默认Db
func Db() *gorm.DB {
	if commDb == nil {
		config := Conf()
		dbName := config.Database.Name
		//创建不存在的数据库
		model.CreateTable(dbName, config)
		dsn := model.NewDsn(dbName, config)
		commDb = model.NewDb(dsn, config.Database.Prefix)
	}
	return commDb
}

// 获取当前指定Db
//func NewDb() *gorm.DB {
//	if db == nil {
//		config := Conf()
//		dbName := config.Database.Name
//		//创建不存在的数据库
//		model.CreateTable(dbName, config)
//		dsn := model.NewDsn(dbName, config)
//		fmt.Println("dsn",dsn)
//		db = model.NewDb(dsn, config.Database.Prefix)
//	}
//	return db
//}

// 手动指定Db
func ManualDb(dbName string) *gorm.DB {
	config := Conf()
	dsn := model.NewDsn(dbName, config)
	return model.NewDb(dsn, config.Database.Prefix)
}

func RedisDb() *redis.Client {
	if redisDb == nil {
		database := Conf().Redis
		empty := data.Redis{}
		if database != empty {
			if database.Host == "" {
				panic("redis host not empty")
			}

			if database.Port == "" {
				panic("redis port not empty")
			}
			redisDb = redis.NewClient(&redis.Options{
				Addr:     database.Host + ":" + database.Port,
				Password: database.Pwd,      // no password set
				DB:       database.Database, // use default DB
			})
			fmt.Println("RedisDb：", redisDb)
			result, err := redisDb.Ping().Result()
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("redis连接状态：", result)
		}
	}
	return redisDb
}

func ManualRedisDb(host string, pwd string) (*redis.Client, error) {
	database := Conf().Redis
	empty := data.Redis{}
	if database != empty && database.Enabled {
		if host == "" {
			return redisDb, errors.New("redis host not empty")
		}

		redisDb = redis.NewClient(&redis.Options{
			Addr:     host + ":" + database.Port,
			Password: pwd,               // no password set
			DB:       database.Database, // use default DB
		})
		fmt.Println("RedisDb：", redisDb)
		result, err := redisDb.Ping().Result()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("redis连接状态：", result)
	}
	return redisDb, nil
}
