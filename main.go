package main

import (
	"core"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	//r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	r := gin.Default()

	r.Static("/css", "templates/css")                              //	Serve CSS Folder
	r.Static("/images", "templates/images")                        //	Serve Images Folder
	r.Static("/fonts", "templates/fonts")                          //	Serve Fonts Folder
	r.StaticFile("favicon.ico", "templates/resources/favicon.ico") //	Serve single File

	//	Process templates at start-up -> they don't have to be loaded from the disk again.
	r.LoadHTMLGlob("templates/*.tmpl") //	This fails if i don't do this, because it finds ./css directory
	r.LoadHTMLGlob("templates/*.html") //	^ Refers to usage of LoadHTMLGlob(templates/*)
	//	filepath.Join(os.Getenv("GOPATH")
	//r.Use(static.Serve("/templates/css/")) // static files have higher priority over dynamic routes
	//r.NotFound(static.Serve("/public"))

	//	Define route for the index page
	r.GET("/", func(c *gin.Context) {

		//	gin.Context.HTML() method -> Remder a template
		//	code int, name string, obj interface{}
		c.HTML(
			//	HTTP status -> 200
			http.StatusOK,
			//	Use templates/hello-world.html
			"hello-world.html",
			//	Pass in data to be rendered by the page
			gin.H{
				"title": "Home Page - Hello World",
			},
		)
	})

	//	Ping - Pong Test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//	gin.Default().	Methods()			->		GET(relativePath string, handlers ...HandlerFunc) IRoutes
	//	GET
	r.GET("/testGET")
	//	PUT
	r.PUT("/testPUT")
	//	POST
	r.POST("/testPOST")
	//	DELETE
	r.DELETE("/testDELETE")
	//	PATCH
	r.PATCH("/testPATCH")
	//	HEAD
	r.HEAD("/testHEAD")
	//	OPTIONS
	r.OPTIONS("/testHEAD")

	//	Template Examples
	//	Get -> Arrays
	r.GET("/array", func(c *gin.Context) {
		var values []int
		for i := 0; i < 10; i++ {
			values = append(values, i)
		}

		c.HTML(http.StatusOK, "arrays-example.tmpl", gin.H{"values": values})
	})

	//	Get ->	Array of Structs
	r.GET("/array_of_structs", func(c *gin.Context) {
		var values []Foo
		for i := 0; i < 5; i++ {
			values = append(values, Foo{IntegerValue: i, StringValue: strconv.Itoa(i)})
		}

		c.HTML(http.StatusOK, "array-struct-example.tmpl", gin.H{"values": values})
	})

	//	Get	->	Map
	r.GET("/map_example", func(c *gin.Context) {
		values := make(map[string]string)
		values["language"] = "Go"
		values["your"] = "Mom"
		values["what do we want?"] = "ubergolang skillz"
		values["when do we want em?"] = "NOW"

		c.HTML(http.StatusOK, "map-example.tmpl", gin.H{"myMap": values})
	})

	//	Get	->	Map & Keys
	r.GET("/map_and_keys", func(c *gin.Context) {
		values := make(map[string]string)
		values["language"] = "Go"
		values["your"] = "Mom"
		values["what do we want?"] = "ubergolang skillz"
		values["when do we want em?"] = "NOW"

		c.HTML(http.StatusOK, "map-and-keys-example.impl", gin.H{"myMap": values})
	})
	//	Template Examples	-	End	-	Check Hugo

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

	//Testing a simple Layout
	r.GET("/test-site", func(c *gin.Context) {

		c.HTML(
			http.StatusOK,
			"index.html",
			nil,
			//gin.H{
			//	"title": "My golang Testing Grounds",
			//},
		)
	})
	//	Above layout contains a link to this
	r.GET("/test_neo4j.html", func(c *gin.Context) {

		//ret, err := helloWorld("bolt://localhost:7687", "neo4j", "GinWebsite-Graph")
		ret, err := helloWorld()
		if err != nil {
			fmt.Println("Something went wrong!", err)
			fmt.Println("Ret _", ret)
		}
		fmt.Println("Ret AFTER _ ", ret)
		c.HTML(
			http.StatusOK,
			"test_neo4j.html",
			nil,
		)
	})

	r.Run(":8081")
}

type Foo struct {
	IntegerValue int
	StringValue  string
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
