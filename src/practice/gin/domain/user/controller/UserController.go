package controller

import (
	"fmt"
	"gin_practice/src/practice/gin/domain/user/controller/dto/request"
	"gin_practice/src/practice/gin/domain/user/exception"
	"gin_practice/src/practice/gin/domain/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpUserController(r *gin.Engine) *gin.Engine {
	r.POST("/users", func(c *gin.Context) {
		body := &request.SignUpRequest{}
		err := c.Bind(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		signUpUser(*body)
		c.Status(http.StatusCreated)
	})

	return r
}

func signUpUser(request request.SignUpRequest) {
	switch service.CreateUser(request) {
	case nil:
		break
	case exception.UserNotFoundException{}:
		fmt.Println("유저 찾을 수 없는 에러")
		break
	case exception.UserAlreadyExistsException{}:
		fmt.Println("유저가 이미 있을때 에러")
		break
	}
}
