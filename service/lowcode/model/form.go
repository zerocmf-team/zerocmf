package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/model"
)

type SRules struct {
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Min      int    `bson:"min,omitempty" json:"min,omitempty"`
	Max      int    `bson:"max,omitempty" json:"max,omitempty"`
	Pattern  string `bson:"pattern,omitempty" json:"pattern,omitempty"`
	Message  string `bson:"message,omitempty" json:"message,omitempty"`
}

type SProps struct {
	Rules []SRules `json:"rules"`
}

type Form struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ParentId    primitive.ObjectID `bson:"parentId" json:"parentId"`
	Name        string             `bson:"name" json:"name"`
	Icon        string             `bson:"icon" json:"icon"`
	UserId      int64              `bson:"userId" json:"userId"`
	MenuType    int                `bson:",omitempty" json:"menuType"`
	HideInMenu  int                `bson:"hideInMenu,omitempty" json:"hideInMenu"`
	Description string             `bson:"description,omitempty" json:"description"`
	Schema      string             `bson:"schema,omitempty" json:"schema"`
	Rules       []SRules           `bson:"rules" json:"rules"`
	ListOrder   float64            `bson:"listOrder,omitempty" json:"listOrder"`
	Status      int                `bson:"status,omitempty" json:"status"`
	DeleteAt    int64              `bson:"deleteAt,omitempty" json:"deleteAt"`
	model.Time
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 行转树结构体
 * @Date 2021/11/30 12:52:26
 * @Param
 * @return
 **/

type iRouters struct {
	Id         primitive.ObjectID `json:"id"`
	Name       string             `json:"name"`
	Redirect   string             `json:"redirect"`
	Path       string             `json:"path"`
	Icon       string             `json:"icon"`
	HideInMenu int                `json:"hideInMenu"`
	ListOrder  float64            `json:"listOrder"`
	CreateAt   int64              `json:"createAt"`
	CreateTime string             `json:"createTime"`
	Routes     []iRouters         `json:"routes,omitempty"`
}

func RecursionMenu(menus []Form, parentId primitive.ObjectID) (routes []iRouters) {
	var routesResult = make([]iRouters, 0)
	for _, v := range menus {
		if parentId == v.ParentId {
			result := iRouters{
				Id:         v.Id,
				Name:       v.Name,
				Path:       "/admin/form/" + v.Id.Hex(),
				Icon:       v.Icon,
				HideInMenu: v.HideInMenu,
				ListOrder:  v.ListOrder,
				CreateAt:   v.CreateAt,
				CreateTime: time.Unix(v.CreateAt, 0).Format(data.TimeLayout),
			}
			children := RecursionMenu(menus, v.Id)

			if len(children) > 0 {
				result.Redirect = children[0].Path
			}

			result.Routes = children
			routesResult = append(routesResult, result)
		}
	}
	return routesResult
}

func (m *Form) Show(db database.MongoDB, filter bson.M) (err error) {
	collection := db.Collection("form")
	err = db.FindOne(collection, filter, &m)
	if err != nil {
		return
	}
	return
}
