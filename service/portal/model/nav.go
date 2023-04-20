/**
** @创建时间: 2021/1/3 8:30 下午
** @作者　　: return
** @描述　　:
 */

package model

import (
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/util"
)

type Nav struct {
	Id     int    `json:"id"`
	Key    string `gorm:"type:varchar(10);comment:导航类型;not null" json:"key"`
	Name   string `gorm:"type:varchar(20);comment:主导航名称" json:"name"`
	Remark string `gorm:"type:varchar(255);comment:备注" json:"remark"`
}

type NavItem struct {
	Id        int       `json:"id"`
	NavId     int       `gorm:"type:int(11);comment:导航id;not null" json:"nav_id"`
	ParentId  int       `gorm:"type:int(11);comment:所属父类id;default:0" json:"parent_id"`
	Status    int       `gorm:"type:tinyint(3);default:1" json:"status"`
	ListOrder float64   `gorm:"type:float;comment:排序;default:10000" json:"list_order"`
	Name      string    `gorm:"type:varchar(50);comment:路由名称" json:"name"`
	Target    string    `gorm:"type:varchar(10);comment:目标状态" json:"target"`
	Href      string    `gorm:"type:varchar(100);comment:路由路径" json:"href"`
	HrefType  int       `gorm:"type:tinyint(3);comment:路由路径类型（0.系统映射，1.自定义）" json:"href_type"`
	Icon      string    `gorm:"type:varchar(255);comment:图标地址" json:"icon"`
	IconPrev  string    `gorm:"-" json:"icon_prev"`
	Path      string    `gorm:"type:varchar(255);comment:路由路径" json:"path"`
	Children  []NavItem `gorm:"-" json:"children,omitempty"`
}

type NavItemOptions struct {
	Value    int              `json:"value"`
	Title    string           `gorm:"type:varchar(50);comment:路由名称" json:"title"`
	Children []NavItemOptions `json:"children"`
}

func (model *Nav) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Nav{})
	db.AutoMigrate(&NavItem{})
	hasType := db.Migrator().HasIndex(&Nav{}, "idx_key")
	if hasType == false {
		db.Migrator().CreateIndex(&Nav{}, "idx_key")
	}
}

func (model *Nav) Get(db *gorm.DB, current, pageSize int, query string, queryArgs []interface{}) (paginate data.Paginate, err error) {

	var total int64 = 0
	var nav []Nav
	db.Where(query, queryArgs...).Find(&nav).Count(&total)
	tx := db.Where(query, queryArgs...).Limit(pageSize).Offset((current - 1) * pageSize).Find(&nav)
	if tx.Error != nil {
		if util.IsDbErr(tx) != nil {
			err = tx.Error
			return
		}
	}
	paginate = data.Paginate{Data: nav, Current: current, PageSize: pageSize, Total: total}
	if len(nav) == 0 {
		paginate.Data = make([]string, 0)
	}
	return paginate, nil
}

func (model *Nav) Show(db *gorm.DB, query string, queryArgs []interface{}) error {

	tx := db.Where(query, queryArgs...).First(&model)
	if util.IsDbErr(tx) != nil {
		return tx.Error
	}
	return nil
}

func (model *NavItem) Show(db *gorm.DB, query string, queryArgs []interface{}) (NavItem, error) {
	navItem := NavItem{}
	tx := db.Where(query, queryArgs...).Find(&navItem)
	if tx.Error != nil {
		return navItem, tx.Error
	}
	if navItem.Icon != "" {
		navItem.IconPrev = util.FileUrl(navItem.Icon)
	}
	href := navItem.Href
	if href != "" && href[0:1] != "/" {
		href = "/" + href
	}
	navItem.Href = href
	return navItem, nil
}

func (model *NavItem) GetWithChildPaginate(db *gorm.DB, current, pageSize int, query string, queryArgs []interface{}) (data.Paginate, error) {

	var total int64 = 0
	var navItem []NavItem
	tx := db.Where(query, queryArgs...).Find(&navItem).Order("list_order desc").Count(&total)
	tx = db.Where(query, queryArgs...).Limit(pageSize).Offset((current - 1) * pageSize).Order("list_order desc").Find(&navItem)
	if tx.Error != nil {
		return data.Paginate{}, tx.Error
	}
	res := model.recursionNav(navItem, 0)
	paginate := data.Paginate{Data: res, Current: current, PageSize: pageSize, Total: total}
	if len(navItem) == 0 {
		paginate.Data = make([]string, 0)
	}
	return paginate, nil
}

func (model *NavItem) GetWithChild(db *gorm.DB, query string, queryArgs []interface{}) (navItem []NavItem, err error) {

	tx := db.Where(query, queryArgs...).Order("list_order desc").Find(&navItem)
	if tx.Error != nil {
		return navItem, tx.Error
	}
	data := model.recursionNav(navItem, 0)

	if data == nil {
		data = make([]NavItem, 0)
	}

	return data, nil
}

func (model *NavItem) recursionNav(nav []NavItem, parentId int) []NavItem {

	var navItems []NavItem
	// 增加当前层级
	for _, v := range nav {

		href := v.Href

		if parentId == v.ParentId {
			ni := NavItem{
				Id:        v.Id,
				NavId:     v.NavId,
				ParentId:  parentId,
				Status:    v.Status,
				Name:      v.Name,
				Path:      v.Path,
				Target:    v.Target,
				Href:      href,
				Icon:      v.Icon,
				IconPrev:  util.FileUrl(v.Icon),
				ListOrder: v.ListOrder,
			}

			childNav := model.recursionNav(nav, v.Id)
			ni.Children = childNav

			navItems = append(navItems, ni)
		}

	}

	return navItems
}

var navItemOptions []NavItemOptions

/**
 * @Author return <1140444693@qq.com>
 * @Description 查看全部分类菜单
 * @Date 2022/4/14 20:59:8
 * @Param
 * @return
 **/

func (model *NavItem) OptionsList(db *gorm.DB, query string, queryArgs []interface{}) []NavItemOptions {
	var navItem []NavItem
	navItemOptions = make([]NavItemOptions, 0)
	db.Where(query, queryArgs...).Order("list_order desc").Find(&navItem)
	data := model.recursionOptions(navItem, 0)
	return data
}

func (model *NavItem) indent(level int) string {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "    |--"
	}
	return indent
}

func (model *NavItem) recursionOptions(nav []NavItem, parentId int) []NavItemOptions {
	options := make([]NavItemOptions, 0)
	for _, v := range nav {
		if parentId == v.ParentId {
			ops := NavItemOptions{
				Value: v.Id,
				Title: v.Name,
			}
			ops.Children = model.recursionOptions(nav, v.Id)
			options = append(options, ops)
		}
	}
	return options
}
