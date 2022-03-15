/**
** @创建时间: 2020/7/18 10:56 上午
** @作者　　: return
 */
package migrate

import (
	"gincmf/app/model"
	cmf "github.com/gincmf/cmf/bootstrap"
	"time"
)

type role struct {
	Migrate
}

func (_ *role) AutoMigrate() {
	cmf.Db.AutoMigrate(&model.Role{})

	role := []model.Role{
		model.Role{
			Name:      "超级管理员",
			Remark:    "拥有网站最高管理员权限！",
			ListOrder: 0,
			CreateAt:  time.Now().Unix(),
			UpdateAt:  time.Now().Unix(),
		},
		model.Role{
			Name:      "普通管理员",
			Remark:    "权限由最高管理员分配！",
			ListOrder: 0,
			CreateAt:  time.Now().Unix(),
			UpdateAt:  time.Now().Unix(),
		},
	}

	// 添加角色权限
	for _,v := range role{
		cmf.Db.Where(model.Role{Name: v.Name}).FirstOrCreate(&v)
	}

	// 添加role_user关系表
	cmf.Db.AutoMigrate(&model.RoleUser{})

}
