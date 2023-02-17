/**
** @创建时间: 2021/11/25 20:13
** @作者　　: return
** @描述　　: casbin
 */

package database

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/util"
	"github.com/casbin/gorm-adapter/v3"
	"strconv"
)

func (db *Database) NewEnforcer(tenantId string) (e *casbin.Enforcer, err error) {

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
		m = g(r.sub, p.sub) && (menuMatch(r.obj, p.obj) || r.obj == p.obj) && r.act == p.act
		`
	// 		m = g(r.sub, p.sub) && r.obj == p.obj && (r.act == p.act || (keyMatch(r.obj, p.obj) || keyMatch2(r.obj, p.obj)) && regexMatch(r.act, p.act))
	m, _ := model.NewModelFromString(text)
	e, err = casbin.NewEnforcer(m, a)
	//e, err = casbin.NewEnforcer("config/rbac_model.conf", a)

	e.AddFunction("menuMatch", menuMatchFunc)

	return
}

func menuMatch(key1 string, key2 string) (bool bool) {
	bool = util.RegexMatch(key2, key1)
	if bool == false {
		bool = util.KeyMatch2(key1, key2)
	}
	return
}

func menuMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return (bool)(menuMatch(name1, name2)), nil
}
