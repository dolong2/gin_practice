package main

import (
	"gin_practice/src/practice/gin/domain/auth/controller"
	exampleController "gin_practice/src/practice/gin/domain/example/controller"
	userController "gin_practice/src/practice/gin/domain/user/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	engine := gin.Default()
	exampleController.SetupRouter(engine)
	userController.SetUpUserController(engine)
	controller.SetUpAuthController(engine)
	return engine
}
