package main

import (
	"github.com/520MianXiangDuiXiang520/GoTools/dao"
	"lets_diagram/src"
	"log"
)

func initialization() {
	// 初始化配置文件
	src.InitSetting("./setting.json")
	// 初始化 MySQL 连接
	err := dao.InitDBSetting(src.GetSetting().MySQLSetting)
	if err != nil {
		log.Fatalf("fail to init MySQL connection: %v", err)
	}
	log.Println("SUCCESS TO INIT MYSQL !!!")

	// 初始化 Redis 连接
	err = dao.InitRedisPool(src.GetSetting().RedisSetting)
	if err != nil {
		log.Fatalf("fail to init redis connection: %v \n", err)
	}
	log.Println("SUCCESS TO INIT REDIS")
	// 初始化 mongoDB 连接
	// err = utils.InitMongoDBConn(src.GetSetting().MongoDBConn.Host,
	// 	src.GetSetting().MongoDBConn.Port, src.GetSetting().MongoDBConn.Database)
	// if err != nil {
	// 	log.Fatalf("fail to init MongoDB connection: %v", err)
	// }
}
