package controller

import (
	"fmt"
	"gin_practice/src/practice/gin/domain/auth/controller/dto/request"
	"gin_practice/src/practice/gin/domain/auth/service"
	"gin_practice/src/practice/gin/global/exception"
	"gin_practice/src/practice/gin/global/security/filter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpAuthController(r *gin.Engine) *gin.Engine {
	group := r.Group("/auth")
	group.Use(filter.JwtReqFilter)
	r.POST("/auth", func(c *gin.Context) {
		body := &request.SignInRequest{}
		err := c.Bind(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		response, err := service.SignIn(*body)
		switch err.(type) {
		case exception.GlobalException:
			fmt.Println(err.Error())
			globalException := err.(exception.GlobalException)
			response := globalException.ToErrorResponse()
			c.JSON(globalException.Code, response)
			return
		}
		c.JSON(http.StatusOK, response)
	})

	group.DELETE("", func(c *gin.Context) {
		err := service.SignOut()
		switch err.(type) {
		case exception.GlobalException:
			fmt.Println(err.Error())
			globalException := err.(exception.GlobalException)
			response := globalException.ToErrorResponse()
			c.JSON(globalException.Code, response)
			return
		}
		c.Status(http.StatusOK)
	})

	return r
}
