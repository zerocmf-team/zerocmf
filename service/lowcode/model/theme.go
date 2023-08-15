package model

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/database"
)

type Theme struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type        int                `json:"type"` // 0=>web;1=>wei
	Default     int                `json:"default"`
	Key         string             `json:"key"`
	Name        string             `json:"name"`
	Version     string             `json:"version"`
	Thumbnail   string             `json:"thumbnail"`
	Description string             `json:"description"`
	UserId      int                `json:"user_id"`
	CreateAt    int64              `json:"createAt"`
	UpdateAt    int64              `json:"updateAt"`
	CreateTime  string             `bson:"-" json:"createTime"`
	UpdateTime  string             `bson:"-" json:"updateTime"`
	ListOrder   float64            `json:"listOrder"`
	DeleteAt    int64              `json:"deleteAt"`
}

func (t Theme) Migrate(db database.MongoDB) (err error) {
	collection := db.Collection("theme")
	theme := Theme{}
	findErr := db.FindOne(collection, bson.M{"type": 0, "default": 1}, &theme)
	if findErr != nil && !errors.Is(findErr, mongo.ErrNoDocuments) {
		return findErr
	}

	if theme.Id.IsZero() {
		_, err = db.InsertOne(collection, Theme{
			Type:     0,
			Default:  1,
			Key:      "default",
			Name:     "默认主题",
			Version:  "0.0.1",
			UserId:   1,
			CreateAt: time.Now().Unix(),
			UpdateAt: time.Now().Unix(),
		})
		if err != nil {
			return err
		}

		setting := Settings{Key: "portal"}
		setting.Store(db, bson.M{
			"theme": "default",
		})

	}

	theme = Theme{}
	findErr = db.FindOne(collection, bson.M{"type": 1, "default": 1}, &theme)
	if findErr != nil && !errors.Is(findErr, mongo.ErrNoDocuments) {
		return findErr
	}

	if theme.Id.IsZero() {
		_, err = db.InsertOne(collection, Theme{
			Type:     1,
			Default:  1,
			Key:      "default",
			Name:     "默认主题",
			Version:  "0.0.1",
			UserId:   1,
			CreateAt: time.Now().Unix(),
			UpdateAt: time.Now().Unix(),
		})
		if err != nil {
			return err
		}
		setting := Settings{Key: "wei"}
		setting.Store(db, bson.M{
			"theme": "default",
		})
	}
	return
}

func (t Theme) List(db database.MongoDB, current int, pageSize int, filter bson.M) (result interface{}, err error) {
	collection := db.Collection("theme")
	//	如果pageSize为0，则无需分页
	var (
		cursor  *mongo.Cursor
		findErr error
	)
	var themes = make([]Theme, 0)
	if pageSize == 0 {
		cursor, findErr = collection.Find(context.TODO(), filter)
		if findErr != nil && !errors.Is(findErr, mongo.ErrNoDocuments) {
			err = findErr
			return
		}
		defer cursor.Close(context.TODO())

		for cursor.Next(context.TODO()) {
			var theme Theme
			if err = cursor.Decode(&result); err != nil {
				// 处理解码错误
				return
			}
			createTime := time.Unix(theme.CreateAt, 0).Format("2006-01-02 15:04:05")
			theme.CreateTime = createTime

			updateTime := time.Unix(theme.UpdateAt, 0).Format("2006-01-02 15:04:05")
			theme.CreateTime = updateTime
			themes = append(themes, theme)
		}

		result = themes
	} else {

		paginate := data.Paginate{}

		pipeline := []bson.M{
			// 匹配过滤条件
			{
				"$match": filter,
			},
			// 分页和总记录数计算
			{
				"$facet": bson.M{
					"data": []bson.M{
						{"$skip": (current - 1) * pageSize},
						{"$limit": pageSize},
					},
					"totalCount": []bson.M{
						{"$count": "count"},
					},
				},
			},
		}

		// 执行聚合查询
		cursor, err = collection.Aggregate(context.TODO(), pipeline)
		if err != nil {
			// 处理错误
			return
		}
		defer cursor.Close(context.TODO())

		if cursor.Next(context.TODO()) {
			var aggregationResult struct {
				Data       []Theme `bson:"data"`
				TotalCount []struct {
					Count int `bson:"count"`
				} `bson:"totalCount"`
			}

			if err = cursor.Decode(&aggregationResult); err != nil {
				// 处理解码错误
				return
			}

			for k, v := range aggregationResult.Data {
				createTime := time.Unix(v.CreateAt, 0).Format("2006-01-02 15:04:05")
				aggregationResult.Data[k].CreateTime = createTime

				updateTime := time.Unix(v.UpdateAt, 0).Format("2006-01-02 15:04:05")
				aggregationResult.Data[k].UpdateTime = updateTime
			}

			paginate.Data = aggregationResult.Data
			paginate.Current = current
			if len(aggregationResult.TotalCount) > 0 {
				paginate.Total = int64(aggregationResult.TotalCount[0].Count)
			}

		}

		result = paginate

	}
	return
}
