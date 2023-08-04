package main

import (
	"gin_practice/src/practice/gin/domain/example/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	engine := gin.Default()
	controller.SetupRouter(engine)
	return engine
}
