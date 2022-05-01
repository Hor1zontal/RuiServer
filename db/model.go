package db

import "time"

type DBUser struct {
	ID          int       `bson:"_id"`         //学号
	IDCard      string    `bson:"IDCard"`      //身份证
	UserName    string    `bson:"userName"`    //姓名
	Sex         string    `bson:"sex"`         //性别
	Email       string    `bson:"email"`       //email
	ClassID     int       `bson:"class"`       //班级
	MajorID     int       `bson:"majorID"`     //专业
	AcademyID   int       `bson:"academyID"`   //学院
	Age         int       `bson:"age"`         //年龄
	Nation      string    `bson:"nation"`      //民族
	CreateTime  time.Time `bson:"createTime"`  //入学时间
	Status      int       `bson:"status"`      //用户角色 1--管理员 2--学生
	Avatar      string    `bson:"avatar"`      //头像
	Description string    `bson:"description"` //描述
	Enabled     bool      `bson:"enabled"`     //封禁状态  1--可用  0--封禁
}

type DBCourse struct {
	CourseID   int    `bson:"courseID"`   //课程ID
	CourseName string `bson:"courseName"` //课程名
}

type DBBoard struct {
	ID         int       `bson:"id" json:"id"`
	Title      string    `bson:"title" json:"title"`
	Content    string    `bson:"content"json:"content"`
	CreateTime time.Time `bson:"createTime" json:"createTime"`
	TypeName   string    `bson:"typeName" json:"typeName"`
	Published  bool      `bson:"published" json:"published"`
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
