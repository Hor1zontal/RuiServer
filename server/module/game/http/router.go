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
