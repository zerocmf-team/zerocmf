package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
	"zerocmf/common/bootstrap/database"
)

type Region struct {
	Code     string   `bson:"code" json:"code"`
	Name     string   `bson:"name" json:"name"`
	Children []Region `bson:"children,omitempty" json:"children"`
}

type sRegion struct {
	Code     int    `bson:"code" json:"code"`
	Name     string `bson:"name" json:"name"`
	Type     int    `bson:"type" json:"type"`
	ParentId int    `bson:"parentId" json:"parentId"`
}

func findRegion(region []Region, _type int, parentId int) (result []sRegion) {
	for _, v := range region {
		code, _ := strconv.Atoi(v.Code)

		str := strconv.Itoa(code)
		length := len(str)
		if length <= 6 {
			for i := 0; i < 6-length; i++ {
				str = str + "0"
			}
		}

		code, _ = strconv.Atoi(str)

		result = append(result, sRegion{
			Code:     code,
			Name:     v.Name,
			Type:     _type,
			ParentId: parentId,
		})

		if len(v.Children) > 0 {
			_type += 1
			children := findRegion(v.Children, _type, code)
			result = append(result, children...)
		}
	}
	return
}

func TestRegion(t *testing.T) {

	data, err := os.ReadFile("./region.json")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new District struct instance
	var region []Region

	// Unmarshal JSON data into the struct
	err = json.Unmarshal(data, &region)
	if err != nil {
		log.Fatal(err)
	}

	result := findRegion(region, 0, 0)

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

	collection := mongoDB.Collection("region")

	var documents []interface{}
	for _, r := range result {
		documents = append(documents, r)
	}

	collection.InsertMany(context.Background(), documents)

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
