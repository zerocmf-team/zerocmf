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

type FieldData struct {
	Text  string      `json:"text,omitempty" bson:"text,omitempty"`
	Value interface{} `json:"value,omitempty" bson:"value,omitempty"`
}

type Options struct {
	Label string `json:"label" bson:"label,omitempty"`
	Value string `json:"value" bson:"value,omitempty"`
}

type ColumnsProps struct {
	FieldId       string     `json:"fieldId" bson:"fieldId"`
	Label         string     `json:"label" bson:"label"`
	ComponentName string     `json:"componentName" bson:"componentName"`
	Unique        bool       `json:"unique"`
	Rules         []SRules   `json:"rules,omitempty" bson:"rules,omitempty"`
	FieldData     *FieldData `json:"fieldData,omitempty" bson:"fieldData,omitempty"`
	Options       []Options  `json:"options,omitempty" bson:"options,omitempty"`
}

type Form struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ParentId primitive.ObjectID `bson:"parentId" json:"parentId"`
	Key      string             `bson:"key" json:"key"`
	Name     string             `bson:"name" json:"name"`
	/*Path        string             `bson:"path" json:"path"`
	Icon        string             `bson:"icon" json:"icon"`*/
	UserId int64 `bson:"userId" json:"userId"`
	/*MenuType    int            `bson:",omitempty" json:"menuType"`
	HideInMenu  int            `bson:"hideInMenu" json:"hideInMenu"`*/
	Description string         `bson:"description" json:"description"`
	Schema      string         `bson:"schema,omitempty" json:"schema"`
	Columns     []ColumnsProps `bson:"columns" json:"columns"`
	ListOrder   float64        `bson:"listOrder" json:"listOrder"`
	Status      int            `bson:"status" json:"status"`
	DeleteAt    int64          `bson:"deleteAt" json:"deleteAt"`
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
	ParentId   string             `bson:"parentId" json:"parentId"`
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
				Id:   v.Id,
				Name: v.Name,
				Path: "/admin/form/" + v.Id.Hex(),
				//Icon:       v.Icon,
				//HideInMenu: v.HideInMenu,
				ListOrder:  v.ListOrder,
				CreateAt:   v.CreateAt,
				CreateTime: time.Unix(v.CreateAt, 0).Format(data.TimeLayout),
			}

			if v.ParentId.IsZero() == false {
				result.ParentId = v.ParentId.Hex()
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
