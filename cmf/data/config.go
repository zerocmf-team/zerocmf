/**
** @创建时间: 2020/10/4 9:13 下午
** @作者　　: return
** @描述　　:
 */
package data

// App 配置文件app对象 定义了系统的基础信息
type App struct {
	Port     string
	AppName  string
	Evn      string
	AppDebug bool
	Domain   string
	NsqIp    string
	NsqaPort string
	NsqdPort string
}

// Database 配置文件数据库对象 定义了数据库的基础信息
type Database struct {
	Default  string `json:"default"`
	Type     string `json:"type"`
	Host     string `json:"host"`
	Name     string `json:"database"`
	User     string `json:"username"`
	Pwd      string `json:"password"`
	Port     string `json:"port"`
	Charset  string `json:"charset"`
	Prefix   string `json:"prefix"`
	AuthCode string `json:"authcode"`
}

// Redis 配置文件对象
type Redis struct {
	Enabled  bool   `json:"enabled"`
	Type     string `json:"type"`
	Host     string `json:"host"`
	Pwd      string `json:"password"`
	Port     string `json:"port"`
	Database int    `json:"database"`
}

//ConfigDefault 定义了配置文件初始结构
type ConfigDefault struct {
	App      App
	Template TemplateMapStruct `json:"template"`
	Database Database
	Redis    Redis
	Token    string
}

//ConfigData 定义一个空结构体
type ConfigData struct {
}

type TempDataBase struct {
	Default string   `json:"default"`
	Mysql   Database `json:"mysql"`
	Redis   Redis    `json:"redis"`
}

// 定义读取配置缓存文件
type TempConfig struct {
	Database TempDataBase
}
