/**
** @创建时间: 2021/12/20 08:58
** @作者　　: return
** @描述　　:
 */

package model

import (
	"encoding/json"
	"github.com/gincmf/bootstrap/config"
	"gorm.io/gorm"
	"io/ioutil"
	"strings"
)

type AdminMenu struct {
	Id         int     `json:"id"`
	Object     string  `gorm:"type:varchar(255);comment:唯一资源名称（用作）" json:"object"`
	ParentId   int     `gorm:"type:int(11);default:0;comment:父级id" json:"parent_id"`
	Name       string  `gorm:"type:varchar(30);comment:路由名称" json:"name"`
	Path       string  `gorm:"type:varchar(100);comment:路由路径" json:"path"`
	Icon       string  `gorm:"type:varchar(30);comment:图标名称" json:"icon"`
	HideInMenu int     `gorm:"type:tinyint(3);comment:菜单中隐藏;default:0" json:"hide_in_menu"`
	ListOrder  float64 `gorm:"type:float;default:10000;comment:排序（越大越靠前）" json:"list_order"`
}

type adminMenu struct {
	Object     string         `json:"object"`
	ParentId   int            `json:"parent_id"`
	Name       string         `json:"name"`
	Path       string         `json:"path"`
	Icon       string         `json:"icon"`
	HideInMenu int            `json:"hide_in_menu"`
	ListOrder  float64        `json:"list_order"`
	ApiVersion string         `json:"api_version"`
	Api        []AdminMenuApi `json:"api"`
	Children   []adminMenu    `json:"children"`
}

type AdminMenuApi struct {
	Id     int    `json:"id"`
	Object string `gorm:"type:varchar(255);comment:唯一资源名称（用作）" json:"object"`
	Desc   string `gorm:"type:varchar(100);comment:资源描述" json:"desc"`
	Path   string `gorm:"type:varchar(100);comment:路由路径" json:"path"`
	Method string `gorm:"type:varchar(10);comment:路由方法" json:"method"`
	Status int    `gorm:"type:tinyint(3);default:1" json:"status"`
}

func (_ *AdminMenu) AutoMigrate(db *gorm.DB) {
	db.Migrator().AutoMigrate(&AdminMenu{})
	db.Migrator().AutoMigrate(&AdminMenuApi{})
	initMenus(db)
}

func initMenus(db *gorm.DB) {
	var menus []adminMenu
	bytes, err := ioutil.ReadFile("data/menu.json")
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(bytes, &menus)
	if err != nil {
		// 增加json中的菜单
		panic("菜单配置文件有误，请检查菜单配置项文件：" + err.Error())
	}
	recursionAddMenu(menus, 0, db)
}

/**
 * @Author return
 * @Description //递归增加菜单
 * @Date 8:09 上午 2020/8/5
 * @Param
 * @return
 **/

var apiVersion = "v1"

func recursionAddMenu(menus []adminMenu, parentId int, db *gorm.DB) {

	// 增加当前层级
	for _, v := range menus {

		menu := AdminMenu{
			ParentId:   parentId,
			Object:     v.Object,
			Name:       v.Name,
			Path:       v.Path,
			HideInMenu: v.HideInMenu,
			ListOrder:  v.ListOrder,
		}

		if v.ApiVersion != "" {
			apiVersion = v.ApiVersion
		}

		// 保存菜单
		if v.Path != "" {
			tx := db.Where("object = ?", menu.Object).First(&menu)
			if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
				panic(tx.Error)
			}
			tx = db.Save(&menu)
			if tx.Error != nil {
				panic(tx.Error)
			}
			if len(v.Children) > 0 {
				recursionAddMenu(v.Children, menu.Id, db)
			}
		}

		// 插入权限api
		for _, v := range v.Api {
			v.Object = menu.Object
			v.Path = "/api/" + apiVersion + v.Path

			v.Method = strings.ToUpper(v.Method)

			api := AdminMenuApi{}
			tx := db.Where("object = ?", menu.Object).First(&api)
			if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
				panic(tx.Error)
			}

			if api.Id > 0 {
				v.Id = api.Id
			}

			tx = db.Save(&v)
			if tx.Error != nil {
				panic(tx.Error)
			}
		}
	}
}

func inRule(object string, title string, db *gorm.DB) error {
	conf := config.Config()
	authRule := AuthRule{
		App:   conf.App.Name,
		Name:  object,
		Title: title,
	}
	tx := db.Where("name = ?", object).FirstOrCreate(&authRule)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
