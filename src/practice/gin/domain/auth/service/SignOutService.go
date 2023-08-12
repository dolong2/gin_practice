package service

import (
	"gin_practice/src/practice/gin/global/redis"
	"gin_practice/src/practice/gin/global/security"
	"github.com/gin-gonic/gin"
)

func SignOut(c *gin.Context) error {
	user, err := security.GetUser(c)
	if err != nil {
		return err
	}
	redis.DeleteValue("RefreshToken", user.Email)
	return nil
}
