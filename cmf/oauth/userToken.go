/**
** @创建时间: 2020/8/31 12:25 下午
** @作者　　: return
** @描述　　:
 */
package oauth

type userToken struct {
	Id       int    `json:"id"`
	UserId   int    `gorm:"type:int(11);comment:'所属用户id';not null" json:"user_id"`
	ExpireAt int64  `gorm:"type:int(11);comment:'失效时间'" json:"expire_time"`
	CreateAt int64  `gorm:"type:int(11)" json:"create_at"`
	Token    string `gorm:"type:varchar(64);comment:'token'" json:"token"`
}
