/**
** @创建时间: 2020/9/10 10:39 上午
** @作者　　: return
** @描述　　: 商户表
 */
package model

type Merchant struct {
	Id          int    `gorm:"index:inx_id,unique" json:"id"`
	MerchantId  string `gorm:"type:varchar(16);comment:ISV平台商户唯一id;not null" json:"merchant_id"`
	Name        string `gorm:"type:varchar(40);comment:商户名称;not null" json:"name"`
	Description string `gorm:"type:varchar(255);comment:商户描述;not null" json:"description"`
	UserLogin   string `gorm:"type:varchar(20);comment:登录名;not null" json:"user_login"`
	UserPass    string `gorm:"type:varchar(64);comment:密码;not null" json:"user_pass"`
	Logo        string `gorm:"type:varchar(255);comment:商户LOGO;not null" json:"logo"`
	CreateAt    int64  `gorm:"type:int(10);comment:'创建时间';default:0" json:"create_at"`
	UpdateAt    int64  `gorm:"type:int(10);comment:'更新时间';default:0" json:"update_at"`
}
