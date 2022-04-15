/**
** @创建时间: 2020/12/25 2:08 下午
** @作者　　: return
** @描述　　:
 */

package model

import (
	"errors"
	"gincmf/common/bootstrap/data"
	"gincmf/common/bootstrap/util"
	"gorm.io/gorm"
	"strings"
)

// 标签内容

type PortalTag struct {
	Id          int    `json:"id"`
	Status      int    `gorm:"type:tinyint(3);comment:状态,1:发布,0:不发布;default:1;not null;" json:"status"`
	Recommended int    `gorm:"type:tinyint(3);comment:是否推荐,1:推荐;0:不推荐;default:0;not null;" json:"recommended"`
	PostCount   int64  `gorm:"type:bigint(20);comment:标签文章数;default:0;not null;" json:"post_count"`
	Name        string `gorm:"type:varchar(20);comment:标签名称;not null;" json:"name"`
}

// 标签关系

type PortalTagPost struct {
	Id     int `json:"id"`
	TagId  int `gorm:"type:bigint(20);comment:标签id;not null;" json:"tag_id"`
	PostId int `gorm:"type:bigint(20);comment:文章id;not null;" json:"post_id"`
	Status int `gorm:"type:tinyint(3);comment:状态,1:发布,0:不发布;default:1;not null;" json:"status"`
}

type PostTagResult struct {
	PortalTagPost
	Name string `gorm:"->" json:"name"`
}

func (model PortalTag) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model)
	db.AutoMigrate(&PortalTagPost{})
}

func (model *PortalTag) Index(db *gorm.DB, current, pageSize int, query string, queryArgs []interface{}) (data.Paginate, error) {
	// 获取默认的系统分页
	var total int64 = 0
	var tag []PortalTag
	db.Where(query, queryArgs...).Find(&tag).Count(&total)
	tx := db.Where(query, queryArgs...).Limit(pageSize).Offset((current - 1) * pageSize).Find(&tag)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return data.Paginate{}, tx.Error
		}
	}
	paginate := data.Paginate{Data: tag, Current: current, PageSize: pageSize, Total: total}
	if len(tag) == 0 {
		paginate.Data = make([]string, 0)
	}
	return paginate, nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description
 * @Date 2021/12/13 12:18:7
 * @Param
 * @return
 **/

func (model *PortalTag) Show(db *gorm.DB, query string, queryArgs []interface{}) (PortalTag, error) {
	tag := PortalTag{}
	tx := db.Where(query, queryArgs...).Find(&tag)
	if tx.Error != nil {
		return tag, nil
	}
	return tag, nil
}

func (model PortalTag) FirstOrSave(db *gorm.DB) (PortalTag, error) {
	// 新建
	if model.Id == 0 {
		tx := db.Create(&model)
		if tx.Error != nil {
			return PortalTag{}, tx.Error
		}
	} else {
		// 更新
		// 统计文章标签数
		var count int64
		tx := db.Where("tag_id = ?", model.Id).Group("post_id").Find(&PortalTagPost{}).Count(&count)
		if util.IsDbErr(tx) != nil {
			return PortalTag{}, tx.Error
		}
		model.PostCount = count
		tx = db.Save(model)
		if tx.Error != nil {
			return PortalTag{}, tx.Error
		}
	}
	return model, nil

}

func (model PortalTag) Save(db *gorm.DB, postId int) error {
	var count int64
	db.Where("post_id = ?", postId).Find(&PortalTagPost{}).Group("post_id").Count(&count)
	db.Where("id = ?", model.Id).First(&model)
	model.PostCount = count
	tx := db.Save(&model)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (model *PortalTagPost) FirstOrSave(db *gorm.DB, kId []int) error {

	// [0,1,2]  [1,3,4]
	postId := model.PostId
	// 查出原来的
	var tagPost []PortalTagPost
	db.Where("post_id = ?", postId).Find(&tagPost)
	// 待添加的
	var toAdd []PortalTagPost
	for _, v := range kId {
		if !new(PortalTagPost).inAddArray(v, tagPost) || len(tagPost) == 0 {
			toAdd = append(toAdd, PortalTagPost{
				TagId:  v,
				PostId: postId,
			})
		}
	}
	//待删除的
	var toDel []string
	var toDelArgs []interface{}
	for _, v := range tagPost {
		if !new(PortalTagPost).inDelArray(v.Id, kId) {
			toDel = append(toDel, "id = ?")
			toDelArgs = append(toDelArgs, v.Id)
		}
		if len(kId) == 0 {
			toDel = append(toDel, "id = ?")
			toDelArgs = append(toDelArgs, v.Id)
		}
	}
	// 删除要删除的
	if len(toDel) > 0 {
		delStr := strings.Join(toDel, " OR ")
		db.Where(delStr, toDelArgs...).Delete(&PortalTagPost{})
	}
	if len(toAdd) > 0 {
		// 增加待增加的
		db.Create(toAdd)
	}
	// 统计当前标签文章数
	for _, v := range kId {
		err := PortalTag{Id: v}.Save(db, postId)
		if err != nil {
			return err
		}

	}
	return nil
}

func (model *PortalTagPost) inDelArray(s int, kId []int) bool {

	for _, v := range kId {
		if s == v {
			return true
		}
	}
	return false

}

func (model *PortalTagPost) inAddArray(s int, tagPost []PortalTagPost) bool {

	for _, v := range tagPost {
		if s == v.Id {
			return true
		}
	}
	return false

}
