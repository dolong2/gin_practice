package filter

import (
	"errors"
	"gin_practice/src/practice/gin/global/exception"
	"gin_practice/src/practice/gin/global/jwt"
	"gin_practice/src/practice/gin/global/security"
	"github.com/gin-gonic/gin"
)

func JwtReqFilter(c *gin.Context) {
	user, err := jwt.GetCurrentUser(c.Request)
	if err != nil {
		var globalException exception.GlobalException
		errors.As(err, &globalException)
		response := globalException.ToErrorResponse()
		c.JSON(globalException.Code, response)
		c.Abort()
	}
	security.SetUser(c, *user)
}
