/**
** @创建时间: 2020/9/7 10:48 上午
** @作者　　: return
** @描述　　:
 */
package model

type AlipayAuth struct {
	Id              int    `gorm:"index:inx_id,unique" json:"id"`
	UserId          string `gorm:"type:varchar(16);comment:授权商户的user_id;not null" json:"user_id"`
	AuthAppId       string `gorm:"type:varchar(20);comment:授权商户的appid;not null" json:"auth_app_id"`
	AppAuthToken    string `gorm:"type:varchar(40);comment:应用授权令牌;not null" json:"app_auth_token"`
	AppRefreshToken string `gorm:"type:varchar(40);comment:刷新令牌;not null" json:"app_refresh_token"`
	ExpiresIn       string `gorm:"type:varchar(16);comment:应用授权令牌的有效时间（从接口调用时间作为起始时间），单位到秒;not null" json:"expires_in"`
	ReExpiresIn     string `gorm:"type:varchar(16);comment:刷新令牌的有效时间（从接口调用时间作为起始时间），单位到秒;not null" json:"re_expires_in"`
	CreateAt        int64  `gorm:"type:int(10);comment:'创建时间';default:0" json:"create_at"`
	UpdateAt        int64  `gorm:"type:int(10);comment:'更新时间';default:0" json:"update_at"`
}
