/**
** @创建时间: 2022/3/23 20:43
** @作者　　: return
** @描述　　:
 */

package casbin

import (
	"gincmf/common/bootstrap/db"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"strconv"
)

func NewEnforcer(tenantId string) (e *casbin.Enforcer, err error) {
	curDb := db.Conf()
	driverName := curDb.Type
	username := curDb.Username
	password := curDb.Password
	host := curDb.Host
	port := strconv.Itoa(curDb.Port)

	database := curDb.Database

	if tenantId != "" {
		database = "tenant_" + tenantId
	}

	a, err := gormadapter.NewAdapter(driverName, username+":"+password+"@tcp("+host+":"+port+")/"+database, true) // Your driver and data source.

	if err != nil {
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
