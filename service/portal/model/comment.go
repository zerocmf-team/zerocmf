/**
** @创建时间: 2022/2/24 09:03
** @作者　　: return
** @描述　　:
 */

package model

import (
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/model"
)

type Comment struct {
}

func (c *Comment) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.CommentReply{})
	db.AutoMigrate(&model.CommentLikePost{})
	db.AutoMigrate(&model.CommentReplyLikePost{})
}
