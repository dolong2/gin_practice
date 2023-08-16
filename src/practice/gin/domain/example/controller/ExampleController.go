package controller

import (
	"gin_practice/src/practice/gin/domain/example/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// Disable Console Color
	// gin.DisableConsoleColor()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		service.Pong(c)
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		service.GetUser(c)
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

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		service.AddAdmin(c)
	})
}
