/**
** @创建时间: 2022/8/22 12:55
** @作者　　: return
** @描述　　:
 */

package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type Redis struct {
	Enabled  bool
	Host     string
	Database int
	Password string
	Port     int
}

var (
	curRedis *redis.Client
)

func NewRedis(database Redis) *redis.Client {
	if curRedis == nil {
		curRedis = redis.NewClient(&redis.Options{
			Addr:     database.Host + ":" + strconv.Itoa(database.Port),
			Password: database.Password, // no password set
			DB:       database.Database, // use default DB
		})
	}
	result, err := curRedis.Ping().Result()
	if err != nil {
		logx.Error("redis异常", err.Error())
	}
	fmt.Println("redis连接状态：", result)
	return curRedis
}