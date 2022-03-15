/**
** @创建时间: 2020/7/18 10:53 上午
** @作者　　: return
 */
package model

type Role struct {
	Id        int     `json:"id"`
	ParentId  int     `gorm:"type:int(11);comment:'所属父类id';default:0" json:"parent_id"`
	Name      string  `gorm:"type:varchar(30);comment:'名称'" json:"name"`
	Remark    string  `gorm:"type:varchar(255);comment:'备注'" json:"remark"`
	ListOrder float64 `gorm:"type:float;comment:'排序';default:'10000'" json:"list_order"`
	CreateAt  int64   `gorm:"type:int(11)" json:"create_at"`
	UpdateAt  int64   `gorm:"type:int(11)" json:"update_at"`
	Status    int     `gorm:"type:tinyint(3);comment:'状态';default:1" json:"status"`
}

type RoleUser struct {
	Id     int `json:"id"`
	RoleId int `gorm:"type:int(11);comment:'角色id';not null" json:"role_id"`
	UserId int `gorm:"type:int(11);comment:'所属用户id';not null" json:"user_id"`
}
