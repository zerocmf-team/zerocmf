/**
** @创建时间: 2020/11/25 2:01 下午
** @作者　　: return
** @描述　　:
 */

package model

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"time"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/model"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/user/rpc/types/user"
)

type PortalPost struct {
	Id                  int              `json:"id"`
	ParentId            int              `gorm:"type:int(11);comment:父级id;NOT NULL" json:"parent_id"`
	PostType            int              `gorm:"type:tinyint(3);comment:类型（1:文章，2:页面）;default:1;NOT NULL" json:"post_type"`
	PostFormat          int              `gorm:"type:tinyint(3);comment:内容格式（1:html，2:md）;default:1;NOT NULL" json:"post_format"`
	UserId              int              `gorm:"type:int(11);comment:发表者用户id;NOT NULL" json:"user_id"`
	UserLogin           string           `gorm:"type:varchar(60);comment:登录账号" json:"user_login"`
	PostStatus          int              `gorm:"type:tinyint(3);comment:状态（1:已发布，0:未发布）;default:1;NOT NULL" json:"post_status"`
	CommentStatus       int              `gorm:"type:tinyint(3);comment:评论状态（1:允许，0:不允许）;default:1;NOT NULL" json:"comment_status"`
	IsTop               int              `gorm:"type:tinyint(3);comment:是否置顶（1:置顶，0:不置顶）;default:0;NOT NULL" json:"is_top"`
	Recommended         int              `gorm:"type:tinyint(3);comment:是否推荐（1:推荐，0:不推荐）;default:0;NOT NULL" json:"recommended"`
	PostHits            int              `gorm:"type:int(11);comment:查看数;default:0;NOT NULL" json:"post_hits"`
	PostFavorites       int              `gorm:"type:int(11);comment:收藏数;default:0;NOT NULL" json:"post_favorites"`
	PostLike            int              `gorm:"type:int(11);comment:点赞数;default:0;NOT NULL" json:"post_like"`
	CommentCount        int              `gorm:"type:int(11);comment:评论数;default:0;NOT NULL" json:"comment_count"`
	CreateAt            int64            `gorm:"type:bigint(20);NOT NULL" json:"create_at"`
	UpdateAt            int64            `gorm:"type:bigint(20);NOT NULL" json:"update_at"`
	PublishedAt         int64            `gorm:"type:bigint(20);comment:发布时间;NOT NULL" json:"published_at"`
	DeleteAt            int64            `gorm:"type:bigint(20);comment:删除实际;NOT NULL" json:"delete_at"`
	PostTitle           string           `gorm:"type:varchar(100);comment:post标题;NOT NULL" json:"post_title"`
	PostKeywords        string           `gorm:"type:varchar(150);comment:SEO关键词;NOT NULL" json:"post_keywords"`
	PostExcerpt         string           `gorm:"type:longtext;comment:post摘要;NOT NULL" json:"post_excerpt"`
	ListOrder           float64          `gorm:"type:double;comment:排序;default:10000;NOT NULL" json:"list_order"`
	PostSource          string           `gorm:"type:varchar(500);comment:转载文章的来源;NOT NULL" json:"post_source"`
	SeoTitle            string           `gorm:"type:varchar(100);comment:三要素标题;not null" json:"seo_title"`
	SeoKeywords         string           `gorm:"type:varchar(255);comment:三要素关键字;not null" json:"seo_keywords"`
	SeoDescription      string           `gorm:"type:varchar(255);comment:三要素描述;not null" json:"seo_description"`
	Thumbnail           string           `gorm:"type:varchar(100);comment:缩略图;NOT NULL" json:"thumbnail"`
	ThumbPrevPath       string           `gorm:"-" json:"thumb_prev_path"`
	PostContent         string           `gorm:"type:longtext;comment:文章内容;NOT NULL" json:"post_content"`
	PostContentFiltered string           `gorm:"type:longtext;comment:处理过的文章内容;NOT NULL" json:"post_content_filtered"`
	More                string           `gorm:"type:json;comment:扩展属性,如缩略图。格式为json;NOT NULL" json:"more"`
	MoreJson            More             `gorm:"-" json:"more_json"`
	Category            []PortalCategory `gorm:"-" json:"category"`
	Tags                []PostTagResult  `gorm:"-" json:"tags"`
	Template            string           `gorm:"-" json:"template"`
	CreateTime          string           `gorm:"-" json:"create_time"`
	UpdateTime          string           `gorm:"-" json:"update_time"`
	PublishedTime       string           `gorm:"-" json:"published_time"`
	DeleteTime          string           `gorm:"-" json:"delete_time"`
	userRpc             user.UserClient
}

