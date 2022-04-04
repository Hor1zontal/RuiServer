package http

import (
	"RuiServer/server/config"
	"RuiServer/server/module/game/http/api"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Init() {
	if config.Server.HTTPAddress == "" {
		log.Println("config.Server.HTTPAddress error!")
		return
	}
	r := gin.Default()
	r.Use(cors())
	r.POST("/web/doLogin", api.DoLogin)
	r.GET("/web/user", api.WebUser)
	r.POST("/web/studentInfo/upload", api.WebStudentInfoUpload)
	r.GET("/web/changeBoard/published", api.WebChangeBoardPublished)
	r.GET("/web/user/allRoles", api.WebUserAllRoles)
	r.GET("/web/school/details", api.WebSchoolDetails)
	r.GET("/web/student/personalInfo", api.WebStudentPersonalInfo)
	r.POST("/web/init/courseInfo", api.WebInitCourseInfo)
	r.GET("/web/init/board", api.WebInitBoard)
	r.GET("/web/getBoard", api.WebGetBoard)
	r.GET("/web/board/type", api.WebBoardType)
	r.POST("/web/student/info", api.WebStudentInfo)
	r.GET("/web/academy/initAll", api.WebAcademyInitAll)
	r.GET("/web/major/initAll", api.WebMajorInitAll)
	r.GET("/web/class/getByMajorId", api.WebClassGetByMajorId)
	r.GET("/web/major/getByAcademyId", api.WebMajorGetByAcademyId)
	r.Run("0.0.0.0:9000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, token, sign")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "false")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		//处理请求
		c.Next()
	}
}
