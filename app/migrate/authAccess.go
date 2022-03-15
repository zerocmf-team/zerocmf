/**
** @创建时间: 2020/8/4 1:09 下午
** @作者　　: return
 */
package migrate

import (
	"gincmf/app/model"
	cmf "github.com/gincmf/cmf/bootstrap"
)

type authAccess struct {
	Migrate
}



func (_ *authAccess) AutoMigrate() {
	cmf.Db.AutoMigrate(model.AuthAccess{})
}

