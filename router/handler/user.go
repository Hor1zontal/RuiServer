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
				"name":     "test1", //
				"id":       10001,   //1-超级管理员 2--学生
				"sex":      "男",
				"phoneNum": "1175952275@qq.com",
				"enabled":  true,
				"avatar":   "https://joeschmoe.io/api/v1/random",
			},
			{
				"name":     "test2", //
				"id":       10002,   //1-超级管理员 2--学生
				"sex":      "男",
				"phoneNum": "1175952275@qq.com",
				"enabled":  true,
				"avatar":   "https://joeschmoe.io/api/v1/random",
			},
			{
				"name":     "test3", //
				"id":       10003,   //1-超级管理员 2--学生
				"sex":      "男",
				"phoneNum": "1175952275@qq.com",
				"enabled":  true,
				"avatar":   "https://joeschmoe.io/api/v1/random",
			},
			{
				"name":     "test4", //
				"id":       10004,   //1-超级管理员 2--学生
				"sex":      "男",
				"phoneNum": "1175952275@qq.com",
				"enabled":  true,
				"avatar":   "https://joeschmoe.io/api/v1/random",
			},
			{
				"name":     "test5", //
				"id":       10005,   //1-超级管理员 2--学生
				"sex":      "男",
				"phoneNum": "1175952275@qq.com",
				"enabled":  true,
				"avatar":   "https://joeschmoe.io/api/v1/random",
			},
			{
				"name":     "test6", //
				"id":       10006,   //1-超级管理员 2--学生
				"sex":      "男",
				"phoneNum": "1175952275@qq.com",
				"enabled":  true,
				"avatar":   "https://joeschmoe.io/api/v1/random",
			},
			{
				"name":     "test7", //
				"id":       10007,   //1-超级管理员 2--学生
				"sex":      "男",
				"phoneNum": "1175952275@qq.com",
				"enabled":  true,
				"avatar":   "https://joeschmoe.io/api/v1/random",
			},
			{
				"name":     "test8", //
				"id":       10008,   //1-超级管理员 2--学生
				"sex":      "男",
				"phoneNum": "1175952275@qq.com",
				"enabled":  true,
				"avatar":   "https://joeschmoe.io/api/v1/random",
			},
			{
				"name":     "test9", //
				"id":       10009,   //1-超级管理员 2--学生
				"sex":      "男",
				"phoneNum": "1175952275@qq.com",
				"enabled":  true,
				"avatar":   "https://joeschmoe.io/api/v1/random",
			},
			{
				"name":     "test10", //
				"id":       10010,    //1-超级管理员 2--学生
				"sex":      "男",
				"phoneNum": "1175952275@qq.com",
				"enabled":  true,
				"avatar":   "https://joeschmoe.io/api/v1/random",
			},
		},
		"total": 10,
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
	var res []gin.H
	for id, name := range utils.AcademyMap {
		res = append(res, gin.H{"id": id, "name": name})
	}
	//temp := gin.H{"id": 10001, "name": "academyName"}
	//temp1 := gin.H{"id": 10002, "name": "academyName"}
	c.JSON(200, res)
}

func WebMajorInitAll(c *gin.Context) {
	//var res []gin.H
	//for id, name := range utils.AcademyIDMajorMap {
	//	res = append(res, gin.H{"id": id, "name": name})
	//}
	//temp := gin.H{"id": 10001, "name": "majorName"}
	//temp1 := gin.H{"id": 10002, "name": "majorName"}
	//c.JSON(200, []gin.H{
	//	temp,
	//	temp1,
	//})
}

func WebClassGetByMajorId(c *gin.Context) {
	var res []gin.H
	for id, name := range utils.ClassesMap {
		res = append(res, gin.H{"id": id, "name": name})
	}
	//temp := gin.H{"id": 10001, "name": "className"}
	//temp1 := gin.H{"id": 10002, "name": "className"}
	c.JSON(200, res)

}

type ReqMajorGetByAcademyId struct {
	ID int `form:"academyId"`
}

func WebMajorGetByAcademyId(c *gin.Context) {
	req := &ReqMajorGetByAcademyId{}
	helper.CheckReq(c, req)
	//temp := gin.H{"id": 10001, "name": "majorName"}
	//temp1 := gin.H{"id": 10002, "name": "majorName"}
	var res []gin.H
	//id, _ := strconv.Atoi(req.ID)
	for id, name := range utils.AcademyIDMajorMap[req.ID] {
		res = append(res, gin.H{"id": id, "name": name})
	}
	c.JSON(200, res)
}
