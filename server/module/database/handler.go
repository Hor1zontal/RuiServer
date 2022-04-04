package database

import (
	"RuiServer/server/config"
	"RuiServer/server/exception"
	"context"
	"github.com/qiniu/qmgo"
)

func Init() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: config.Server.Database.Address})
	if err != nil {
		exception.ExceptionCustom("server init", exception.DatabaseError, err)
	}
	defer func() {
		if err = client.Close(ctx); err != nil {
			exception.ExceptionCustom("server close", exception.DatabaseError, err)
		}
	}()
	db := client.Database("rui")
	db.Collection("test")
	//创建索引
	//client.CreateOneIndex(context.Background(), options.IndexModel{Key: []string{"name"}, Unique: true})
	//client.CreateIndexes(context.Background(), []options.IndexModel{{Key: []string{"id2", "id3"}}})
}
