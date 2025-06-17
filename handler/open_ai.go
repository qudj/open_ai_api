package handler

import (
	"github.com/gin-gonic/gin"
)

func openAiHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"version": "1.0.0",
	})
}
