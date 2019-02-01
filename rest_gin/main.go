package main

import (
	"github.com/gin-gonic/gin"
	"projects/rest_gin/app"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/instructions", app.GetInstructions)
		v1.GET("/instructions/:id", app.GetInstruction)
		v1.POST("/instructions", app.PostInstruction)
		v1.PUT("/instructions/:id", app.UpdateInstruction)
		v1.DELETE("/instructions/:id", app.DeleteInstruction)
	}

	return router
}

func main() {
	// Disable Console Color
	// // gin.DisableConsoleColor()

	// // Creates a gin router with default middleware:
	// // logger and recovery (crash-free) middleware
	// router := gin.Default()
	// v1 := router.Group("api/v1")
	// {
	// 	v1.GET("/instructions", app.GetInstructions)
	// 	v1.GET("/instructions/:id", app.GetInstruction)
	// 	v1.POST("/instructions", app.PostInstruction)
	// 	v1.PUT("/instructions", app.UpdateInstruction)
	// 	v1.DELETE("/instructions", app.DeleteInstruction)

	// }

	// router.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Hello world!")
	// })
	router := SetupRouter()
	router.Run()
	// router.Run(":3000") for a hard coded port
}