type More struct {
	Photos        []Path            `json:"photos"`
	Files         []Path            `json:"files"`
	Extends       []Extends         `json:"extends"`
	ExtendsObj    map[string]string `json:"extends_obj"`
	Audio         string            `json:"audio"`
	AudioPrevPath string            `json:"audio_prev_path"`
	Video         string            `json:"video"`
	VideoPrevPath string            `json:"video_prev_path"`
	Template      string            `json:"template"`
	Alias         string            `json:"alias"`
}

type Extends struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Path struct {
	RemarkName string `json:"remark_name"`
	PrevPath   string `json:"prev_path"`
	FilePath   string `json:"file_path"`
}

type PortalCategoryResult struct {
	Id         int    `json:"id"`
	Name       string `gorm:"type:varchar(200);comment:唯一名称;not null" json:"name"`
	Alias      string `gorm:"type:varchar(200);comment:唯一名称;not null" json:"alias"`
	PostId     int    `gorm:"type:bigint(20);comment:文章id;not null" json:"post_id"`
	CategoryId int    `gorm:"type:int(11);comment:分类id;not null" json:"category_id"`
}

// 分类关系表

type PortalCategoryPost struct {
	Id         int     `json:"id"`
	PostId     int     `gorm:"type:int(11);comment:文章id;not null" json:"post_id"`
	CategoryId int     `gorm:"type:int(11);comment:分类id;not null" json:"category_id"`
	ListOrder  float64 `gorm:"type:float(0);comment:排序;default:10000;not null" json:"list_order"`
	Status     int     `gorm:"type:tinyint(3);comment:状态,1:发布;0:不发布;default:1;not null" json:"status"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 点赞关系表
 * @Date 2022/2/12 17:31:17
 * @Param
 * @return
 **/

type PostLikePost struct {
	model.LikePost
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 收藏关系表
 * @Date 2022/2/12 17:31:17
 * @Param
 * @return
 **/

type PostFavoritesPost struct {
	model.LikePost
}

func NewPost(userRpc user.UserClient) PortalPost {
	return PortalPost{
		userRpc: userRpc,
	}
}

func (model *PortalPost) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model)
	db.AutoMigrate(&PortalCategoryPost{})
	db.AutoMigrate(&PostLikePost{})
	db.AutoMigrate(&PostFavoritesPost{})
}

func (model *PortalPost) PortalList(db *gorm.DB, query string, queryArgs []interface{}) ([]PortalPost, error) {

	query += " AND delete_at = ?"
	queryArgs = append(queryArgs, 0)

	var post []PortalPost
	tx := db.Where(query, queryArgs...).Order("list_order desc,id desc").Find(&post)
	for k, v := range post {
		m := More{}
		json.Unmarshal([]byte(v.More), &m)
		post[k].MoreJson = m
		post[k].PublishedTime = time.Unix(v.PublishedAt, 0).Format(data.TimeLayout)
	}
	if tx.Error != nil {
		return post, tx.Error
	}
	return post, nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 根据分类显示文章列表
 * @Date 2020/11/26 13:24:01
 * @Param
 * @return
 **/

func (model *PortalPost) ListByCategory(db *gorm.DB, current, pageSize int, query string, queryArgs []interface{}, extra map[string]string) (result data.Paginate, err error) {

	order := "p.list_order desc,p.id desc"
	if extra["hot"] == "1" {
		order = "p.post_hits desc," + order
	}

	var total int64 = 0
	conf := database.Config()
	prefix := conf.Prefix
	db.Table(prefix+"portal_post p").Distinct("p.id").
		Joins("LEFT JOIN "+prefix+"portal_category_post cp ON p.id = cp.post_id").
		Joins("LEFT JOIN "+prefix+"portal_category pc ON pc.id = cp.category_id").
		Where(query, queryArgs...).
		Order(order).
		Count(&total)

	var portalPostData []PortalPost
	tx := db.Table(prefix+"portal_post p").Select("p.*,pc.name").
		Joins("LEFT JOIN "+prefix+"portal_category_post cp ON p.id = cp.post_id").
		Joins("LEFT JOIN "+prefix+"portal_category pc ON pc.id = cp.category_id").
		Where(query, queryArgs...).Limit(pageSize).Offset((current - 1) * pageSize).
		Order(order).
		Group("p.id").Scan(&portalPostData)

	if tx.Error != nil {
		err = tx.Error
		return
	}
	for k, v := range portalPostData {
		portalPostData[k].ThumbPrevPath = util.FileUrl(v.Thumbnail)
		category := PortalCategory{}
		categoryItem, _ := category.FindPostCategory(db, "p.id = ? AND  p.delete_at = ?", []interface{}{v.Id, 0})
		portalPostData[k].Category = categoryItem

		createTime := time.Unix(v.CreateAt, 0).Format("2006-01-02 15:04:05")
		portalPostData[k].CreateTime = createTime

		updateTime := time.Unix(v.UpdateAt, 0).Format("2006-01-02 15:04:05")
		portalPostData[k].UpdateTime = updateTime

		publishTime := time.Unix(v.PublishedAt, 0).Format("2006-01-02 15:04:05")
		portalPostData[k].PublishedTime = publishTime

		m := More{}
		json.Unmarshal([]byte(v.More), &m)

		extendsObj := map[string]string{}
		extends := m.Extends
		for _, extend := range extends {
			extendsObj[extend.Key] = extend.Value
		}

		m.ExtendsObj = extendsObj

		portalPostData[k].MoreJson = m

		//data, _ := new(PortalTag).ListByPostId(v.Id)
		//
		//portalPostArr[k].Tags = data
		portalPostData[k].Tags = []PostTagResult{}
	}

	result = data.Paginate{Data: portalPostData, Current: current, PageSize: pageSize, Total: total}

	if len(portalPostData) == 0 {
		result.Data = make([]string, 0)
	}

	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询文章所在的分类
 * @Date 2020/11/27 13:24:46
 * @Param
 * @return
 **/

func (model *PortalCategory) FindPostCategory(db *gorm.DB, query string, queryArgs []interface{}) ([]PortalCategory, error) {

	var category []PortalCategory

	conf := database.Config()
	prefix := conf.Prefix

	result := db.Table(prefix+"portal_post p").Select("pc.*").
		Joins("INNER JOIN "+prefix+"portal_category_post pcp ON pcp.post_id = p.id").
		Joins("INNER JOIN "+prefix+"portal_category pc ON pc.id = pcp.category_id").
		Where(query, queryArgs...).Scan(&category)

	for k, v := range category {
		topCategory := new(PortalCategory)
		err := topCategory.GetTopCategory(db, v.Id)
		if err != nil {
			category[k].TopAlias = topCategory.Alias
			category[k].PrevPath = util.FileUrl(v.Thumbnail)
		}
	}

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return category, result.Error
		}
	}

	if len(category) == 0 {
		category = make([]PortalCategory, 0)
	}

	return category, nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查看单条文章
 * @Date 2021/12/12 18:31:44
 * @Param
 * @return
 **/

func (model *PortalPost) Show(db *gorm.DB, query string, queryArgs []interface{}) (err error) {

	tx := db.Where(query, queryArgs...).
		Order("id desc").
		First(&model)

	if tx.Error != nil {
		err = tx.Error
		return
	}

	m := More{}
	json.Unmarshal([]byte(model.More), &m)

	extendsObj := map[string]string{}
	extends := m.Extends
	for _, extend := range extends {
		extendsObj[extend.Key] = extend.Value
	}

	m.ExtendsObj = extendsObj

	model.MoreJson = m
	model.Template = m.Template

	createTime := time.Unix(model.CreateAt, 0).Format("2006-01-02 15:04:05")
	model.CreateTime = createTime

	updateTime := time.Unix(model.UpdateAt, 0).Format("2006-01-02 15:04:05")
	model.UpdateTime = updateTime

	publishTime := time.Unix(model.PublishedAt, 0).Format("2006-01-02 15:04:05")
	model.PublishedTime = publishTime

	model.ThumbPrevPath = util.FileUrl(model.Thumbnail)
	model.MoreJson.AudioPrevPath = util.FileUrl(model.MoreJson.Audio)
	model.MoreJson.VideoPrevPath = util.FileUrl(model.MoreJson.Video)
	return

}

func (model *PortalPost) Store(db *gorm.DB) (err error) {
	tx := db.Create(&model)
	if tx.Error != nil {
		return nil
	}
	return nil
}

func (model *PortalPost) Update(db *gorm.DB) (err error) {
	tx := db.Save(&model)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return nil
}
