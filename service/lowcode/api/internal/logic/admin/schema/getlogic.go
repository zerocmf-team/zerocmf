package schema

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"time"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"
	"zerocmf/service/lowcode/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type iRoute struct {
	Key        string   `json:"key"`
	Name       string   `json:"name"`
	Plugin     string   `json:"plugin"` // 微服务应用
	MenuType   int      `json:"menuType"`
	HideInMenu int      `json:"hideInMenu"`
	Path       string   `json:"path"`
	Redirect   string   `json:"redirect"`
	Component  string   `json:"component"`
	FilePath   string   `json:"filePath"`
	ListOrder  float64  `json:"listOrder"`
	Routes     []iRoute `json:"routes"`
}

func (l *GetLogic) Get(req *types.InitReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	err = new(model.Lowcode).Migrate(db)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Get 无法获取当前目录：", err)
		return
	}

	routeData, rErr := os.ReadFile(dir + "/template/menu.json")
	if rErr != nil {
		fmt.Println("无法获取当前目录：", rErr)
		return
	}

	var routes []iRoute

	json.Unmarshal(routeData, &routes)

	err = importLogic{
		db: db,
	}.recursionImport(routes)
	if err != nil {
		return
	}

	resp.Success("操作成功！", nil)
	return
}

type importLogic struct {
	db       database.MongoDB
	parentId primitive.ObjectID
	key      string
	plugin   string

	level int
}

func (i importLogic) recursionImport(routes []iRoute) (err error) {

	var dir string
	dir, err = os.Getwd()
	if err != nil {
		fmt.Println("无法获取当前目录：", err)
		return
	}

	db := i.db
	_key := i.key
	_parentId := i.parentId
	_plugin := i.plugin

	level := i.level

	for _, v := range routes {

		key := _key
		parentId := _parentId
		plugin := v.Plugin

		if key == "" {
			key = v.Key
		} else {
			key = key + "." + v.Key
		}

		if level == 0 {
			i.plugin = plugin
		} else if plugin == "" {
			plugin = _plugin
		}

		var file []byte
		if v.FilePath != "" {
			file, err = os.ReadFile(dir + v.FilePath)
			if err != nil {
				fmt.Println("无法获取当前目录：", err)
				return
			}
		}

		menuType := v.MenuType
		var formId primitive.ObjectID

		// 表单
		if menuType == 0 {
			schema := string(file)
			// 初始化默认脚本到form表中
			form := model.Form{}
			filter := bson.M{
				"key": key,
			}
			collection := db.Collection("form")
			err = db.FindOne(collection, filter, &form)
			if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
				return
			}

			form.Key = key
			form.Schema = schema
			form.Name = v.Name
			form.Status = 1

			// 未查询到，则导入

			if form.Id.IsZero() {
				form.CreateAt = time.Now().Unix()
				form.UpdateAt = time.Now().Unix()
				var one *mongo.InsertOneResult
				one, err = db.InsertOne(collection, &form)
				if err != nil {
					return
				}
				formId = one.InsertedID.(primitive.ObjectID)
			} else {
				form.UpdateAt = time.Now().Unix()
				db.UpdateOne(collection, filter, bson.M{
					"$set": form,
				})
				formId = form.Id
			}

		}

		// 创建路由
		collection := db.Collection("adminMenu")
		menu := model.AdminMenu{}
		filter := bson.M{
			"key": key,
		}

		err = db.FindOne(collection, filter, &menu)
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			return
		}

		menu.MenuType = v.MenuType
		menu.Key = key
		menu.Plugin = plugin
		menu.HideInMenu = v.HideInMenu
		menu.Name = v.Name
		menu.Path = v.Path
		menu.Redirect = v.Redirect
		menu.Component = v.Component
		menu.ParentId = parentId
		menu.FormId = &formId
		menu.ListOrder = v.ListOrder
		menu.CreateAt = time.Now().Unix()
		menu.UpdateAt = time.Now().Unix()

		var menuId primitive.ObjectID
		if menu.Id.IsZero() {
			var one *mongo.InsertOneResult
			one, err = db.InsertOne(collection, menu)
			if err != nil {
				return err
			}
			menuId = one.InsertedID.(primitive.ObjectID)
		} else {

			//fmt.Println("menu", menu.Name, menu.Plugin)

			db.UpdateOne(collection, filter, bson.M{
				"$set": menu,
			})
			menuId = menu.Id
		}

		if v.Routes != nil {
			i.key = key
			i.parentId = menuId
			i.level = level + 1
			err = i.recursionImport(v.Routes)
			if err != nil {
				return err
			}
		}
	}
	return
}
