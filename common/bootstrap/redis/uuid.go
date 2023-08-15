package redis

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/pjebs/optimus-go"
)

/*
 ** 唯一uid编号生成逻辑
 ** 日期 + 当天排号数量
 */

func (r Redis) EncryptUid(key string, salt int) (uid int64, err error) {

	t1 := time.Now()
	date := t1.Format("20060102")
	timeStr := t1.Format("150405")
	key = key + date
	// 设置当天失效时间
	year, month, day := t1.Date()
	today := time.Date(year, month, day, 23, 59, 59, 59, time.Local)
	curRedis := r.Client
	if curRedis == nil {
		panic("redis 链接失败")
	}
	curRedis.ExpireAt(key, today)
	val, valErr := curRedis.Incr(key).Result()
	if valErr != nil {
		_, _, line, _ := runtime.Caller(0)
		fmt.Println("redis err"+strconv.Itoa(line), valErr.Error())
		err = valErr
		return
	}

	saltStr := ""
	if salt > 0 {
		saltStr = strconv.Itoa(salt)
	}

	incr := strconv.FormatInt(val, 10)

	// 0 20230506 1
	nStr := saltStr + date + timeStr + incr
	n, _ := strconv.ParseInt(nStr, 10, 64)

	uid = EncodeId(uint64(n))
	return
}

// 加密id 纯数字

func EncodeId(id uint64) int64 {
	o := optimus.New(561604931, 848718699, 1452111999)
	number := o.Encode(id)
	return int64(number)
}

// 解密id 纯数字

func DecodeId(id uint64) int {
	o := optimus.New(561604931, 848718699, 1452111999)
	origId := o.Decode(id)
	return int(origId)
}
