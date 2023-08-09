package main

import (
	"gin_practice/src/practice/gin/domain/auth/controller"
	exampleController "gin_practice/src/practice/gin/domain/example/controller"
	userController "gin_practice/src/practice/gin/domain/user/controller"
	"gin_practice/src/practice/gin/global/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	redis.SetUpRedis()
	defer redis.CloseConn()
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
