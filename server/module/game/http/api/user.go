package api

import (
	"RuiServer/server/constant"
	"RuiServer/server/module/game/http/helper"
	"RuiServer/server/module/game/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func DoLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"user": gin.H{
			"avatar":      "", //头像
			"status":      1,  //1-超级管理员 2--学生
			"description": "",
		},
	})
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
	helper.FormatReq(c, req)
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

func WebGetBoard(c *gin.Context) {
	c.JSON(
		200,
		gin.H{
			"id":        111,
			"typeName":  "this is typeName",
			"content":   "this is content!!!",
			"title":     "this is title",
			"published": true,
		},
	)
}

func WebInitCourseInfo(c *gin.Context) {
	//1必修 2 选修 3 公选
	temp := gin.H{"id": 10001, "status": 1, "academy": gin.H{"id": 10001, "name": "academyName"}, "credit": 99, "total": 11, "period": 11, "createTime": "createTime", "updateTime": "createTime", "name": "courseName"}
	temp1 := gin.H{"id": 10002, "status": 2, "academy": gin.H{"id": 10002, "name": "academyName"}, "credit": 99, "total": 11, "period": 11, "createTime": "createTime", "updateTime": "createTime", "name": "courseName"}
	c.JSON(200, gin.H{
		"list": []gin.H{
			temp,
			temp1,
		},
		"total": 2,
	})
}

func WebChangeBoardPublished(c *gin.Context) {
	//改变publish的值
}

func WebInitBoard(c *gin.Context) {
	c.JSON(200, gin.H{
		"list": []gin.H{
			{"typeName": "typeName",
				"title":      "title",
				"published":  true, //
				"creatTime":  "createTime",
				"updateTime": "updateTime",
			},
		},
	})
}

func WebBoardType(c *gin.Context) {
	c.JSON(
		200, gin.H{
			"list": []gin.H{
				{"id": "10001", "title": "title", "createTime": time.Now().Format("2006-01-02 15:04:05")},
				{"id": "10002", "title": "title", "createTime": time.Now().Format("2006-01-02 15:04:05")},
				{"id": "10003", "title": "title", "createTime": time.Now().Format("2006-01-02 15:04:05")},
				{"id": "10004", "title": "title", "createTime": time.Now().Format("2006-01-02 15:04:05")},
				{"id": "10005", "title": "title", "createTime": time.Now().Format("2006-01-02 15:04:05")},
				{"id": "10006", "title": "title", "createTime": time.Now().Format("2006-01-02 15:04:05")},
			},
		},
	)
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
	helper.FormatReq(c, req)
	total, users := service.GetUsersByQuery(req.ID, req.Name, req.SelectedAcademy, req.SelectedClass, req.SelectedMajor, req.StartTime, req.EndTime, req.PageNum, req.PageSize)

	var listValue []gin.H
	for _, user := range users {
		listValue = append(listValue, gin.H{
			"id":         strconv.Itoa(user.ID),
			"name":       user.UserName,
			"sex":        user.Sex,
			"email":      user.Email,
			"classes":    gin.H{"id": user.ClassID, "name": constant.ClassesMap[user.ClassID]},
			"major":      gin.H{"id": user.MajorID, "name": constant.AcademyIDMajorMap[user.AcademyID][user.MajorID]},
			"academy":    gin.H{"id": user.AcademyID, "name": constant.AcademyMap[user.AcademyID]},
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
