package handler

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// 健康检查路由
	r.GET("/health", healthCheck)

	// API分组
	api := r.Group("/api")
	{
		api.POST("/open_ai/", openAiHandler)
	}
}

// healthCheck 健康检查处理函数
func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"version": "1.0.0",
	})
}
