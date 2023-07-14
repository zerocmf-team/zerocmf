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

func (t Settings) Migrate(db database.MongoDB) (err error) {
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
			"key":   "portal",
			"value": bson.M{},
		})
		if err != nil {
			return err
		}
	}
	return
}
