package model

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/model"
)

type Site struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SiteId   int64              `bson:"siteId" json:"siteId"`
	Name     string             `bson:"name" json:"name"`
	Domain   string             `bson:"domain,omitempty" json:"domain"`
	Desc     string             `bson:"desc,omitempty" json:"desc"`
	Dsn      string             `bson:"dsn,omitempty" json:"dsn"`
	Status   int                `bson:"status" json:"status"`
	DeleteAt int64              `bson:"deleteAt" json:"deleteAt"`
	model.Time
}

type SiteUser struct {
	Id        int64              `bson:"_id,omitempty" json:"id"`
	TenantId  int64              `bson:"tenantId" json:"tenantId"`
	SiteId    int64              `bson:"siteId" json:"siteId"`
	Uid       int64              `bson:"uid" json:"uid"`
	Oid       primitive.ObjectID `bson:"oid" json:"oid"`
	IsOwner   int                `bson:"isOwner" json:"isOwner"`
	ListOrder float64            `bson:"listOrder" json:"listOrder"`
	Status    int                `bson:"status" json:"status"`
}

func (m *Site) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&m)
}

func (m *Site) Show(db database.MongoDB, match bson.M) (err error) {

	collection := db.Collection("siteUser")

	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "site",
				"localField":   "siteId",
				"foreignField": "siteId",
				"as":           "siteInfo",
			},
		},
		{
			"$match": match,
		},
		{
			"$limit": 1, // 限制结果集只返回一条数据
		},
		{
			"$unwind": "$siteInfo", // 平铺数组字段
		},
		{
			"$project": bson.M{
				"domain":    1,
				"desc":      1,
				"oid":       1,
				"isOwner":   1,
				"createAt":  1,
				"deleteAt":  1,
				"listOrder": 1,
				"status":    1,
				"siteId":    "$siteInfo.siteId",
				"name":      "$siteInfo.name",
			},
		},
	}

	var (
		cursor *mongo.Cursor
	)
	cursor, err = collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return
	}

	if !cursor.Next(context.TODO()) {
		err = errors.New("未查询到该站点！")
		return
	}

	var result bson.M
	err = cursor.Decode(&result)
	if err != nil {
		// 处理错误
		return
	}

	bsonBytes, err := bson.Marshal(result)
	if err != nil {
		// 处理错误
		return
	}

	err = bson.Unmarshal(bsonBytes, &m)
	if err != nil {
		// 处理错误
		return
	}

	return
}
