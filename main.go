package main

import (
	"core"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	//r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//	Issue Cookie according to Gin Example https://gin-gonic.com/docs/examples/cookie/
	//									using https://stackoverflow.com/a/38418781
	//	Error is handled within
	r.GET("/cookie-pols", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			/*//	Set in case of no cookie message
			cookie = "NotSet"
			*/
			//	Cookie Not Set
			fmt.Println("[GIN] Remote Host Cookie - NotSet")

			//	Issue new Cookie using	./core/core.go
			newcval := core.GenCookie("username:password")
			fmt.Println(newcval)
			c.SetCookie("gin_cookie", newcval, 3600, "/", "localhost", false, false)
			cookie = newcval
			fmt.Println("[GIN] Remote Host Cookie - Issued -", cookie)
		}

		fmt.Printf("[GIN] Remote Host Cookie - %s \n", cookie)

	})

	r.Run(":8081")
}

/*
//	This portion is part of the Gin Quickstart Tutorial

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}
*/
