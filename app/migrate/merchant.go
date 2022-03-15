/**
** @创建时间: 2020/9/10 10:51 上午
** @作者　　: return
** @描述　　:
 */
package migrate

import (
	"gincmf/app/model"
	cmf "github.com/gincmf/cmf/bootstrap"
)

type merchant struct {
	Migrate
}

func (_ *merchant) AutoMigrate() {
	cmf.Db.AutoMigrate(&model.Merchant{})

	// 检查索引
	b := cmf.Db.Migrator().HasIndex(&model.Merchant{}, "idx_id")
	if !b {
		// 新建索引
		cmf.Db.Migrator().CreateIndex(&model.Merchant{}, "idx_id")
	}
}
