package db

import (
	"RuiServer/server/config"
	"RuiServer/server/exception"
	"context"
	"github.com/qiniu/qmgo"
)

var ctx context.Context

var collectionName = []string{
	"test",
	"user",
	"role",
}

var collectionMap = make(map[string]*qmgo.Collection)

//var DataBaseHandler *qmgo.Client
func Init() {
	ctx = context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: config.Server.Database.Address})
	if err != nil {
		exception.ExceptionCustom("server init", exception.DatabaseError, err)
	}

	db := client.Database("rui")

	for _, collName := range collectionName {
		collectionMap[collName] = db.Collection(collName)
	}
	//创建索引
	//client.CreateOneIndex(context.Background(), options.IndexModel{Key: []string{"name"}, Unique: true})
	//client.CreateIndexes(context.Background(), []options.IndexModel{{Key: []string{"id2", "id3"}}})

}

func InsertOne(collName string, data interface{}) {
	if result, err := collectionMap[collName].InsertOne(ctx, data); err != nil {
		println(err.Error())
	} else {
		println(result)
	}
}
