/**
** @创建时间: 2020/8/5 7:28 上午
** @作者　　: return
 */
package model

type AuthRule struct {
	Id     int    `json:"id"`
	Name   string `gorm:"type:varchar(100);comment:'规则唯一英文标识,全小写'" json:"name"`
	Param  string `gorm:"type:varchar(100);comment:'额外url参数'" json:"param"`
	Title  string `gorm:"type:varchar(100);comment:'规则描述'" json:"title"`
	Status int    `gorm:"type:tinyint(3);default:1" json:"status"`
}
