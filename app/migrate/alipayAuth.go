/**
** @创建时间: 2020/9/7 11:01 上午
** @作者　　: return
** @描述　　:
 */
package migrate

import (
	"gincmf/app/model"
	cmf "github.com/gincmf/cmf/bootstrap"
)

type AlipayAuth struct {
	Migrate
}

func (_ *AlipayAuth) AutoMigrate() {
	cmf.Db.AutoMigrate(&model.AlipayAuth{})

	// 检查索引
	b := cmf.Db.Migrator().HasIndex(&model.AlipayAuth{}, "idx_id")
	if !b {
	// 新建索引
		cmf.Db.Migrator().CreateIndex(&model.AlipayAuth{}, "idx_id")
	}
}



