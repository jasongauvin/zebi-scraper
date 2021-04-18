package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/zebi-scraper/api/controllers"
	"github.com/jasongauvin/zebi-scraper/api/middleware"
)

// InitializeRoutes set up the routes for the server
func InitializeRoutes(r *gin.Engine) {
	// HTML routes for the GUI
	r.GET("/", controllers.SayHello)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/scrapi/:name", controllers.GetFact)


	// API routes
	api := r.Group("/api")
	// Check Authorization for api grouip
	api.Use(middleware.CheckAuthorization)

	{
		// api.POST("/scapi/:id/instagram", controllers.TestSearch)
	}
}
