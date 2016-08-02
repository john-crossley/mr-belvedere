package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/john-crossley/mr-belvedere/db"
	"net/http"
	"os"
	"os"
)

func respondWithError(code int, message string, c *gin.Context) {
	response := map[string]string{"error": message}

	c.JSON(code, response)
	c.Abort()
}

func TokenAuthenticate(c *gin.Context) {
	token := c.Request.FormValue("api_token")

	if token == "" {
		respondWithError(http.StatusUnauthorized, "API token required", c)
		return
	}

	if token != os.Getenv("API_TOKEN") {
		respondWithError(http.StatusUnauthorized, "Invalid API token", c)
		return
	}

	c.Next()
}

// Connect middleware clones the database session for each request
// and makes the `db` object available for each handler
func Connect(c *gin.Context) {
	session := db.Session.Clone()
	defer session.Close()

	c.Set("db", session.DB(db.Mongo.Database))
	c.Next()
}

// ErrorHandler is a middleware to handle errors encountered during requests.
func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": true, "message": c.Errors.String()})
	}
}
