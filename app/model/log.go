package model

type Log struct {
	Id             int    `json:"id"`
	ModuleName     string `gorm:"type:varchar(50)" json:"module_name"`
	ControllerName string `gorm:"type:varchar(50)" json:"controller_name"`
	ActionName     string `gorm:"type:varchar(50)" json:"action_name"`
	Url            string `gorm:"type:varchar(255)" json:"Url"`
	RequestIp      string `gorm:"type:varchar(20)" json:"request_ip"`
	UserId         int    `gorm:"type:int(11)" json:"userid"`
	UserNickname   string `gorm:"type:varchar(50)" json:"user_nickname"`
	Message        string `gorm:"type:varchar(255)" json:"message"`
	CreateAt       int64  `gorm:"type:int(11)" json:"create_at"`
}
