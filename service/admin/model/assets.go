/**
** @创建时间: 2021/11/26 08:45
** @作者　　: return
** @描述　　:
 */

package model

import (
	"errors"
	"gincmf/common/bootstrap/paginate"
	"gincmf/common/bootstrap/util"
	"gorm.io/gorm"
)

/**
 * @Author return <1140444693@qq.com>
 * @Description 资源模型
 * @Date 2021/11/26 8:47:47
 * @Param
 * @return
 **/

type Assets struct {
	Id         int    `json:"id"`
	UserId     int    `gorm:"type:int(11);not null;comment:所属用户id" json:"user_id"`
	FileSize   int64  `gorm:"type:int(11);not null;comment:文件大小" json:"file_size"`
	CreateAt   int64  `gorm:"type:int(10);default:0;comment:上传时间" json:"create_at"`
	Status     int    `gorm:"type:tinyint(3);default:1;comment:文件状态" json:"status"`
	FileKey    string `gorm:"type:varchar(64);not null;comment:文件唯一码(md5)" json:"file_key"`
	RemarkName string `gorm:"type:varchar(100);not null;comment:文件名" json:"remark_name"`
	FileName   string `gorm:"type:varchar(100);not null;comment:文件名" json:"file_name"`
	FilePath   string `gorm:"type:varchar(100);not null;comment:文件路径" json:"file_path"`
	PrevPath   string `gorm:"-" json:"prev_path"` // 前台预览地址
	Suffix     string `gorm:"type:varchar(10);not null;comment:文件后缀" json:"suffix"`
	AssetType  int    `gorm:"column:type;type:tinyint(3);not null;comment:资源类型" json:"asset_type"`
	More       string `gorm:"type:longtext;comment:更多配置" json:"more"`
}

func (model *Assets) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Assets{})
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取分页数据
 * @Date 2021/12/6 22:37:38
 * @Param
 * @return
 **/

func (model *Assets) Get(db *gorm.DB, current int, pageSize int, query string, queryArgs []interface{}) (paginateData paginate.Paginate, err error) {

	var assets []Assets

	var total int64 = 0

	tx := db.Where(query, queryArgs...).Find(&assets).Count(&total)
	if err := util.IsDbErr(tx); err != nil {
		return paginateData, errors.New("数据库连接出错：" + err.Error())
	}

	tx = db.Where(query, queryArgs...).Limit(pageSize).Offset((current - 1) * pageSize).Order("id desc").Find(&assets)

	if err := util.IsDbErr(tx); err != nil {
		return paginateData, errors.New("数据库连接出错：" + err.Error())
	}

	for k, v := range assets {
		prevPath := util.FileUrl(v.FilePath)
		assets[k].PrevPath = prevPath
	}

	paginateData = paginate.Paginate{Data: assets, Current: current, PageSize: pageSize, Total: total}
	return paginateData, nil

}


