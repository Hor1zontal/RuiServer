package http

import (
	"RuiServer/server/config"
	"RuiServer/server/exception"
	"RuiServer/server/module/game/http/api"
	"RuiServer/server/module/game/http/helper"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

func Init() {
	if config.Server.HTTPAddress == "" {
		log.Println("config.Server.HTTPAddress error!")
		return
	}
	r := gin.New()
	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.StaticFS("/", http.Dir("./dist"))
	//r.LoadHTMLGlob("./dist/*.html")
	//r.LoadHTMLFiles("static/*/*")
	r.Static("/static", "./dist/static")
	r.StaticFile("/login", "./dist/index.html")
	r.StaticFile("/favicon.ico", "./dist/favicon.ico")
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	r.Use(Recovery(), gin.Logger())

	r.Use(cors())
	r.POST("/web/doLogin", api.DoLogin)
	r.POST("/web/init/courseInfo", api.WebInitCourseInfo)
	r.POST("/web/studentInfo/upload", api.WebStudentInfoUpload)
	r.POST("/web/student/info", api.WebStudentInfo)

	r.GET("/web/user", api.WebUser)
	r.GET("/web/changeBoard/published", api.WebChangeBoardPublished)
	r.GET("/web/user/allRoles", api.WebUserAllRoles)
	r.GET("/web/school/details", api.WebSchoolDetails)
	r.GET("/web/student/personalInfo", api.WebStudentPersonalInfo)
	r.GET("/web/init/board", api.WebInitBoard)
	r.GET("/web/getBoard", api.WebGetBoard)
	r.GET("/web/board/type", api.WebBoardType)
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

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//log.Debug("error:%v",err)
				switch err.(type) {
				case exception.ErrorCode:
					helper.ResponseWithCode(c, err.(exception.ErrorCode))
				default:
					log.Println(err, "--", string(debug.Stack()))
					log.Println(err)
					debug.PrintStack()
					helper.ResponseWithCode(c, exception.InternalError)
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}
