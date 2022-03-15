/**
** @创建时间: 2021/11/26 11:02
** @作者　　: return
** @描述　　:
 */

package model

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
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
	Object     string      `json:"object"`
	ParentId   int         `json:"parent_id"`
	Name       string      `json:"name"`
	Path       string      `json:"path"`
	Icon       string      `json:"icon"`
	HideInMenu int         `json:"hide_in_menu"`
	ListOrder  float64     `json:"list_order"`
	Children   []adminMenu `json:"children"`
}

func (_ *AdminMenu) AutoMigrate(db *gorm.DB) {
	db.Migrator().AutoMigrate(&AdminMenu{})
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
		panic( "菜单配置文件有误，请检查菜单配置项文件：" + err.Error())
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

		// 保存菜单
		if v.Path != "" {
			tx := db.Where("object = ?", menu.Object).FirstOrCreate(&menu)
			if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
				fmt.Println("tx.err",tx.Error)
				return
			}
			if len(v.Children) > 0 {
				recursionAddMenu(v.Children, menu.Id, db)
			}
		}
	}
}
