package controller

import (
	"errors"
	"fmt"
	"gin_practice/src/practice/gin/domain/auth/controller/dto/request"
	"gin_practice/src/practice/gin/domain/auth/service"
	"gin_practice/src/practice/gin/global/exception"
	"gin_practice/src/practice/gin/global/security/filter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpAuthController(g *gin.RouterGroup) {
	group := g.Group("/auth")
	group.Use(filter.JwtReqFilter)
	g.POST("/auth", func(c *gin.Context) {
		body := &request.SignInRequest{}
		err := c.Bind(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		response, err := service.SignIn(*body)
		var globalException exception.GlobalException
		switch {
		case errors.As(err, &globalException):
			fmt.Println(err.Error())
			var globalException exception.GlobalException
			errors.As(err, &globalException)
			response := globalException.ToErrorResponse()
			c.JSON(globalException.Code, response)
			return
		}
		c.JSON(http.StatusOK, response)
	})

	group.DELETE("", func(c *gin.Context) {
		err := service.SignOut(c)
		var globalException exception.GlobalException
		switch {
		case errors.As(err, &globalException):
			fmt.Println(err.Error())
			var globalException exception.GlobalException
			errors.As(err, &globalException)
			response := globalException.ToErrorResponse()
			c.JSON(globalException.Code, response)
			return
		}
		c.Status(http.StatusOK)
	})

	group.PATCH("", func(c *gin.Context) {
		response, err := service.ReissueToken(c)
		var globalException exception.GlobalException
		switch {
		case errors.As(err, &globalException):
			fmt.Println(err.Error())
			var globalException exception.GlobalException
			errors.As(err, &globalException)
			response := globalException.ToErrorResponse()
			c.JSON(globalException.Code, response)
			return
		}
		c.JSON(http.StatusOK, response)
	})
}
