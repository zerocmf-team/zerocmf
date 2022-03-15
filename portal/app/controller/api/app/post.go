/**
** @创建时间: 2022/1/9 14:11
** @作者　　: return
** @描述　　:
 */

package app

import (
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/config"
	"github.com/gincmf/bootstrap/controller"
	bsModal "github.com/gincmf/bootstrap/model"
	"github.com/gincmf/bootstrap/paginate"
	"github.com/gincmf/bootstrap/util"
	"strconv"
	"strings"
)

type Post struct {
	controller.Rest
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取文章列表
 * @Date 2021/1/10 17:10:34
 * @Param
 * @return
 **/

func (rest *Post) Get(c *gin.Context) {

	var rewrite struct {
		Id int `uri:"id"`
	}

	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	categoryId := rewrite.Id

	db := util.GetDb(c)

	ids, err := new(model.PortalCategory).ChildIds(db, categoryId)

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	var query []string
	var queryArgs []interface{}

	for _, v := range ids {
		query = append(query, "cp.category_id = ?")
		queryArgs = append(queryArgs, v)
	}

	queryRes := []string{"p.post_type = 1 AND p.delete_at = 0"}

	queryStr := strings.Join(query, " OR ")
	queryRes = append(queryRes, "("+queryStr+")")

	current, pageSize, err := new(paginate.Paginate).Default(c)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	data, err := new(model.PortalPost).IndexByCategory(db, current, pageSize, strings.Join(queryRes, " AND "), queryArgs, nil)

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	rest.Success(c, "获取成功！", data)

}

func (rest *Post) ListWithCid(c *gin.Context) {

	var form struct {
		Ids []int `json:"ids"`
		Hot int   `json:"hot"` // 根据浏览量和权重排序
	}

	if err := c.BindJSON(&form); err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	extra := map[string]string{}

	hot := form.Hot

	if hot == 1 {
		extra["hot"] = "1"
	}

	var query []string
	var queryArgs []interface{}

	for _, v := range form.Ids {
		query = append(query, "cp.category_id = ?")
		queryArgs = append(queryArgs, v)
	}
	queryRes := []string{"p.post_type = 1 AND p.delete_at = 0"}
	if len(query) > 0 {
		queryStr := strings.Join(query, " OR ")
		queryRes = append(queryRes, queryStr)
	}
	db := util.GetDb(c)
	current, pageSize, err := new(paginate.Paginate).Default(c)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	data, err := new(model.PortalPost).IndexByCategory(db, current, pageSize, strings.Join(queryRes, " AND "), queryArgs, extra)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	rest.Success(c, "获取成功！", data)
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 显示单条文章
 * @Date 2022/2/12 18:41:40
 * @Param
 * @return
 **/

func (rest *Post) Show(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		rest.Error(c, "id不能为空", nil)
		return
	}
	typ := c.Query("type")
	if typ == "" {
		rest.Error(c, "页面类型错误", nil)
		return
	}
	db := util.GetDb(c)
	var query = "p.id = ? AND p.post_type = ? and p.delete_at = ?"
	var queryArgs = []interface{}{id, typ, 0}
	post, err := new(model.PortalPost).Show(db, query, queryArgs)

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	if post.Id == 0 {
		rest.Error(c, "该文章不存在或已被删除", nil)
		return
	}

	// 更新访问量 +1
	postHits := post.PostHits
	postHits += 1
	post.PostHits = postHits

	tx := db.Model(model.PortalPost{Id: post.Id}).Update("post_hits", postHits)
	if tx.Error != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	var result struct {
		model.PortalPost
		PrevPost *model.PortalPost `json:"prev_post"`
		NextPost *model.PortalPost `json:"next_post"`
	}

	result.PortalPost = post

	if typ == "1" {

		// 查询上一篇
		query = "p.id < ? AND p.post_type = ? and p.delete_at = ?"
		queryArgs = []interface{}{id, typ, 0}
		prevPost, err := new(model.PortalPost).Show(db, query, queryArgs)
		if err != nil {
			rest.Error(c, err.Error(), nil)
			return
		}

		// 查询下一篇
		query = "p.id > ? AND p.post_type = ? and p.delete_at = ?"
		queryArgs = []interface{}{id, 1, 0}
		nextPost, err := new(model.PortalPost).Show(db, query, queryArgs)
		if err != nil {
			rest.Error(c, err.Error(), nil)
			return
		}

		if prevPost.Id > 0 {
			result.PrevPost = &prevPost
		}

		if nextPost.Id > 0 {
			result.NextPost = &nextPost
		}
	}

	rest.Success(c, "获取成功！", result)
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 点赞单条文章
 * @Date 2022/2/12 18:41:53
 * @Param
 * @return
 **/

func (rest *Post) Like(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		rest.Error(c, "id不能为空", nil)
		return
	}

	db := util.GetDb(c)
	var query = "p.id = ? AND p.post_type = ? and p.delete_at = ?"
	var queryArgs = []interface{}{id, 1, 0}
	post, err := new(model.PortalPost).Show(db, query, queryArgs)

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	if post.Id == 0 {
		rest.Error(c, "该文章不存在或已被删除", nil)
		return
	}

	userId, _ := c.Get("userId")
	userIdInt, _ := strconv.Atoi(userId.(string))

	query = "post_id = ? AND user_id = ?"
	queryArgs = []interface{}{post.Id, userId}

	postLikePost := model.PostLikePost{}

	prefix := config.Config().Database.Prefix
	postLikePost.Table = prefix + "post_like_post"

	postLike, status, err := postLikePost.Like(db, bsModal.Post{
		Id:       post.Id,
		UserId:   userIdInt,
		PostLike: post.PostLike,
	}, query, queryArgs)

	msg := "点赞成功！"

	if status == false {
		msg = "取消点赞！"
	}

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	// 更新post_like
	tx := db.Where("id", post.Id).Model(&model.PortalPost{}).Update("post_like", postLike)

	if tx.Error != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	rest.Success(c, msg, post)

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 判断当前用户是否已经点赞
 * @Date 2022/2/23 12:14:34
 * @Param
 * @return
 **/

func (rest *Post) IsLike(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		rest.Error(c, "id不能为空", nil)
		return
	}

	db := util.GetDb(c)
	userId, _ := c.Get("userId")
	postLikePost := new(model.PostLikePost)
	prefix := config.Config().Database.Prefix
	postLikePost.Table = prefix + "post_like_post"

	err := postLikePost.IsLike(db, id, userId.(string))
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	rest.Success(c, "获取成功！", postLikePost.Status)

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 收藏单条文章
 * @Date 2022/2/12 18:41:53
 * @Param
 * @return
 **/

func (rest *Post) Favorite(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		rest.Error(c, "id不能为空", nil)
		return
	}

	db := util.GetDb(c)
	var query = "p.id = ? AND p.post_type = ? and p.delete_at = ?"
	var queryArgs = []interface{}{id, 1, 0}
	post, err := new(model.PortalPost).Show(db, query, queryArgs)

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	if post.Id == 0 {
		rest.Error(c, "该文章不存在或已被删除", nil)
		return
	}

	userId, _ := c.Get("userId")

	userIdInt, _ := strconv.Atoi(userId.(string))

	query = "post_id = ? AND user_id = ?"
	queryArgs = []interface{}{post.Id, userId}

	postFavoritesPost := model.PostFavoritesPost{}
	prefix := config.Config().Database.Prefix
	postFavoritesPost.Table = prefix + "post_favorites_post"

	postFavorite, status, err := postFavoritesPost.Like(db, bsModal.Post{
		Id:       post.Id,
		UserId:   userIdInt,
		PostLike: post.PostLike,
	}, query, queryArgs)

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	msg := "收藏成功！"

	if status == false {
		msg = "取消收藏！"
	}

	// 更新post_like
	tx := db.Where("id", post.Id).Model(&model.PortalPost{}).Update("post_favorites", postFavorite)

	if tx.Error != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	rest.Success(c, msg, post)

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 判断当前用户是否已经收藏
 * @Date 2022/2/23 12:14:34
 * @Param
 * @return
 **/

func (rest *Post) IsFavorite(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		rest.Error(c, "id不能为空", nil)
		return
	}

	db := util.GetDb(c)
	userId, _ := c.Get("userId")
	favoritesPost := new(model.PostFavoritesPost)

	prefix := config.Config().Database.Prefix
	favoritesPost.Table = prefix + "post_favorites_post"

	err := favoritesPost.IsLike(db, id, userId.(string))
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	rest.Success(c, "获取成功！", favoritesPost.Status)

}
