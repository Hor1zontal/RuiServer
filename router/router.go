package router

import (
	"RuiServer/config"
	"RuiServer/exception"
	"RuiServer/router/handler"
	"RuiServer/router/helper"
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
	//r.StaticFS("/", router.Dir("./dist"))
	//r.LoadHTMLGlob("./dist/*.html")
	//r.LoadHTMLFiles("static/*/*")

	//r.Static("/static", "./resource/dist/static")
	//r.StaticFile("/login", "./resource/dist/index.html")
	//r.StaticFile("/favicon.ico", "./resource/dist/favicon.ico")
	//r.GET("/", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "index.html", nil)
	//})

	r.Use(Recovery(), gin.Logger())

	r.Use(cors())
	r.POST("/web/doLogin", handler.DoLogin)
	r.POST("/web/init/courseInfo", handler.WebInitCourseInfo)

	r.POST("/web/studentInfo/upload", handler.WebStudentInfoUpload)
	r.POST("/web/student/info", handler.WebStudentInfo)

	r.GET("/web/user", handler.WebUser)

	r.GET("/web/changeBoard/published", handler.WebChangeBoardPublished)
	r.GET("/web/init/board", handler.WebInitBoard)
	r.GET("/web/getBoard", handler.WebGetBoard)
	r.GET("/web/board/type", handler.WebBoardType)
	r.POST("/web/add/board", handler.WebAddBoard)

	r.GET("/web/user/allRoles", handler.WebUserAllRoles)
	r.GET("/web/school/details", handler.WebSchoolDetails)
	r.GET("/web/student/personalInfo", handler.WebStudentPersonalInfo)

	r.GET("/web/academy/initAll", handler.WebAcademyInitAll)
	r.GET("/web/class/getByMajorId", handler.WebClassGetByMajorId)

	r.GET("/web/major/initAll", handler.WebMajorInitAll)
	r.GET("/web/major/getByAcademyId", handler.WebMajorGetByAcademyId)

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
