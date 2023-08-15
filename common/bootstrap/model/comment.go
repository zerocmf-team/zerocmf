/**
** @创建时间: 2022/2/26 18:14
** @作者　　: return
** @描述　　:
 */

package model

import (
	"errors"
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/util"
)

/**
 * @Author return <1140444693@qq.com>
 * @Description 评论表
 * @Date 2022/2/23 21:25:3
 * @Param
 * @return
 **/

type Comment struct {
	Id               int    `json:"id"`
	TopicId          int    `gorm:"type:bigint(20);comment:主题id;not null" json:"topic_id"`
	TopicType        int    `gorm:"type:tinyint(3);comment:类型,0:文章;default:0;not null" json:"topic_type"`
	Content          string `gorm:"type:varchar(500);comment:评论内容" json:"content"`
	FromUserId       int    `gorm:"type:bigint(20);comment:用户id;not null" json:"from_user_id"`
	FromUserNickname string `gorm:"type:varchar(50);comment:用户昵称;not null" json:"from_user_nickname"`
	FromUserAvatar   string `gorm:"type:varchar(255);comment:用户头像;not null" json:"from_user_avatar"`
	Status           int    `gorm:"type:tinyint(3);comment:状态,1:上架;0:审核中;2：拒绝;default:1;not null" json:"status"`
	DeleteAt         int64  `gorm:"type:bigint(20);NOT NULL" json:"delete_at"`
	PostLike         int    `gorm:"type:int(11);comment:点赞数;default:0;NOT NULL" json:"post_like"`
	IsLike           int    `gorm:"-" json:"is_like"`
	Time
	CommentReply     []CommentReply `gorm:"-" json:"comment_reply"`
	ReplyAllCount    int            `gorm:"-" json:"reply_all_count"`
	ReplyOthersCount int            `gorm:"-" json:"reply_others_count"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 回复表
 * @Date 2022/2/23 21:57:0
 * @Param
 * @return
 **/

type CommentReply struct {
	Id               int    `json:"id"`
	CommentId        int    `gorm:"type:bigint(20);comment:评论id;not null" json:"comment_id"`
	ReplyId          int    `gorm:"type:bigint(20);comment:回复目标id;没有回复内容则为空;not null" json:"reply_id"`
	ReplyType        int    `gorm:"type:tinyint(3);comment:类型（0:评论，1：回复）;default:0;not null" json:"reply_type"`
	Content          string `gorm:"type:varchar(500);comment:回复内容" json:"content"`
	PostLike         int    `gorm:"type:int(11);comment:点赞数;default:0;NOT NULL" json:"post_like"`
	FromUserId       int    `gorm:"type:bigint(20);comment:用户id;not null" json:"from_user_id"`
	FromUserNickname string `gorm:"type:varchar(50);comment:用户昵称;not null" json:"from_user_nickname"`
	FromUserAvatar   string `gorm:"type:varchar(255);comment:用户头像;not null" json:"from_user_avatar"`
	ToUserId         int    `gorm:"type:bigint(20);comment:用户id;not null" json:"to_user_id"`
	ToUserNickname   string `gorm:"type:varchar(50);comment:用户昵称;not null" json:"to_user_nickname"`
	ToUserAvatar     string `gorm:"type:varchar(255);comment:用户头像;not null" json:"to_user_avatar"`
	Status           int    `gorm:"type:tinyint(3);comment:状态,1:上架;0:审核中;2：拒绝;default:1;not null" json:"status"`
	DeleteAt         int64  `gorm:"type:bigint(20);NOT NULL" json:"delete_at"`
	IsLike           int    `gorm:"-" json:"is_like"`
	Time
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 评论点赞表
 * @Date 2022/2/27 9:8:4
 * @Param
 * @return
 **/

type CommentLikePost struct {
	LikePost
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 回复点赞表
 * @Date 2022/2/27 14:28:24
 * @Param
 * @return
 **/

type CommentReplyLikePost struct {
	LikePost
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取分页评论列表
 * @Date 2022/2/26 18:15:33
 * @Param
 * @return
 **/

type CommentData struct {
	CommentAllCount int `json:"comment_all_count"`
	data.Paginate
}

func (model *Comment) Paginate(db *gorm.DB, current int, pageSize int, query string, queryArgs []interface{}, userId int) (Data CommentData, err error) {

	var total int64 = 0
	var CommentAllCount = 0
	var comment []Comment

	tx := db.Where(query, queryArgs...).Find(&comment).Count(&total)
	if err = util.IsDbErr(tx); err != nil {
		return Data, errors.New("数据库连接出错：" + err.Error())
	}

	tx = db.Where(query, queryArgs...).Limit(pageSize).Offset((current - 1) * pageSize).Order("id desc").Find(&comment)

	if err = util.IsDbErr(tx); err != nil {
		return Data, errors.New("数据库连接出错：" + err.Error())
	}

	CommentAllCount += int(total)

	for k, v := range comment {
		// 查询当前用户是否已经点赞
		if userId > 0 {
			var cPost CommentLikePost
			db.Where("user_id = ? AND post_id = ?", userId, v.Id).First(&cPost)
			if cPost.Status == 1 {
				comment[k].IsLike = 1
			}
		}

		//	查出两条评论和回复
		var commentReply []CommentReply
		db.Where("comment_id = ?", []interface{}{v.Id}).Limit(2).Order("create_at").Find(&commentReply)

		for rk, rv := range commentReply {
			// 查询当前用户是否已经点赞
			if userId > 0 {
				var crPost CommentReplyLikePost
				db.Where("user_id = ? AND post_id = ?", userId, rv.Id).First(&crPost)
				if crPost.Status == 1 {
					commentReply[rk].IsLike = 1
				}
			}
		}

		var total int64 = 0
		db.Where("comment_id = ?", []interface{}{v.Id}).Find(&model.CommentReply).Count(&total)

		if len(commentReply) > 0 {
			CommentAllCount += int(total)
			comment[k].CommentReply = commentReply
			comment[k].ReplyAllCount = int(total)
			if total > 2 {
				comment[k].ReplyOthersCount = int(total - 2)
			}
		}
	}

	Data = CommentData{
		CommentAllCount: CommentAllCount,
		Paginate: data.Paginate{
			Data: comment, Current: current, PageSize: pageSize, Total: total,
		},
	}
	return

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取当前单条评论
 * @Date 2022/2/27 9:13:20
 * @Param
 * @return
 **/

func (model *Comment) Show(db *gorm.DB, query string, queryArgs []interface{}) (err error) {
	tx := db.Where(query, queryArgs...).First(&model)
	if util.IsDbErr(tx) != nil {
		return tx.Error
	}
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取当前单条回复
 * @Date 2022/2/27 9:13:20
 * @Param
 * @return
 **/

func (model *CommentReply) Show(db *gorm.DB, query string, queryArgs []interface{}) (err error) {
	tx := db.Where(query, queryArgs...).First(&model)
	if util.IsDbErr(tx) != nil {
		return tx.Error
	}
	return
}
