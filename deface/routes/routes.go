package routes

import (
	"deface/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	websiteRoutes := router.Group("/websites")
	{
		websiteRoutes.POST("", controllers.CreateWebsite)
		websiteRoutes.GET("", controllers.GetWebsites)
		websiteRoutes.GET("/:id", controllers.GetWebsiteByID)
	}

	return router
}
