package routes

import (
	"mysqlgin/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	comp := router.Group("company")
	{

		comp.GET("/", controllers.ShowAllComp)
		comp.GET("/:id", controllers.ShowComp)
		comp.POST("/", controllers.CreateComp)
		comp.PUT("/", controllers.UpdateComp)
		comp.DELETE("/:id", controllers.DeleteComp)
	}

	emp := router.Group("employee")
	{

		emp.GET("/", controllers.ShowAllEmpl)
		emp.GET("/:id", controllers.ShowEmpl)
		emp.POST("/", controllers.CreateEmpl)
		emp.PUT("/", controllers.UpdateEmpl)
		emp.DELETE("/:id", controllers.DeleteEmpl)
	}

	use := router.Group("user")
	{

		use.GET("/", controllers.ShowAllUse)
		use.GET("/:id", controllers.ShowUse)
		use.POST("/", controllers.CreateUse)
		use.PUT("/", controllers.UpdateUse)
		use.DELETE("/:id", controllers.DeleteUse)
	}

	return router
}
