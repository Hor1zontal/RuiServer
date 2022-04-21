package handler

import (
	"RuiServer/exception"
	"RuiServer/router/helper"
	"RuiServer/service"
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
