/**
** @创建时间: 2022/2/24 12:09
** @作者　　: return
** @描述　　:
 */

package app

import (
	"gincmf/app/grpc/user"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/config"
	"github.com/gincmf/bootstrap/controller"
	"github.com/gincmf/bootstrap/model"
	"github.com/gincmf/bootstrap/paginate"
	"github.com/gincmf/bootstrap/util"
	"github.com/jinzhu/copier"
	"strconv"
	"time"
)

type Comment struct {
	controller.Rest
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取当前文章全部评论列表
 * @Date 2022/2/24 12:12:44
 * @Param
 * @return
 **/

func (rest *Comment) Get(c *gin.Context) {

	topicId := c.Param("id")
	typ := c.DefaultQuery("type", "0")

	db := util.GetDb(c)

	query := "topic_id = ? AND topic_type = ?"
	queryArgs := []interface{}{topicId, typ}

	userId := c.Query("userId")
	userIdInt := 0

	if userId != "" {
		tenantId := util.TenantId(c)
		uidInt, _ := strconv.Atoi(userId)
		tenantIdInt, _ := strconv.Atoi(tenantId)

		userData, err := new(user.User).Request(uidInt, tenantIdInt)
		if err != nil {
			rest.Error(c, err.Error(), nil)
			return
		}

		userIdInt = int(userData.Id)
	}

	current, pageSize, err := new(paginate.Paginate).Default(c)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	data, err := new(model.Comment).Paginate(db, current, pageSize, query, queryArgs, userIdInt)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	rest.Success(c, "获取成功！", data)

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 评论文章
 * @Date 2022/2/24 12:15:49
 * @Param
 * @return
 **/

func (rest *Comment) Comment(c *gin.Context) {

	topicId := c.Param("id")
	topicIdInt, _ := strconv.Atoi(topicId)

	var form struct {
		TopicType int    `json:"topic_type"`
		Content   string `json:"content"`
	}

	if err := c.ShouldBindJSON(&form); err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	userId, _ := c.Get("userId")
	tenantId := util.TenantId(c)

	userIdInt, _ := strconv.Atoi(userId.(string))
	tenantIdInt, _ := strconv.Atoi(tenantId)

	userData, err := new(user.User).Request(userIdInt, tenantIdInt)

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	db := util.GetDb(c)

	now := time.Now().Unix()
	comment := model.Comment{
		TopicId:          topicIdInt,
		TopicType:        form.TopicType,
		Content:          form.Content,
		FromUserId:       int(userData.Id),
		FromUserNickname: userData.UserNickname,
		FromUserAvatar:   userData.Avatar,
	}

	comment.CreateAt = now
	comment.UpdateAt = now

	tx := db.Create(&comment)
	if tx.Error != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	rest.Success(c, "评论成功！", comment)

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 回复评论
 * @Date 2022/2/25 12:35:43
 * @Param
 * @return
 **/

func (rest *Comment) Reply(c *gin.Context) {

	commentId := c.Param("id")
	commentIdInt, _ := strconv.Atoi(commentId)

	var form struct {
		ReplyId   int    `json:"reply_id"`
		ReplyType int    `json:"reply_type"`
		Content   string `json:"content"`
		ToUserId  int    `json:"to_user_id"`
	}

	if err := c.ShouldBindJSON(&form); err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	userId, _ := c.Get("userId")
	tenantId := util.TenantId(c)

	userIdInt, _ := strconv.Atoi(userId.(string))
	tenantIdInt, _ := strconv.Atoi(tenantId)

	userData, err := new(user.User).Request(userIdInt, tenantIdInt)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	toUserId := form.ToUserId
	var toUserData *user.Data

	db := util.GetDb(c)
	now := time.Now().Unix()
	reply := model.CommentReply{}
	copier.Copy(&reply, &form)

	reply.CommentId = commentIdInt
	reply.FromUserId = int(userData.Id)
	reply.FromUserNickname = userData.UserNickname
	reply.FromUserAvatar = userData.Avatar
	reply.CreateAt = now
	reply.UpdateAt = now

	if toUserId != 0 {
		toUserData, err = new(user.User).Request(toUserId, tenantIdInt)
		if err != nil {
			rest.Error(c, err.Error(), nil)
			return
		}
		reply.ToUserId = int(toUserData.Id)
		reply.ToUserNickname = toUserData.UserNickname
		reply.ToUserAvatar = toUserData.Avatar
	}

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	tx := db.Create(&reply)
	err = tx.Error
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	rest.Success(c, "回复成功！", reply)

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 点赞单条评论
 * @Date 2022/2/12 18:41:53
 * @Param
 * @return
 **/

func (rest *Comment) Like(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		rest.Error(c, "id不能为空", nil)
		return
	}

	db := util.GetDb(c)
	var query = "id = ? AND status = 1 AND delete_at = 0"
	var queryArgs = []interface{}{id}
	comment := new(model.Comment)
	err := comment.Show(db, query, queryArgs)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	if comment.Id == 0 {
		rest.Error(c, "该评论不存在或已被删除", nil)
		return
	}
	userId, _ := c.Get("userId")
	userIdInt, _ := strconv.Atoi(userId.(string))

	query = "post_id = ? AND user_id = ?"
	queryArgs = []interface{}{comment.Id, userId}

	postLikePost := model.CommentLikePost{}

	prefix := config.Config().Database.Prefix
	postLikePost.Table = prefix + "comment_like_post"

	postLike, status, err := postLikePost.Like(db, model.Post{
		Id:       comment.Id,
		UserId:   userIdInt,
		PostLike: comment.PostLike,
	}, query, queryArgs)

	msg := "点赞成功！"
	comment.IsLike = 1

	if status == false {
		msg = "取消点赞！"
		comment.IsLike = 0
	}

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	// 更新post_like
	tx := db.Where("id", comment.Id).Model(&model.Comment{}).Update("post_like", postLike)

	comment.PostLike = postLike

	if tx.Error != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	rest.Success(c, msg, comment)

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 点赞单条回复评论
 * @Date 2022/2/27 14:32:44
 * @Param
 * @return
 **/

func (rest *Comment) ReplyLike(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		rest.Error(c, "id不能为空", nil)
		return
	}

	db := util.GetDb(c)
	var query = "id = ? AND status = 1 AND delete_at = 0"
	var queryArgs = []interface{}{id}
	commentReply := new(model.CommentReply)
	err := commentReply.Show(db, query, queryArgs)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	if commentReply.Id == 0 {
		rest.Error(c, "该回复不存在或已被删除", nil)
		return
	}
	userId, _ := c.Get("userId")
	userIdInt, _ := strconv.Atoi(userId.(string))

	query = "post_id = ? AND user_id = ?"
	queryArgs = []interface{}{commentReply.Id, userId}

	postLikePost := model.CommentReplyLikePost{}

	prefix := config.Config().Database.Prefix
	postLikePost.Table = prefix + "comment_reply_like_post"

	postLike, status, err := postLikePost.Like(db, model.Post{
		Id:       commentReply.Id,
		UserId:   userIdInt,
		PostLike: commentReply.PostLike,
	}, query, queryArgs)

	msg := "点赞成功！"
	commentReply.IsLike = 1

	if status == false {
		msg = "取消点赞！"
		commentReply.IsLike = 0
	}

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	// 更新post_like
	tx := db.Where("id", commentReply.Id).Model(&model.CommentReply{}).Update("post_like", postLike)

	commentReply.PostLike = postLike

	if tx.Error != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	rest.Success(c, msg, commentReply)

}
