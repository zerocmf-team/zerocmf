package model

type User struct {
	Id                int     `json:"id"`
	UserType          int     `gorm:"type:tinyint(3);not null" json:"user_type"`
	Gender            int     `gorm:"type:tinyint(2)" json:"gender"`
	Birthday          int64   `gorm:"type:int(11)" json:"birthday"`
	LastLoginAt       int64   `gorm:"type:int(11)" json:"last_login_at"`
	Score             int     `json:"score"`
	Coin              int     `json:"coin"`
	Balance           float64 `gorm:"type:decimal(10,2);not null" json:"balance"`
	CreateAt          int64   `gorm:"type:int(11)" json:"create_at"`
	UpdateAt          int64   `gorm:"type:int(11)" json:"update_at"`
	UserStatus        int     `gorm:"type:tinyint(3);not null" json:"user_status"`
	UserLogin         string  `gorm:"type:varchar(60);not null" json:"user_login"`
	UserPass          string  `gorm:"type:varchar(64);not null" json:"-"`
	UserNickname      string  `gorm:"type:varchar(50);not null" json:"user_nickname"`
	UserRealName      string  `gorm:"type:varchar(50);not null" json:"user_realname"`
	UserEmail         string  `gorm:"type:varchar(100);not null" json:"user_email"`
	UserUrl           string  `gorm:"type:varchar(100);not null" json:"user_url"`
	Avatar            string  `json:"avatar"`
	Signature         string  `json:"signature"`
	LastLoginIp       string  `json:"last_loginip"`
	UserActivationKey string  `json:"user_activation_key"`
	Mobile            string  `gorm:"type:varchar(20);not null" json:"mobile"`
	DepartmentId      int     `gorm:"type:int(11);comment:'部门id'" json:"department_id"`
	more              string  `gorm:"type:text" json:"more"`
}
