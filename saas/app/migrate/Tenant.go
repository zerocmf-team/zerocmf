/**
** @创建时间: 2021/11/24 12:54
** @作者　　: return
** @描述　　:
 */

package migrate

type Tenant struct {
	Id           int    `json:"id"`
	TenantId     int    `gorm:"type:int(11);not null;index:idx_tenant_id;comment:租户id" json:"tenant_id"`
	Company      string `gorm:"type:varchar(100);not null;comment:公司名称" json:"company"`
	UserLogin    string `gorm:"type:varchar(60);not null;index:idx_user_login;comment:租户账号" json:"user_login"`
	AliasName    string `gorm:"type:varchar(60);not null;index:idx_user_name;comment:子账户登录别名" json:"alias_name"`
	Mobile       string `gorm:"type:varchar(20);not null;comment:手机号" json:"mobile"`
	UserRealName string `gorm:"type:varchar(50);comment:租户真实姓名" json:"user_realname"`
	UserEmail    string `gorm:"type:varchar(100);comment:租户邮箱" json:"user_email"`
	UserStatus   int    `gorm:"type:tinyint(3);default:1;not null;comment:租户状态" json:"user_status"`
	Type         int    `gorm:"type:tinyint(3);not null;comment:订阅类型(0:体验版)" json:"type"`
	CreateAt     int64  `gorm:"type:bigint(20);comment:启用时间" json:"create_at"`
	UpdateAt     int64  `gorm:"type:bigint(20);comment:更新时间" json:"update_at"`
	ExpireAt     int64  `gorm:"type:bigint(20);comment:失效时间,-1无限期" json:"expire_at"`
}
