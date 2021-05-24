package utils

// 用于初始化 mongodb 连接

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
)

var client *qmgo.Client

func InitMongoDBConn(host string, port int, database string) error {
    var err error
	ctx := context.Background()
	uri := fmt.Sprintf("mongodb://%s:%d", host, port)
	client, err = qmgo.NewClient(ctx, &qmgo.Config{Uri: uri})
	if err != nil {
		return err
	}
	mongoDB = client.Database(database)
	return nil
}

var mongoDB *qmgo.Database

func GetMgoDB() *qmgo.Database {
	if mongoDB == nil {
		panic("mongodb connection is not initialized")
	}
	return mongoDB
}

func GetMgoClient() *qmgo.Client {
    if client == nil {
        panic("mongodb connection is not initialized")
    }
    return client
}

func GetMgoCollection(collection string) *qmgo.Collection {
	if mongoDB == nil {
		panic("mongodb connection is not initialized")
	}
	return mongoDB.Collection(collection)
}
