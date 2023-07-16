package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"zerocmf/common/bootstrap/database"
)

type Settings struct {
	Id    primitive.ObjectID `bson:"_id" json:"id"`
	Key   string             `bson:"key" json:"key"`
	Value bson.M             `bson:"value" json:"value"`
}

func (t *Settings) Migrate(db database.MongoDB) (err error) {
	collection := db.Collection("settings")
	//
	settings := Settings{}
	findErr := db.FindOne(collection, bson.M{
		"key": "portal",
	}, &settings)
	if findErr != nil && !errors.Is(findErr, mongo.ErrNoDocuments) {
		return err
	}

	if settings.Id.IsZero() {
		_, err = db.InsertOne(collection, bson.M{
			"key": "portal",
			"value": bson.M{
				"theme": "default",
			},
		})
		if err != nil {
			return err
		}
	}
	return
}

func (t *Settings) Show(db database.MongoDB, filter bson.M) (err error) {
	collection := db.Collection("settings")
	err = db.FindOne(collection, filter, &t)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return
	}
	return
}

func (t *Settings) Store(db database.MongoDB, params bson.M) (saveData bson.M, err error) {

	collection := db.Collection("settings")
	// 新增
	filter := bson.M{
		"key": t.Key,
	}
	result := bson.M{}
	err = db.FindOne(collection, filter, &result)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return
		}
	}

	var objectId primitive.ObjectID
	id := result["_id"]
	if id != nil {
		objectId = id.(primitive.ObjectID)
	}

	saveData = bson.M{
		"key":   t.Key,
		"value": params,
	}

	if objectId.IsZero() {
		//var one *mongo.InsertOneResult
		_, err = db.InsertOne(collection, &saveData)
		if err != nil {
			return
		}
	} else {
		_, err = db.UpdateOne(collection, filter, bson.M{
			"$set": saveData,
		})
		if err != nil {
			return
		}
	}
	return
}
