/**
Desc: 后台菜单
Author: daifuyang
Contact: github.com/daifuyang
Date: Date: 2023-07-06 19:46:18
*/

package model

import (
	"context"
	"errors"
	"fmt"
	"time"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
Desc: 菜单结构体
Author: daifuyang
Contact: github.com/daifuyang
Date: Date: 2023-07-06 19:46:38
*/

type AdminMenu struct {
	Id         primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	ParentId   primitive.ObjectID  `bson:"parentId" json:"parentId"`
	FormId     *primitive.ObjectID `bson:"formId" json:"formId,omitempty"`
	Plugin     string              `bson:"plugin" json:"plugin,omitempty"` //菜单唯一标识，恒定，用来确定子页面的路由 portal=> 门户 shop => 商城
	MenuType   int                 `bson:"menuType" json:"menuType"`
	Key        string              `bson:"key" json:"key"` //菜单唯一标识
	Name       string              `bson:"name" json:"name"`
	Path       string              `bson:"path" json:"path"`
	Redirect   string              `bson:"redirect" json:"redirect"`
	Component  string              `bson:"component" json:"component,omitempty"`
	Icon       string              `bson:"icon" json:"icon"`
	HideInMenu int                 `bson:"hideInMenu" json:"hideInMenu"`
	ListOrder  float64             `bson:"listOrder" json:"listOrder"`
	CreateAt   int64               `bson:"createAt" json:"createAt"`
	UpdateAt   int64               `bson:"updateAt" json:"updateAt"`
}

type adminMenu struct {
	Id         primitive.ObjectID  `json:"id"`
	ParentId   *primitive.ObjectID `json:"parentId,omitempty"`
	FormId     *primitive.ObjectID `json:"formId,omitempty"`
	MenuType   int                 `json:"menuType"`
	Plugin     string              `json:"plugin"`
	Name       string              `json:"name"`
	Path       string              `json:"path"`
	Redirect   string              `json:"redirect,omitempty"`
	Component  string              `json:"component,omitempty"`
	Icon       string              `json:"icon"`
	HideInMenu int                 `json:"hideInMenu"`
	ListOrder  float64             `json:"listOrder"`
	Routes     []adminMenu         `json:"routes"`
	CreateAt   int64               `json:"createAt"`
	CreateTime string              `json:"createTime"`
	UpdateAt   int64               `bson:"updateAt" json:"updateAt"`
	UpdateTime string              `json:"UpdateTime"`
}

/**
Desc: 递归获取树型菜单
Author: daifuyang
Contact: github.com/daifuyang
Date: Date: 2023-07-06 20:09:37
*/

func recursionMenu(menus []AdminMenu, parentId primitive.ObjectID) (routes []adminMenu) {
	var adminMenus = make([]adminMenu, 0)
	for _, v := range menus {
		if parentId == v.ParentId {
			menu := adminMenu{
				Id:         v.Id,
				MenuType:   v.MenuType,
				Name:       v.Name,
				Plugin:     v.Plugin,
				Path:       v.Path,
				Redirect:   v.Redirect,
				Component:  v.Component,
				Icon:       v.Icon,
				HideInMenu: v.HideInMenu,
				ListOrder:  v.ListOrder,
				CreateAt:   v.CreateAt,
				CreateTime: time.Unix(v.CreateAt, 0).Format(data.TimeLayout),
				UpdateAt:   v.UpdateAt,
				UpdateTime: time.Unix(v.UpdateAt, 0).Format(data.TimeLayout),
			}

			if !v.ParentId.IsZero() {
				menu.ParentId = &v.ParentId
			}

			if !v.FormId.IsZero() {
				menu.FormId = v.FormId
			}

			children := recursionMenu(menus, v.Id)
			//if len(children) > 0 {
			//	menu.Redirect = children[0].Path
			//}

			menu.Routes = children
			adminMenus = append(adminMenus, menu)
		}
	}
	return adminMenus
}

/**
Desc: 获取全部菜单，无需分页
Author: daifuyang
Contact: github.com/daifuyang
Date: Date: 2023-07-06 19:49:49
*/

func (rest *AdminMenu) GetTrees(db database.MongoDB, ctx context.Context, filter bson.M) (menus []adminMenu, err error) {
	collection := db.Collection("adminMenu")
	opts := options.Find().SetSort(bson.D{{"listOrder", -1}})
	cursor, fErr := collection.Find(ctx, filter, opts)
	if fErr != nil {
		err = fErr
		return
	}
	result := make([]AdminMenu, 0)
	for cursor.Next(context.Background()) {
		menu := AdminMenu{}
		err = cursor.Decode(&menu)
		if err != nil {
			return
		}
		result = append(result, menu)
	}

	menus = recursionMenu(result, primitive.ObjectID{})
	return
}

func (rest *AdminMenu) Show(db database.MongoDB, filter bson.M) (err error) {
	collection := db.Collection("adminMenu")
	err = db.FindOne(collection, &filter, &rest)
	if err != nil {
		return
	}
	return
}

func (rest *AdminMenu) Save(db database.MongoDB, filter bson.M) (err error) {
	collection := db.Collection("adminMenu")
	menu := new(AdminMenu)
	showErr := menu.Show(db, filter)
	if showErr != nil && !errors.Is(showErr, mongo.ErrNoDocuments) {
		err = showErr
		return
	}
	// 新增
	if rest.Id.IsZero() {
		rest.CreateAt = time.Now().Unix()
		rest.UpdateAt = time.Now().Unix()
		one, oneErr := db.InsertOne(collection, &rest)
		fmt.Println(one, oneErr)
		if oneErr != nil {
			err = oneErr
			return
		}
		rest.Id = one.InsertedID.(primitive.ObjectID)
	} else {
		rest.UpdateAt = time.Now().Unix()
		_, oneErr := db.UpdateOne(collection, filter, &rest)
		if oneErr != nil {
			err = oneErr
			return
		}
		rest.Id = menu.Id
	}
	return
}
