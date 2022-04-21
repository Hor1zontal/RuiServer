package db

import (
	"RuiServer/config"
	"RuiServer/exception"
	"context"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

var ctx context.Context

var collectionName = []string{
	"test",
	"user",
	"role",
	"course",
	"board",
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
	//coll.CreateOneIndex(context.Background(), options.IndexModel{Key: []string{"name"}, Unique: true})
	//coll.CreateIndexes(context.Background(), []options.IndexModel{{Key: []string{"id2", "id3"}}})
	updateCourseMap()
}

func InsertOne(collName string, data interface{}) {
	_, err := collectionMap[collName].InsertOne(ctx, data)
	if err != nil {
		exception.ExceptionCustom("insert one", exception.DatabaseError, err)
	}
}

func GenId(collName string) (int64, error) {
	IDInt64 := struct {
		Value int64 `bson:"max_id"`
	}{Value: 1}
	err := collectionMap[collName].Find(ctx, bson.M{}).Apply(qmgo.Change{Update: bson.M{"$inc": IDInt64}, Upsert: true, ReturnNew: true}, &IDInt64)
	return IDInt64.Value, err
}

func FindOne(collName string, query bson.M, data interface{}) error {
	return collectionMap[collName].Find(ctx, query).One(data)
}

// FindAll 返回query查询后总共的数量
func FindAll(collName string, query bson.M, pageNum, PageSize int, data interface{}) int {
	queryI := collectionMap[collName].Find(ctx, query)
	total, _ := queryI.Count()
	err := queryI.Limit(int64(PageSize)).Skip(int64((pageNum - 1) * PageSize)).All(data)
	if err != nil {
		exception.ExceptionCustom("FindAll", exception.DatabaseError, err)
	}
	return int(total)
}

func QueryCount(collName string, query bson.M) (int64, error) {
	return collectionMap[collName].Find(ctx, query).Count()
}

var CourseMap = make(map[int]string) //courseID-name

func updateCourseMap() {
	var courses []*DBCourse
	if err := collectionMap["course"].Find(ctx, bson.M{}).All(&courses); err != nil {
		exception.ExceptionCustom("update course map", exception.DatabaseError, err)
	}
	for _, course := range courses {
		CourseMap[course.CourseID] = course.CourseName
	}
}
