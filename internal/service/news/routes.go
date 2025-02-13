package news

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Маршруты
	newsGroup := router.Group("/news")
	{
		newsGroup.GET("", GetNews)
		newsGroup.POST("", CreateNews)
		newsGroup.PUT("/:id", UpdateNews)
		newsGroup.DELETE("/:id", DeleteNews)
	}

	return router
}
