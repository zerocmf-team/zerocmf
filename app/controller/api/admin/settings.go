package admin

import (
	"encoding/json"
	"fmt"
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	"github.com/gincmf/cmf/controller"
	"strings"
)

type SettingsController struct {
	controller string
	rc         controller.RestController
}

/*
 * @restApi(
 *     'name'  	=> '获取系统设置',
 *	   'desc'  	=> '获取系统设置'
 *     'url'   	=> 'api/admin/settings',
 *	   'param' 	=>  '',
 *	   'method'	=> "get",
 *	   'status' => 1
 * )
 */
func (rest *SettingsController) Get(c *gin.Context) {
	option := &model.Option{}
	siteResult := cmf.Db.First(option, "option_name = ?", "site_info") // 查询
	if siteResult.RowsAffected > 0 {
		rest.rc.Success(c, "获取成功", option)
	} else {
		rest.rc.Error(c, "获取失败", nil)
	}
}

func (rest *SettingsController) Show(c *gin.Context) {
	var rewrite struct {
		id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	rest.rc.Success(c, "操作成功show", nil)
}

func (rest *SettingsController) Edit(c *gin.Context) {
	rest.rc.Success(c, "操作成功Edit", nil)
}

/*
 * @restApi(
 *     'name'  		=> '更新系统设置',
 *	   'desc'  		=> '更新系统设置'
 *     'url'   		=> 'api/admin/settings',
 *	   'param' 		=>  '',
 *	   'method'		=> 'post',
 *	   'list_order' => '10000',
 *	   'status'		=> 1
 * )
 */
func (rest *SettingsController) Store(c *gin.Context) {
	//siteInfo := &model.SiteInfo{
	//	SiteName:           c.PostForm("site_name"),
	//	AdminPassword:      c.PostForm("admin_password"),
	//	SiteSeoTitle:       c.PostForm("site_seo_title"),
	//	SiteSeoKeywords:    c.PostForm("site_seo_keywords"),
	//	SiteSeoDescription: c.PostForm("site_seo_description"),
	//	SiteIcp:            c.PostForm("site_icp"),
	//	SiteGwa:            c.PostForm("site_gwa"),
	//	SiteAdminEmail:     c.PostForm("site_admin_email"),
	//	SiteAnalytics:      c.PostForm("site_analytics"),
	//	OpenRegistration:   c.PostForm("open_registration"),
	//}


	c.Request.ParseForm()
	var params = make(map[string]interface{},len(c.Request.Form))
	for k,v := range c.Request.Form{
		if len(v) > 0 {
			params[k] = strings.Join(v,"")
		}
	}
	siteInfoValue, _ := json.Marshal(params)
	fmt.Println("siteInfoValue", string(siteInfoValue))
	cmf.Db.Model(&model.Option{}).Where("option_name = ?", "site_info").Update("option_value", string(siteInfoValue))
	rest.rc.Success(c, "修改成功", params)
}

func (rest *SettingsController) Delete(c *gin.Context) {
	rest.rc.Success(c, "操作成功Delete", nil)
}
