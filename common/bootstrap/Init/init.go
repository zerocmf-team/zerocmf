/**
** @创建时间: 2022/3/14 12:57
** @作者　　: return
** @描述　　:
 */

package Init

type iData interface {
	InitContext()
}

type Data struct {
	*context
}

func (rest *Data) Context() (data *Data) {
	return &Data{
		context: new(context),
	}
}
