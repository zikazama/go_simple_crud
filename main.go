package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		value := db
		c.JSON(http.StatusOK, gin.H{"data": value})
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		}
	})

	r.POST("/user", func(c *gin.Context) {

		// Parse JSON
		var json struct {
			Name  string `json:"name" binding:"required"`
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[json.Name] = json.Value
			c.JSON(http.StatusCreated, gin.H{"message": "create success"})
		}
	})

	r.PUT("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")

		_, ok := db[user]
		if ok {
			// Parse JSON
			var json struct {
				Value string `json:"value" binding:"required"`
			}

			if c.Bind(&json) == nil {
				db[user] = json.Value
				c.JSON(http.StatusOK, gin.H{"message": "update success"})
			}
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		}
	})

	r.DELETE("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")

		_, ok := db[user]
		if ok {
			delete(db, user)
			c.JSON(http.StatusNoContent, gin.H{"message": "remove success"})

		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
