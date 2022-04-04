package db

type DBUser struct {
	ID         int    `bson:"id"`
	IDCard     string `bson:"IDCard"`
	UserName   string `bson:"userName"`
	Sex        string `bson:"sex"`
	Email      string `bson:"email"`
	ClassID    int    `bson:"class"`     //班级
	MajorID    int    `bson:"majorID"`   //专业
	AcademyID  int    `bson:"academyID"` //学院
	Age        int    `bson:"age"`
	CreateTime string `bson:"createTime"` //入学时间
}

//
//func (u *DBUser) CustomFields() field.CustomFieldsBuilder {
//	return field.NewCustom().SetCreateAt("CreateTimeAt").SetUpdateAt("UpdateTimeAt").SetId("_id")
//}

type DBMajor struct {
	ID        int    `bson:"_id" gorm:""`
	MajorName string `bson:"majorName"`
}

type DBCourse struct {
	ID         int    `bson:"_id" gorm:"AUTO_INCREMENT"`
	CourseName string `bson:"courseNamename"`
}

type DBAcademy struct {
	ID          int    `bson:"_id" gorm:"AUTO_INCREMENT"`
	AcademyName string `bson:"academyName"`
}
