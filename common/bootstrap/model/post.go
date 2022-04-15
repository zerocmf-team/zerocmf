/**
** @创建时间: 2022/2/23 19:26
** @作者　　: return
** @描述　　: 各种页面的点赞收藏评论
 */

package model

import (
	"gincmf/common/bootstrap/util"
	"gorm.io/gorm"
	"time"
)

/**
 * @Author return <1140444693@qq.com>
 * @Description 文章信息
 * @Date 2022/2/23 21:24:46
 * @Param
 * @return
 **/

type Post struct {
	Id       int `json:"id"`
	UserId   int `json:"user_id"`
	PostLike int `json:"post_like"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 点赞收藏表
 * @Date 2022/2/23 20:20:17
 * @Param
 * @return
 **/

type LikePost struct {
	Table      string `gorm:"-" json:"table"`
	Id         int    `json:"id"`
	PostId     int    `gorm:"type:bigint(20);comment:对象id，例如：文章，评论等;not null" json:"post_id"`
	UserId     int    `gorm:"type:bigint(20);comment:用户id;not null" json:"user_id"`
	Status     int    `gorm:"type:tinyint(3);comment:状态,1:点赞;0:未点赞;default:1;not null" json:"status"`
	CreateAt   int64  `gorm:"type:bigint(20);NOT NULL" json:"create_at"`
	UpdateAt   int64  `gorm:"type:bigint(20);NOT NULL" json:"update_at"`
	CreateTime string `gorm:"-" json:"create_time"`
	UpdateTime string `gorm:"-" json:"update_time"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 点赞收藏内容
 * @Date 2022/2/23 19:31:36
 * @Param
 * @return
 **/

func (model *LikePost) Like(db *gorm.DB, post Post, query string, queryArgs []interface{}) (postLike int, status bool, err error) {

	err = model.Show(db, query, queryArgs)
	if err != nil {
		return
	}

	now := time.Now().Unix()

	// 不存在点赞
	savePost := LikePost{
		PostId:   post.Id,
		UserId:   post.UserId,
		Status:   1,
		CreateAt: now,
		UpdateAt: now,
	}

	var tx *gorm.DB

	status = true
	postLike = post.PostLike

	if model.Id == 0 {
		postLike += 1
		tx = db.Table(model.Table).Create(&savePost)
	} else {
		savePost.Id = model.Id
		if model.Status == 1 {
			savePost.Status = 0
			postLike -= 1
			status = false
		} else {
			savePost.Status = 1
			postLike += 1
		}
		savePost.UpdateAt = now
		tx = db.Table(model.Table).Save(&savePost)
	}

	if tx.Error != nil {
		err = tx.Error
		return
	}

	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查看单条点赞收藏关系
 * @Date 2022/2/23 12:52:55
 * @Param
 * @return
 **/

func (model *LikePost) Show(db *gorm.DB, query string, queryArgs []interface{}) (err error) {
	tx := db.Table(model.Table).Where(query, queryArgs...).Scan(&model)
	if util.IsDbErr(tx) != nil {
		err = tx.Error
		return
	}
	return nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 判断当前用户是否已经喜欢或收藏
 * @Date 2022/2/23 12:53:10
 * @Param
 * @return
 **/

func (model *LikePost) IsLike(db *gorm.DB, postId string, userId string) (err error) {
	query := "post_id = ? AND user_id = ?"
	queryArgs := []interface{}{postId, userId}
	err = model.Show(db, query, queryArgs)
	return err
}

/**
 * @Author return <1140444693@qq.com>
 * @Description
 * @Date 2022/2/26 18:13:59
 * @Param
 * @return
 **/
