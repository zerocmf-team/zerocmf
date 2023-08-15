/**
** @创建时间: 2022/8/22 12:55
** @作者　　: return
** @描述　　:
 */

package redis

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/logx"
)

type Redis struct {
	*redis.Client `json:",optional"`
	Enabled       bool
	Host          string
	Database      int
	Password      string
	Port          int
}

func NewRedis(database Redis) Redis {
	curRedis := redis.NewClient(&redis.Options{
		Addr:     database.Host + ":" + strconv.Itoa(database.Port),
		Password: database.Password, // no password set
		DB:       database.Database, // use default DB
	})
	result, err := curRedis.Ping().Result()
	if err != nil {
		logx.Error("redis异常", err.Error())
	}
	fmt.Println("redis连接状态：", result)
	database.Client = curRedis
	return database
}
