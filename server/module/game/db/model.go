package db

import "time"

type DBUser struct {
	ID         int       `bson:"_id"`
	IDCard     string    `bson:"IDCard"`
	UserName   string    `bson:"userName"`
	Sex        string    `bson:"sex"`
	Email      string    `bson:"email"`
	ClassID    int       `bson:"class"`     //班级
	MajorID    int       `bson:"majorID"`   //专业
	AcademyID  int       `bson:"academyID"` //学院
	Age        int       `bson:"age"`
	Nation     string    `bson:"nation"`
	CreateTime time.Time `bson:"createTime"` //入学时间
}

type DBCourse struct {
	CourseID   int    `bson:"courseID"`
	CourseName string `bson:"courseName"`
}

//
//func (u *DBUser) CustomFields() field.CustomFieldsBuilder {
//	return field.NewCustom().SetCreateAt("CreateTimeAt").SetUpdateAt("UpdateTimeAt").SetId("_id")
//}

//type DBMajor struct {
//	MajorID   int    `bson:"majorID"`
//	MajorName string `bson:"majorName"`
//}
//

//type DBAcademy struct {
//	AcademyID   int    `bson:"majorID"`
//	AcademyName string `bson:"academyName"`
//}
