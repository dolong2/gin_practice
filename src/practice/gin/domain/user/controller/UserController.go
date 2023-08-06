package controller

import (
	"gin_practice/src/practice/gin/domain/user/controller/dto/request"
	"gin_practice/src/practice/gin/domain/user/service"
	"gin_practice/src/practice/gin/global/exception"
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
		err = service.CreateUser(*body)
		switch err.(type) {
		case exception.GlobalException:
			globalException := err.(exception.GlobalException)
			response := globalException.ToErrorResponse()
			c.JSON(globalException.Code, response)
		}
		c.Status(http.StatusCreated)
	})

	return r
}
