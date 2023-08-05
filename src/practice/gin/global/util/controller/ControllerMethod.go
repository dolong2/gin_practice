package controller

import (
	"gin_practice/src/practice/gin/domain/user/controller/dto/request"
	"gin_practice/src/practice/gin/global/exception"
	"github.com/gin-gonic/gin"
	"go/types"
	"net/http"
)

type Request interface {
	request.SignUpRequest | types.Nil
}

type Response interface {
}

func POST[R Request](r *gin.Engine, url string, body R, responseStatus int, execute func(request R) error) {
	r.POST(url, func(c *gin.Context) {
		body := &body
		err := c.Bind(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		err = execute(*body)
		switch err.(type) {
		case exception.GlobalException:
			globalException := err.(exception.GlobalException)
			response := globalException.ToErrorResponse()
			c.JSON(globalException.Code, response)
		}
		c.Status(responseStatus)
	})
}

func GET(r *gin.Engine, url string, f func(c *gin.Context)) {
	r.GET(url, f)
}

func PUT(r *gin.Engine, url string, f func(c *gin.Context)) {
	r.POST(url, f)
}

func PATCH(r *gin.Engine, url string, f func(c *gin.Context)) {
	r.PATCH(url, f)
}

func DELETE(r *gin.Engine, url string, f func(c *gin.Context)) {
	r.DELETE(url, f)
}
