package service

import (
	"RuiServer/db"
	"github.com/gin-gonic/gin"
	"strconv"
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

func GetUserCourses(ID int) interface{} {
	var res []resResult
	results := db.GetUserResults(ID)
	user := db.GetUser(strconv.Itoa(ID))
	for _, v := range results.Results {
		tmp := resResult{
			UserID:     user.ID,
			Name:       user.UserName,
			CourseName: GetCourseName(v.CourseID),
			Score:      v.Score,
		}
		res = append(res, tmp)
	}
	return res
}

type resResult struct {
	UserID     int    `json:"userID"`
	Name       string `json:"name"`
	CourseName string `json:"courseName"`
	Score      int    `json:"score"`
}

func GetCourseName(ID int) string {
	return db.GetCourseName(ID)
}
