package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"strconv"
	"time"
)

type MongoDB struct {
	client   *mongo.Client   `json:",optional"`
	database *mongo.Database `json:",optional"`
	Host     string
	DbName   string
	Username string
	Password string
	Port     int
	Prefix   string
	AuthCode string
}

type Mongo struct {
	Host     string
	DbName   string
	Username string
	Password string
	Port     int
	Prefix   string
	AuthCode string
}

// 连接到MongoDB
func NewMongoDB(mongoDb Mongo, dbName string) (db MongoDB, err error) {

	if dbName == "" {
		dbName = mongoDb.DbName
	} else {
		dbName = "site_" + dbName
	}

	clientOptions := mongoDb.newConn()
	var client *mongo.Client
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return
	}

	db = MongoDB{
		client:   client,
		database: client.Database(dbName),
		Host:     mongoDb.Host,
		DbName:   dbName,
		Username: mongoDb.Username,
		Password: mongoDb.Password,
		Port:     mongoDb.Port,
		Prefix:   mongoDb.Prefix,
		AuthCode: mongoDb.AuthCode,
	}

	return
}

// 获取集合
func (mDB *MongoDB) Collection(name string) *mongo.Collection {
	return mDB.database.Collection(name)
}

// 查询文档
func (mDB *MongoDB) FindOne(collection *mongo.Collection, filter interface{}, result interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

// 更新文档
func (mDB *MongoDB) UpdateOne(collection *mongo.Collection, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 删除文档
func (mDB *MongoDB) DeleteOne(collection *mongo.Collection, filter interface{}) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 插入文档

func (mDB *MongoDB) InsertOne(collection *mongo.Collection, document interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (mDB *MongoDB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := mDB.client.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}

func (db *Mongo) newConn() (clientOptions *options.ClientOptions) {
	username := db.Username
	pwd := db.Password
	host := db.Host
	HOST := os.Getenv("MYSQL_HOST")
	if HOST != "" {
		host = HOST
	}
	port := strconv.Itoa(db.Port)
	uri := "mongodb://" + username + ":" + pwd + "@" + host + ":" + port
	clientOptions = options.Client().ApplyURI(uri)
	return

}

func (mDB *MongoDB) Ping() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = mDB.client.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}
	return
}
