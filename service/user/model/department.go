package model

import (
	"gorm.io/gorm"
	"strconv"
	"time"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/util"
)

type Department struct {
	Id         int     `json:"id"`
	ParentId   int     `gorm:"type:int(11);default:0;comment:父级id" json:"parent_id"`
	Name       string  `gorm:"type:varchar(30);comment:名称" json:"name"`
	Leader     string  `gorm:"type:varchar(30);comment:负责人" json:"leader"`
	Mobile     string  `gorm:"type:varchar(20);comment:联系电话" json:"mobile"`
	Email      string  `gorm:"type:varchar(100);comment:联系邮箱" json:"email"`
	Status     int     `gorm:"type:tinyint(3);default:1;comment:文件状态" json:"status"`
	ListOrder  float64 `gorm:"type:float;default:10000;comment:排序（越大越靠前）" json:"list_order" validate:"required" label:"排序"`
	CreateAt   int64   `gorm:"type:bigint(20);NOT NULL" json:"create_at"`
	CreateTime string  `gorm:"-" json:"create_time"`
	UpdateAt   int64   `gorm:"type:bigint(20);comment:更新时间" json:"update_at"`
	UpdateTime string  `gorm:"-" json:"update_time"`
}

type department struct {
	Department
	Path     string       `json:"path"`
	Children []department `json:"children"`
}

func (_ *Department) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Department{})
}

/**
 * @Author return
 * @Description 递归获取树形表格数据
 * @Date 2022/12/2 16:39
 * @Param
 * @return
 **/

func (rest *Department) TreeList(db *gorm.DB, query string, queryArgs []interface{}) (result []department, err error) {
	var dep []Department
	tx := db.Where(query, queryArgs...).Find(&dep)
	if util.IsDbErr(tx) != nil {
		err = tx.Error
		return
	}

	result = recursionDepartment(dep, 0, "")

	return
}

func recursionDepartment(depInArr []Department, parentId int, parentIndex string) []department {
	var resArr []department
	index := 0
	for _, v := range depInArr {
		if parentId == v.ParentId {
			iStr := strconv.Itoa(index)
			var curIndex string
			if parentIndex == "" {
				curIndex = iStr
			} else {
				curIndex = parentIndex + "-" + iStr
			}

			v.CreateTime = time.Unix(v.CreateAt, 0).Format(data.TimeLayout)
			v.UpdateTime = time.Unix(v.UpdateAt, 0).Format(data.TimeLayout)

			result := department{
				Department: v,
				Path:       curIndex,
			}
			index++
			children := recursionDepartment(depInArr, v.Id, curIndex)
			result.Children = children
			resArr = append(resArr, result)
		}
	}
	return resArr
}

func (rest *Department) Paginate(db *gorm.DB, current, pageSize int, query string, queryArgs []interface{}) (result data.Paginate, err error) {
	var department []Department
	var total int64 = 0
	tx := db.Where(query, queryArgs...).Find(&department).Count(&total)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	tx = db.Limit(pageSize).Where(query, queryArgs...).Offset((current - 1) * pageSize).Find(&department)
	if util.IsDbErr(tx) != nil {
		err = tx.Error
		return
	}

	result = data.Paginate{Data: department, Current: current, PageSize: pageSize, Total: total}
	if len(department) == 0 {
		result.Data = make([]string, 0)
	}
	return result, nil
}
