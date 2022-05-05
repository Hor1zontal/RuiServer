package handler

import (
	"RuiServer/router/helper"
	"RuiServer/service"
	"RuiServer/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ReqDoLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func DoLogin(c *gin.Context) {

	req := &ReqDoLogin{}
	helper.FormatReq(c, req)
	user := service.GetUser(req.Username)
	//c.JSON(200, gin.H{
	//	"user": gin.H{
	//		"avatar":      "https://cdn.jsdelivr.net/gh/Hor1zontal/BLOG_IMG/2022/3/202204252141295.png", //头像
	//		"status":      1,
	//		"description": "",
	//		"role":        1, //1-超级管理员 2--学生
	//	},
	//})
	c.JSON(200, user)
}

type ReqWebStudentInfoUpload struct {
	AcademyID  int    `json:"academyId"`
	Age        int    `json:"age"`
	ClassID    int    `json:"classId"`
	CreateTime string `json:"createTime"`
	Email      string `json:"email"`
	IDCard     string `json:"idCard"`
	MajorID    int    `json:"majorId"`
	Name       string `json:"name"`
	Sex        string `json:"sex"`
}

func WebStudentInfoUpload(c *gin.Context) {
	req := &ReqWebStudentInfoUpload{}
	helper.CheckReq(c, req)
	service.UploadUser(req.Name, req.Sex, req.Email, req.CreateTime, req.IDCard, req.ClassID, req.MajorID, req.AcademyID, req.Age)
	c.JSON(200, gin.H{})
	//println(req.Age)

}

func WebUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"list": []gin.H{
			{
				"name":     "邓飞宇", //
				"id":       1,     //1-超级管理员 2--学生
				"sex":      "男",
				"phoneNum": "123456789",
				"enabled":  true,
				"avatar":   "https://cdn.jsdelivr.net/gh/Hor1zontal/BLOG_IMG/2022/3/202204252141295.png",
			},
			{
				"name":     "睿子", //
				"id":       2,    //1-超级管理员 2--学生
				"sex":      "女",
				"phoneNum": "123456789",
				"enabled":  true,
				"avatar":   "https://cdn.jsdelivr.net/gh/Hor1zontal/BLOG_IMG/2022/3/202204252141295.png",
			},
		},
		"total": 2,
	})
}

func WebUserAllRoles(c *gin.Context) {

}

func WebSchoolDetails(c *gin.Context) {

}

func WebStudentPersonalInfo(c *gin.Context) {

}

type ReqWebStudentPersonalCourses struct {
	ID string `form:"id"` //学生id
}

func WebStudentPersonalCourses(c *gin.Context) {
	req := &ReqWebStudentPersonalCourses{}
	helper.CheckReq(c, req)
	id, _ := strconv.Atoi(req.ID)
	res := service.GetUserCourses(id)
	c.JSON(200, res)
}

type ReqWebStudentInfo struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	PageNum         int    `json:"pageNum"`
	PageSize        int    `json:"pageSize"`
	SelectedAcademy string `json:"selectedAcademy"`
	SelectedClass   string `json:"selectedClass"`
	SelectedMajor   string `json:"selectedMajor"`
	StartTime       string `json:"startTime"`
	EndTime         string `json:"endTime"`
}

func WebStudentInfo(c *gin.Context) {
	req := &ReqWebStudentInfo{}
	helper.CheckReq(c, req)
	total, users := service.GetUsersByQuery(req.ID, req.Name, req.SelectedAcademy, req.SelectedClass, req.SelectedMajor, req.StartTime, req.EndTime, req.PageNum, req.PageSize)

	var listValue []gin.H
	for _, user := range users {
		listValue = append(listValue, gin.H{
			"id":         strconv.Itoa(user.ID),
			"name":       user.UserName,
			"sex":        user.Sex,
			"email":      user.Email,
			"classes":    gin.H{"id": user.ClassID, "name": utils.ClassesMap[user.ClassID]},
			"major":      gin.H{"id": user.MajorID, "name": utils.AcademyIDMajorMap[user.AcademyID][user.MajorID]},
			"academy":    gin.H{"id": user.AcademyID, "name": utils.AcademyMap[user.AcademyID]},
			"nation":     user.Nation,
			"createTime": user.CreateTime.Local().Format("2006-01-02"),
		})
	}
	c.JSON(
		200, gin.H{
			"list":  listValue,
			"total": total,
		},
	)
}

func WebAcademyInitAll(c *gin.Context) {
	temp := gin.H{"id": 10001, "name": "academyName"}
	temp1 := gin.H{"id": 10002, "name": "academyName"}
	c.JSON(200, []gin.H{
		temp,
		temp1,
	})
}

func WebMajorInitAll(c *gin.Context) {
	temp := gin.H{"id": 10001, "name": "majorName"}
	temp1 := gin.H{"id": 10002, "name": "majorName"}
	c.JSON(200, []gin.H{
		temp,
		temp1,
	})
}

func WebClassGetByMajorId(c *gin.Context) {
	temp := gin.H{"id": 10001, "name": "className"}
	temp1 := gin.H{"id": 10002, "name": "className"}
	c.JSON(200, []gin.H{
		temp,
		temp1,
	})

}

func WebMajorGetByAcademyId(c *gin.Context) {
	temp := gin.H{"id": 10001, "name": "majorName"}
	temp1 := gin.H{"id": 10002, "name": "majorName"}
	c.JSON(200, []gin.H{
		temp,
		temp1,
	})
}
