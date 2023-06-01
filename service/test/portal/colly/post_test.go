package colly

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"gorm.io/gorm"
	"strings"
	"testing"
	"time"
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/portal/model"
)

func TestPost(t *testing.T) {

	db := database.NewGormDb(database.Database{
		Type:     "mysql",
		Host:     "localhost",
		Database: "zerocmf_portal",
		Username: "root",
		Password: "123456",
		Port:     3306,
		Charset:  "utf8mb4",
		Prefix:   "cmf_",
		AuthCode: "KFHlk2ubIlMr5ltqaD",
	})

	c := colly.NewCollector(
		colly.DetectCharset(),
	)

	c.OnHTML(".news-list-box .news-list-cnt h2 a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		e.Request.Visit(link)
	})

	c.OnHTML(".news-main", func(e *colly.HTMLElement) {
		if e.Request.URL.Path != "/" {
			title := e.ChildText(".news-tit h1 span")
			fmt.Println("title", title)

			e.DOM.Find(".news-flex-box .news-cnt div").Remove()

			postExcerpt := e.DOM.Find(".news-flex-box .news-cnt p:nth-child(1)").Text()

			postExcerpt = strings.TrimSpace(postExcerpt)

			content, _ := e.DOM.Find(".news-flex-box .news-cnt").Html()
			fmt.Println("content", content)

			post := model.PortalPost{}

			err := post.Show(db, "post_title = ?", []interface{}{title})

			newPost := model.PortalPost{
				PostTitle:   title,
				PostContent: content,
				PostExcerpt: postExcerpt,
				PostType:    1,
				UserId:      1,
				UserLogin:   "admin",
				PublishedAt: time.Now().Unix(),
				UpdateAt:    time.Now().Unix(),
				ListOrder:   10000,
				PostStatus:  1,
				More:        "{}",
			}

			if err != nil {
				if err == gorm.ErrRecordNotFound {

				}
			} else {
				newPost.Id = post.Id
				err = newPost.Update(db)
			}
			err = newPost.Store(db)

			categoryId := 1

			pcp := &model.PortalCategoryPost{}
			tx := db.Where("post_id = ? AND category_id = ?", newPost.Id, categoryId).First(&pcp)
			if util.IsDbErr(tx) != nil {
				fmt.Println("db err", tx.Error)
				return
			}

			cPost := model.PortalCategoryPost{
				PostId:     newPost.Id,
				CategoryId: categoryId,
				ListOrder:  10000,
				Status:     1,
			}

			cPost.Id = pcp.Id
			db.Save(&cPost)
		}
	})

	c.Visit("http://sh.yuloo.com/zaizhi/wenti/")

}
