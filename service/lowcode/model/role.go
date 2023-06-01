package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/model"
)

type Role struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Root      int                `bson:"root" json:"root,omitempty"`
	ParentId  primitive.ObjectID `bson:"parentId"  json:"parentId"`
	Name      string             `bson:"name" json:"name"`
	Remark    string             `bson:"remark" json:"remark"`
	ListOrder float64            `bson:"listOrder" json:"listOrder"`
	Status    int                `bson:"status" json:"status"`
	DeletedAt int                `bson:"deletedAt" json:"deletedAt"`
	model.Time
}

type RoleUser struct {
	Id     primitive.ObjectID `bson:"_id.o,omitempty" json:"id,omitempty"`
	RoleId primitive.ObjectID `bson:"roleId" json:"roleId"`
	UserId primitive.ObjectID `bson:"userId" json:"userId"`
}

func (_ *Role) AutoMigrate(db database.MongoDB) {
	documents := []interface{}{
		Role{
			Name:      "超级管理员",
			Remark:    "拥有网站最高管理员权限！",
			ListOrder: 0,
			Status:    1,
			Time: model.Time{
				CreateAt: time.Now().Unix(),
				UpdateAt: time.Now().Unix(),
			},
		},
		Role{
			Name:      "超级管理员",
			Remark:    "拥有网站最高管理员权限！",
			ListOrder: 0,
			Status:    1,
			Time: model.Time{
				CreateAt: time.Now().Unix(),
				UpdateAt: time.Now().Unix(),
			},
		},
		Role{
			Name:      "普通管理员",
			Remark:    "权限由最高管理员分配！",
			ListOrder: 0,
			Status:    1,
			Time: model.Time{
				CreateAt: time.Now().Unix(),
				UpdateAt: time.Now().Unix(),
			},
		},
	}
	db.Collection("role").InsertMany(context.TODO(), documents)
}
