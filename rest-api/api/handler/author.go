package handler

import "github.com/gin-gonic/gin"

func ListAuthors(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List Authors",
	})
}
