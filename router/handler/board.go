package handler

import (
	"RuiServer/exception"
	"RuiServer/router/helper"
	"RuiServer/service"
	"RuiServer/utils"
	"github.com/gin-gonic/gin"
)

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
	temp := gin.H{"id": 10001, "status": 1, "academy": gin.H{"id": 1, "name": utils.AcademyMap[1]}, "credit": 2, "total": 100, "period": 36, "createTime": "2022-03-01", "updateTime": "2022-03-01", "name": "机器学习"}
	temp1 := gin.H{"id": 10002, "status": 2, "academy": gin.H{"id": 1, "name": utils.AcademyMap[1]}, "credit": 3, "total": 100, "period": 36, "createTime": "2022-03-01", "updateTime": "2022-03-01", "name": "人工智能"}
	temp2 := gin.H{"id": 10003, "status": 1, "academy": gin.H{"id": 1, "name": utils.AcademyMap[1]}, "credit": 2, "total": 200, "period": 36, "createTime": "2022-03-01", "updateTime": "2022-03-01", "name": "大数据"}
	temp3 := gin.H{"id": 10004, "status": 1, "academy": gin.H{"id": 1, "name": utils.AcademyMap[1]}, "credit": 1, "total": 120, "period": 18, "createTime": "2022-04-01", "updateTime": "2022-04-01", "name": "云计算"}
	temp4 := gin.H{"id": 10005, "status": 2, "academy": gin.H{"id": 1, "name": utils.AcademyMap[1]}, "credit": 2, "total": 100, "period": 36, "createTime": "2022-03-01", "updateTime": "2022-03-01", "name": "数据结构"}
	c.JSON(200, gin.H{
		"list": []gin.H{
			temp,
			temp1,
			temp2,
			temp3,
			temp4,
		},
		"total": 5,
	})
}

type reqBoardPublished struct {
	ID        int  `form:"id"`
	Published bool `form:"published"`
}

func WebChangeBoardPublished(c *gin.Context) {
	//改变publish的值
}

type reqInitBoard struct {
	Title    string `form:"keyword"`
	CurrPage int    `form:"currPage"`
	PageSize int    `form:"pageSize"`
}

type resInitBoard struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

func WebInitBoard(c *gin.Context) {
	req := &reqInitBoard{}
	helper.CheckReq(c, req)
	res := &resInitBoard{}
	res.Total, res.List = service.GetBoardByTitle(req.PageSize, req.CurrPage, req.Title)
	helper.ResponseWithData(c, res)
}

type reqWebBoardType struct {
	CurrPage int    `form:"currPage"`
	PageSize int    `form:"pageSize"`
	TypeName string `form:"typeName"`
}

type resWebBoardType struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

func WebBoardType(c *gin.Context) {
	req := &reqWebBoardType{}
	helper.CheckReq(c, req)
	res := &resWebBoardType{}
	res.Total, res.List = service.GetBoardByType(req.PageSize, req.CurrPage, req.TypeName)
	helper.ResponseWithData(c, res)
	//c.JSON(
	//	200, gin.H{
	//		"total": total,
	//		"list": []gin.H{
	//			{"id": "10001", "title": "title", "createTime": time.Now().Format("2006-01-02 15:04:05")},
	//			{"id": "10002", "title": "title", "createTime": time.Now().Format("2006-01-02 15:04:05")},
	//		},
	//	},
	//)
}

type reqAddBoard struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Published bool   `json:"published"`
	Title     string `json:"title"`
	TypeName  string `json:"typeName"`
}

func WebAddBoard(c *gin.Context) {
	req := &reqAddBoard{}
	helper.CheckReq(c, req)
	service.NewBoard(req.Content, req.Title, req.TypeName, req.Published)
	helper.ResponseWithCode(c, exception.CodeSuccess)
}
