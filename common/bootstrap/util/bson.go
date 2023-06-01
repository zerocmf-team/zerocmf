package util

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func AtoBsonM(obj interface{}) (bsonM bson.M, err error) {
	// 将结构体转换为 BSON 字节切片
	var bsonBytes []byte
	bsonBytes, err = bson.Marshal(obj)
	if err != nil {
		return
	}
	// 将 BSON 字节切片解码为 bson.M
	err = bson.Unmarshal(bsonBytes, &bsonM)
	if err != nil {
		log.Fatal("Error unmarshaling BSON to bson.M:", err)
	}
	return
}
