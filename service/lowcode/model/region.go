package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"zerocmf/common/bootstrap/database"
)

type Region struct {
	Code     int      `bson:"code" json:"code"`
	Name     string   `bson:"name" json:"name"`
	Type     int      `bson:"type" json:"type"`
	ParentId int      `bson:"parentId" json:"parentId"`
	Children []Region `bson:"children,omitempty" json:"children,omitempty"`
}

func (m *Region) List(db database.MongoDB) (result []Region, err error) {
	collection := db.Collection("region")
	var cursor *mongo.Cursor
	cursor, err = collection.Find(context.Background(), bson.M{})
	if err != nil {
		return
	}
	defer cursor.Close(context.Background())
	// 遍历结果
	for cursor.Next(context.Background()) {
		var bsonM bson.M
		if err = cursor.Decode(&bsonM); err != nil {
			return
		}
		var region Region
		// 转换为字节数组
		var bytes []byte
		bytes, err = bson.Marshal(bsonM)
		if err != nil {
			return
		}

		// 转换为结构体
		err = bson.Unmarshal(bytes, &region)
		if err != nil {
			return
		}
		result = append(result, region)
	}

	result = recursionRegion(result, 0)

	if err = cursor.Err(); err != nil {
		return
	}
	return
}

func recursionRegion(region []Region, parentId int) (result []Region) {
	for _, v := range region {
		if v.ParentId == parentId {
			children := recursionRegion(region, v.Code)
			v.Children = children
			result = append(result, v)
		}
	}
	return
}
