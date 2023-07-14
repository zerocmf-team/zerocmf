package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"zerocmf/service/user/model"
)

type FormData struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FormId     primitive.ObjectID `bson:"formId" json:"formId"`
	Schema     []ColumnsProps     `bson:"schema" json:"schema"`
	User       model.User         `bson:"user" json:"user"`
	CreateAt   int64              `bson:"createAt"  json:"createAt"`
	UpdateAt   int64              `bson:"updateAt"  json:"updateAt"`
	CreateTime string             `bson:"-" json:"createTime"`
	UpdateTime string             `bson:"-" json:"updateTime"`
	/*	Data      map[string]interface{}*/
}
