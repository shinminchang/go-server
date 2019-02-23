package route

import (
	"github.com/gin-gonic/gin"
	"go-server/controllers"
)

func Route(){
	router := gin.Default()
	test := router.Group("/api/v1/test")
	{
		test.GET("/:id", controllers.Tests.Test)
		test.PUT("/:id", controllers.Tests.Test)
		test.DELETE("/:id", controllers.Tests.Test)
	}

	member := router.Group("/api/v1/new")
	{
		member.GET("", controllers.News.News)
	}

	databases := router.Group("/api/v1/db")
	{
		databases.GET("", controllers.Dbtest.Test)
	}
	router.Run()
}