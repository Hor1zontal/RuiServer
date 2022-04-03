package api

import "github.com/gin-gonic/gin"

func DoLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"user": gin.H{
			"avatar": "",
			"status": 1,
		},
	})
}
