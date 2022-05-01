package service

import (
	"RuiServer/db"
	"github.com/gin-gonic/gin"
)

func UploadUser(username, sex, email, createTime, IDCard string, ClassID, MajorID, AcademyID, Age int) {
	db.NewUser(username, sex, email, createTime, IDCard, ClassID, MajorID, AcademyID, Age)
}

func GetUser(ID string) gin.H {
	user := db.GetUser(ID)
	return user.PackLoginRes()
}

const (
	role_admin   = 1
	role_student = 2
)

func GetUsersByQuery(ID, Name, SelectedAcademy, SelectedClass, SelectedMajor, startTime, endTime string, PageNum, PageSize int) (int, []*db.DBUser) {
	return db.GetUsersByQuery(ID, Name, SelectedAcademy, SelectedClass, SelectedMajor, startTime, endTime, PageNum, PageSize)
}
