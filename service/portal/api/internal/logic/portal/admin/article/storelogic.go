package article

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
	"zerocmf/common/bootstrap/data"
	comModel "zerocmf/common/bootstrap/model"
	"zerocmf/service/portal/model"
	"zerocmf/service/user/rpc/types/user"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreLogic {
	return &StoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreLogic) Store(req *types.ArticleSaveReq) (resp types.Response) {
	c := l.svcCtx
	resp = save(c, req)
	return
}

func save(c *svc.ServiceContext, req *types.ArticleSaveReq) (resp types.Response) {

	db := c.Db
	userRpc := c.UserRpc
	id := req.Id

	postType := 1
	if req.PostType == 2 {
		postType = 2
	}

	if postType == 1 && len(req.CategoryIds) == 0 {
		resp.Error("分类不能为空！", nil)
		return
	}

	if req.PostTitle == "" {
		resp.Error("标题不能为空！", nil)
		return
	}

	currentTime := time.Now().Unix()

	var more model.More
	copier.Copy(&more, &req)

	if req.Template != "" {
		more.Template = req.Template
	}

	//re, _ := regexp.Compile("[^a-zA-Z0-9]+")
	//slug := re.ReplaceAllString(req.PostTitle, "-")
	//more.Slug = strings.ToLower(slug)

	more.Alias = req.Alias

	moreJson, err := json.Marshal(more)

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	userId, _ := c.Get("userId")

	userIdInt, _ := strconv.Atoi(userId.(string))

	portal := model.NewPost(userRpc)

	if id > 0 {
		query := "id = ?"
		queryArgs := []interface{}{id}
		err := portal.Show(db, query, queryArgs)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}
	}

	copier.Copy(&portal, &req)
	publishAt, _ := time.ParseInLocation(data.TimeLayout, req.PublishTime, time.Local)
	listOrder := req.ListOrder

	if listOrder == 0 {
		listOrder = 10000
	}

	portal.PostType = postType
	portal.UserId = userIdInt

	tenant, exist := db.Get("tenantId")
	tenantId := ""
	if exist {
		tenantId = tenant.(string)
	}

	userReply, err := userRpc.Get(context.Background(), &user.UserRequest{
		UserId:   int64(userIdInt),
		TenantId: tenantId,
	})

	if err != nil {
		return
	}

	portal.UserLogin = userReply.UserLogin
	portal.PostKeywords = strings.Join(req.PostKeywords, ",")
	portal.More = string(moreJson)
	if id == 0 {
		portal.CreateAt = currentTime
	}
	portal.UpdateAt = currentTime
	portal.PublishedAt = publishAt.Unix()
	portal.ListOrder = listOrder

	var (
		data struct {
			model.PortalPost
			Category []model.PortalCategoryPost `json:"category"`
		}
		postData model.PortalPost
	)

	if id == 0 {
		err = portal.Store(db)
	} else {
		portal.Id = id
		err = portal.Update(db)
	}

	if err != nil {
		resp.Error("数据库系统错误", err.Error())
		return
	}

	postData = portal

	if req.Alias != "" {
		fullUrl := "page/" + strconv.Itoa(postData.Id)
		url := req.Alias

		route := comModel.Route{
			Type:    2,
			FullUrl: fullUrl,
			Url:     url,
		}

		err = route.Set(db)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}
	}

	var pcpPost = make([]model.PortalCategoryPost, 0)

	pcp := model.PortalCategoryPost{}

	category := model.PortalCategory{}
	existsCategory, err := category.List(db)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Error("请先添加分类！", nil)
			return
		}
		resp.Error(err.Error(), nil)
		return
	}

	for _, v := range req.CategoryIds {

		cidInt, _ := strconv.Atoi(v)

		if !inArray(cidInt, existsCategory) {
			resp.Error("分类参数非法！", nil)
			return
		}

		pcp.PostId = postData.Id
		pcp.CategoryId = cidInt
		pcpPost = append(pcpPost, pcp)
	}

	pcpData, pcpErr := pcp.Store(db, pcpPost)
	if pcpErr != nil {
		resp.Error(pcpErr.Error(), nil)
		return
	}

	data.Category = pcpData

	var tag []int
	var portalTag model.PortalTag
	for _, v := range req.PostKeywords {
		if strings.TrimSpace(v) != "" {
			// 查询当前tag是否存在
			portalTag, err = new(model.PortalTag).Show(db, "name = ?", []interface{}{v})
			if err != nil {
				resp.Error(err.Error(), nil)
				return
			}
			portalTag.Name = v
			_, err = portalTag.FirstOrSave(db)
			if err != nil {
				resp.Error(err.Error(), err)
			}
			tag = append(tag, portalTag.Id)
		}
	}
	tagPost := model.PortalTagPost{
		PostId: postData.Id,
	}
	err = tagPost.FirstOrSave(db, tag)
	if err != nil {
		resp.Error("数据库系统错误", err.Error())
		return
	}
	data.PortalPost = postData
	message := "添加成功！"
	if id > 0 {
		message = "更新成功！"
	}
	resp.Success(message, data)
	return
}

func inArray(cid int, pc []model.PortalCategory) bool {
	for _, v := range pc {
		if v.Id == cid {
			return true
		}
	}
	return false
}
