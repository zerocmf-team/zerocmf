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

type ThemePage struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Theme       string             `bson:"theme" json:"theme"`
	IsPublic    int                `bson:"isPublic" json:"isPublic"`
	Key         string             `bson:"key" json:"key"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Type        string             `bson:"type" json:"type"`
	Schema      string             `bson:"schema" json:"schema"`
	ListOrder   float64            `bson:"listOrder" json:"listOrder"`
	UserId      int64              `bson:"userId" json:"userId"`
	CreateAt    int64              `bson:"createAt" json:"createAt"`
	UpdateAt    int64              `bson:"updateAt" json:"updateAt"`
	CreateTime  string             `bson:"-" json:"createTime"`
	UpdateTime  string             `bson:"-" json:"updateTime"`
	Status      int                `bson:"status" json:"status"`
	DeleteAt    int64              `bson:"deleteAt" json:"deleteAt"`
	/*SeoTitle       string             `bson:"seoTitle" json:"title"`
	SeoKeywords    string             `bson:"seoKeywords" json:"seoKeywords"`
	SeoDescription string             `bson:"seoDescription" json:"seoDescription"`*/
}

func (t *ThemePage) Migrate(db database.MongoDB) (err error) {
	collection := db.Collection("themePage")
	themePage := ThemePage{}
	findErr := db.FindOne(collection, bson.M{"theme": "default", "key": "home"}, &themePage)
	if findErr != nil && !errors.Is(findErr, mongo.ErrNoDocuments) {
		return findErr
	}

	if themePage.Id.IsZero() {
		var one *mongo.InsertOneResult
		one, err = db.InsertOne(collection, ThemePage{
			Theme:     "default",
			Key:       "home",
			Name:      "主页",
			Type:      "page",
			UserId:    1,
			CreateAt:  time.Now().Unix(),
			UpdateAt:  time.Now().Unix(),
			Status:    1,
			ListOrder: 10000,
		})
		if err != nil {
			return err
		}
		objectId := one.InsertedID.(primitive.ObjectID)
		settings := Settings{Key: "portal"}

		err = settings.Show(db, bson.M{"key": "portal"})
		if err != nil {
			return err
		}
		params := settings.Value
		params["mainPage"] = objectId
		_, err = settings.Store(db, params)
		if err != nil {
			return err
		}
	}

	themePage = ThemePage{}
	findErr = db.FindOne(collection, bson.M{"theme": "default", "key": "list"}, &themePage)
	if findErr != nil && !errors.Is(findErr, mongo.ErrNoDocuments) {
		return findErr
	}

	if themePage.Id.IsZero() {
		_, err = db.InsertOne(collection, ThemePage{
			Theme:     "default",
			Key:       "list",
			Name:      "默认列表",
			Type:      "list",
			UserId:    1,
			CreateAt:  time.Now().Unix(),
			UpdateAt:  time.Now().Unix(),
			Status:    1,
			ListOrder: 10000,
		})
		if err != nil {
			return err
		}
	}

	themePage = ThemePage{}
	findErr = db.FindOne(collection, bson.M{"theme": "default", "key": "article"}, &themePage)
	if findErr != nil && !errors.Is(findErr, mongo.ErrNoDocuments) {
		return findErr
	}

	if themePage.Id.IsZero() {
		_, err = db.InsertOne(collection, ThemePage{
			Theme:     "default",
			Key:       "article",
			Name:      "默认文章",
			Type:      "article",
			UserId:    1,
			CreateAt:  time.Now().Unix(),
			UpdateAt:  time.Now().Unix(),
			Status:    1,
			ListOrder: 10000,
		})
		if err != nil {
			return err
		}
	}

	themePage = ThemePage{}
	findErr = db.FindOne(collection, bson.M{"theme": "default", "key": "page"}, &themePage)
	if findErr != nil && !errors.Is(findErr, mongo.ErrNoDocuments) {
		return findErr
	}

	if themePage.Id.IsZero() {
		_, err = db.InsertOne(collection, ThemePage{
			Theme:     "default",
			Key:       "page",
			Name:      "默认页面",
			Type:      "page",
			UserId:    1,
			CreateAt:  time.Now().Unix(),
			UpdateAt:  time.Now().Unix(),
			Status:    1,
			ListOrder: 10000,
		})
		if err != nil {
			return err
		}
	}

	return
}

func (t *ThemePage) List(db database.MongoDB, current int, pageSize int, filter bson.M) (result interface{}, err error) {

	collection := db.Collection("themePage")
	//	如果pageSize为0，则无需分页
	var (
		cursor  *mongo.Cursor
		findErr error
	)

	var pages = make([]ThemePage, 0)

	if pageSize == 0 {
		cursor, findErr = collection.Find(context.TODO(), filter)
		if err != nil && !errors.Is(findErr, mongo.ErrNoDocuments) {
			return
		}

		defer cursor.Close(context.TODO())

		for cursor.Next(context.TODO()) {
			var page ThemePage
			if err = cursor.Decode(&page); err != nil {
				// 处理解码错误
				return
			}
			createTime := time.Unix(page.CreateAt, 0).Format("2006-01-02 15:04:05")
			page.CreateTime = createTime

			updateTime := time.Unix(page.UpdateAt, 0).Format("2006-01-02 15:04:05")
			page.CreateTime = updateTime
			pages = append(pages, page)
		}

		result = pages

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
				Data       []ThemePage `bson:"data"`
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

func (t *ThemePage) Show(db database.MongoDB, filter bson.M) (err error) {
	collection := db.Collection("themePage")
	err = db.FindOne(collection, filter, &t)
	if err != nil {
		return
	}
	return
}
