/**
** @创建时间: 2022/2/24 09:03
** @作者　　: return
** @描述　　:
 */

package model

import (
	"github.com/gincmf/bootstrap/model"
	"gorm.io/gorm"
)

type Comment struct {

}

func (c *Comment) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.CommentReply{})
	db.AutoMigrate(&model.CommentLikePost{})
	db.AutoMigrate(&model.CommentReplyLikePost{})
}
