/**
** @创建时间: 2021/11/25 20:13
** @作者　　: return
** @描述　　:
 */

package db

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/gorm-adapter/v3"
	"strconv"
)

func (db *database) NewEnforcer(tenantId string) (e *casbin.Enforcer, err error) {

	driverName := db.Type
	username := db.Username
	password := db.Password
	host := db.Host
	port := strconv.Itoa(db.Port)
	database := db.Database
	if tenantId != "" {
		database = "tenant_" + tenantId
	}

	a, err := gormadapter.NewAdapter(driverName, username+":"+password+"@tcp("+host+":"+port+")/"+database, true) // Your driver and data source.

	if err != nil {
		panic(err)
		return
	}

	// 从字符串初始化模型
	text :=
		`
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
		`

	m, _ := model.NewModelFromString(text)

	e, err = casbin.NewEnforcer(m, a)
	//e, err = casbin.NewEnforcer("config/rbac_model.conf", a)

	return
}
