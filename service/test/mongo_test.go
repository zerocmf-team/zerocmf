package test

import (
	"fmt"
	"testing"
	"zerocmf/common/bootstrap/database"
)

func TestMongo(t *testing.T) {

	mongoDB, err := database.NewMongoDB(
		database.Mongo{
			Username: "root",
			Password: "123456",
			Port:     27017,
		},
		"tenant",
	)
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}

	err = mongoDB.Ping()
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}

	//err = client.Ping(ctx, readpref.Primary())
	//if err != nil {
	//	fmt.Println("err", err.Error())
	//	return
	//}
	//
	//db := client.Database("lowcode")
	//collection := db.Collection("test")
	//
	//var iResult *mongo.InsertOneResult
	//
	//var testModel struct {
	//	Name string
	//}
	//
	//testModel.Name = "name"
	//
	//if iResult, err = collection.InsertOne(context.TODO(), testModel); err != nil {
	//	fmt.Print(err)
	//	return
	//}
	////_id:默认生成一个全局唯一ID
	//id := iResult.InsertedID.(primitive.ObjectID)
	//fmt.Println("自增ID", id.Hex())

}
