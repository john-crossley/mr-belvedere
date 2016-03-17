package main

import (
    "github.com/gin-gonic/gin"
    "github.com/john-crossley/mr-belvedere/db"
    "os"
	"github.com/john-crossley/mr-belvedere/api/jobs"
	"github.com/john-crossley/mr-belvedere/middlewares"
)

const (
    // Port at which the server starts listening
    Port = "8000"
)

func init() {
    db.Connect()
}

func main() {
    router := gin.Default()

	router.Use(middlewares.Connect)
	router.Use(middlewares.ErrorHandler)

    v1 := router.Group("api/v1")
    {
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"error": false, "message": "Mr Belvedere at your service..."})
		})
        v1.GET("/jobs", jobs.List)
		v1.GET("/jobs/:id", jobs.Find)
		v1.POST("/jobs", jobs.Create)
		v1.PUT("/jobs/:id", jobs.Update)
		v1.DELETE("/jobs/:id", jobs.Delete)
    }

    port := Port
    if len(os.Getenv("PORT")) > 0 {
        port = os.Getenv("PORT")
    }
    router.Run(":" + port)
}
