/**
** @创建时间: 2020/11/25 1:59 下午
** @作者　　: return
** @描述　　:
 */

package admin

import (
	"encoding/json"
	"errors"
	"gincmf/app/model"
	"gincmf/app/service"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	cmfModel "github.com/gincmf/bootstrap/model"
	"github.com/gincmf/bootstrap/util"
	"gorm.io/gorm"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type PortalPost struct {
	controller.Rest
}

// 获取文章列表

func (rest *PortalPost) Get(c *gin.Context) {

	query := []string{"p.delete_at = ?"}
	queryArgs := []interface{}{0}

	title := c.Query("post_title")
	if title != "" {
		query = append(query, "p.post_title like ?")
		queryArgs = append(queryArgs, "%"+title+"%")
	}

	postType := c.Query("post_type")
	if postType == "2" {
		postType = "2"
	} else {
		postType = "1"
	}

	query = append(query, "p.post_type = ?")
	queryArgs = append(queryArgs, postType)

	startTime := c.Query("start_time")
	endTime := c.Query("end_time")

	if startTime != "" && endTime != "" {
		startTimeStamp, err := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
		if err != nil {
			rest.Error(c, "起始时间非法！", err.Error())
			return
		}

		endTimeStamp, err := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
		if err != nil {
			rest.Error(c, "结束时间非法！", err.Error())
		}

		query = append(query, "((p.publish_at BETWEEN ? AND ?) OR (p.update_at BETWEEN ? AND ?))")
		queryArgs = append(queryArgs, startTimeStamp, endTimeStamp, startTimeStamp, endTimeStamp)
	}

	queryStr := strings.Join(query, " AND ")

	post := new(service.PortalPost)
	data, err := post.IndexByCategory(c, queryStr, queryArgs)

	if err != nil {
		rest.Error(c, "获取失败！", nil)
		return
	}

	rest.Success(c, "获取成功！", data)
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查看单条文章列表
 * @Date 2021/12/12 18:28:46
 * @Param
 * @return
 **/

func (rest *PortalPost) Show(c *gin.Context) {

	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	db := util.GetDb(c)
	query := []string{"id = ?", "delete_at = ?"}
	queryStr := strings.Join(query, " AND ")
	queryArgs := []interface{}{rewrite.Id, 0}

	var result struct {
		model.PortalPost
		UserLogin string                 `json:"user_login"`
		Alias     string                 `json:"alias"`
		Keywords  []string               `json:"keywords"`
		Category  []model.PortalCategory `json:"category"`
		Extends   []model.Extends        `json:"extends"`
		Slug      string                 `json:"slug"`
		model.More
	}

	post, err := new(model.PortalPost).Show(db, queryStr, queryArgs)

	if err != nil {
		rest.Error(c, "查询失败："+err.Error(), nil)
		return
	}

	if post.PostKeywords != "" {
		result.Keywords = strings.Split(post.PostKeywords, ",")
	}

	result.PortalPost = post

	pQueryArgs := []interface{}{rewrite.Id, 0}
	pc := model.PortalCategory{}
	category, err := pc.ListWithPost(db, "p.id = ? AND p.delete_at = ?", pQueryArgs)

	result.Category = category

	result.Extends = post.MoreJson.Extends

	result.Photos = post.MoreJson.Photos
	result.Files = post.MoreJson.Files

	result.Audio = post.MoreJson.Audio
	result.AudioPrevPath = post.MoreJson.AudioPrevPath

	result.Video = post.MoreJson.Video
	result.VideoPrevPath = post.MoreJson.VideoPrevPath

	fullUrl := "page/" + strconv.Itoa(rewrite.Id)
	route := cmfModel.Route{}
	tx := db.Where("full_url", fullUrl).First(&route)

	if util.IsDbErr(tx) != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	result.Alias = route.Url

	rest.Success(c, "获取成功！", result)
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 新增文章
 * @Date 2021/12/16 21:59:59
 * @Param
 * @return
 **/

func (rest *PortalPost) Store(c *gin.Context) {
	rest.Save(c, 0)
}

func (rest *PortalPost) Edit(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	rest.Save(c, rewrite.Id)
}

func (rest *PortalPost) Save(c *gin.Context, id int) {

	var form struct {
		CategoryIds    []string        `json:"category_ids"`
		PostType       int             `json:"post_type"`
		Alias          string          `json:"alias"`
		PostTitle      string          `json:"post_title"`
		Thumbnail      string          `json:"thumbnail"`
		PostKeywords   []string        `json:"post_keywords"`
		ListOrder      float64         `json:"list_order"`
		PublishTime    string          `json:"publish_time"`
		PostSource     string          `json:"post_source"`
		PostExcerpt    string          `json:"post_excerpt"`
		PostContent    string          `json:"post_content"`
		IsTop          int             `json:"is_top"`
		SeoTitle       string          `json:"seo_title"`
		SeoKeywords    string          `json:"seo_keywords"`
		SeoDescription string          `json:"seo_description"`
		Recommended    int             `json:"recommended"`
		PostStatus     int             `json:"post_status"`
		Photos         []model.Path    `json:"photos"`
		Files          []model.Path    `json:"files"`
		Audio          string          `json:"audio"`
		Video          string          `json:"video"`
		Template       string          `json:"template"`
		Extends        []model.Extends `json:"extends"`
	}

	err := c.ShouldBind(&form)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	postType := 1
	if form.PostType == 2 {
		postType = 2
	}

	if postType == 1 && len(form.CategoryIds) == 0 {
		rest.Error(c, "分类不能为空！", nil)
		return
	}

	if form.PostTitle == "" {
		rest.Error(c, "标题不能为空！", nil)
		return
	}

	currentTime := time.Now().Unix()

	var more model.More

	more.Photos = form.Photos
	more.Files = form.Files
	more.Audio = form.Audio
	more.Video = form.Video
	more.Extends = form.Extends

	re, _ := regexp.Compile("[^a-zA-Z0-9]+")
	slug := re.ReplaceAllString(form.PostTitle, "-")
	more.Slug = strings.ToLower(slug)

	if form.Template != "" {
		more.Template = form.Template
	}

	moreJson, err := json.Marshal(more)

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	userId, _ := c.Get("userId")
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	userIdInt, _ := strconv.Atoi(userId.(string))

	publishAt, _ := time.ParseInLocation(model.TimeLayout, form.PublishTime, time.Local)

	listOrder := form.ListOrder

	if listOrder == 0 {
		listOrder = 10000
	}

	portal := model.PortalPost{
		PostType:       postType,
		PostTitle:      form.PostTitle,
		SeoTitle:       form.SeoTitle,
		SeoKeywords:    form.SeoKeywords,
		SeoDescription: form.SeoDescription,
		Thumbnail:      form.Thumbnail,
		UserId:         userIdInt,
		PostKeywords:   strings.Join(form.PostKeywords, ","),
		PostSource:     form.PostSource,
		PostExcerpt:    form.PostExcerpt,
		PostContent:    form.PostContent,
		CreateAt:       currentTime,
		UpdateAt:       currentTime,
		IsTop:          form.IsTop,
		Recommended:    form.Recommended,
		PostStatus:     form.PostStatus,
		More:           string(moreJson),
		PublishedAt:    publishAt.Unix(),
		ListOrder:      listOrder,
	}

	db := util.GetDb(c)

	var data struct {
		model.PortalPost
		Category []model.PortalCategoryPost `json:"category"`
	}

	var (
		postData model.PortalPost
	)

	if id == 0 {
		postData, err = portal.Store(db)
	} else {
		portal.Id = id
		postData, err = portal.Update(db)
	}

	if form.Alias != "" {
		fullUrl := "page/" + strconv.Itoa(postData.Id)
		url := form.Alias

		route := cmfModel.Route{
			Type:    2,
			FullUrl: fullUrl,
			Url:     url,
		}

		err = route.Set(db)
		if err != nil {
			rest.Error(c, err.Error(), nil)
			return
		}
	}

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	var pcpPost = make([]model.PortalCategoryPost, 0)

	pcp := model.PortalCategoryPost{}

	category := model.PortalCategory{}
	existsCategory, err := category.List(db)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			rest.Error(c, "请先添加分类！", nil)
			return
		}
		rest.Error(c, err.Error(), nil)
		return
	}

	for _, v := range form.CategoryIds {

		cidInt, _ := strconv.Atoi(v)

		if !rest.inArray(cidInt, existsCategory) {
			rest.Error(c, "分类参数非法！", nil)
			return
		}

		pcp.PostId = postData.Id
		pcp.CategoryId = cidInt
		pcpPost = append(pcpPost, pcp)
	}

	pcpData, pcpErr := pcp.Store(db, pcpPost)
	if pcpErr != nil {
		rest.Error(c, pcpErr.Error(), nil)
		return
	}

	data.Category = pcpData

	var tag []int
	for _, v := range form.PostKeywords {

		if strings.TrimSpace(v) != "" {
			// 查询当前tag是否存在
			portalTag, err := new(model.PortalTag).Show(db, "name = ?", []interface{}{v})
			if err != nil {
				rest.Error(c, err.Error(), nil)
				return
			}

			portalTag.Name = v

			_, err = portalTag.FirstOrSave(db)
			if err != nil {
				rest.Error(c, err.Error(), err)
			}

			tag = append(tag, portalTag.Id)
		}

	}

	tagPost := model.PortalTagPost{
		PostId: postData.Id,
	}
	tagPost.FirstOrSave(db, tag)
	data.PortalPost = postData
	message := "添加成功！"
	if id > 0 {
		message = "更新成功！"
	}
	rest.Success(c, message, data)

}

func (rest *PortalPost) inArray(cid int, pc []model.PortalCategory) bool {
	for _, v := range pc {
		if v.Id == cid {
			return true
		}
	}
	return false
}
