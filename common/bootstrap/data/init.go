/**
** @创建时间: 2022/3/14 12:57
** @作者　　: return
** @描述　　:
 */

package data

type iData interface {
	InitContext()
}

type Data struct {
	*Context
	Rest
}

func (rest *Data) InitContext() (c *Context) {
	return new(Context)
}