package news

import (
	"github.com/gin-gonic/gin"
)

type RouteInfo struct {
	Method string
	Route  string
}

var Routes []RouteInfo

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Маршруты
	newsGroup := router.Group("/news")
	{
		newsGroup.GET("/getNews", GetNews)
		newsGroup.POST("/createNews", CreateNews)
		newsGroup.PUT("/edit/:id", UpdateNews)
		newsGroup.DELETE("remove/:id/:newsId", DeleteNews)
	}

	return router
}
