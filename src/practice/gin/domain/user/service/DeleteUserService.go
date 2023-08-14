package service

import (
	"gin_practice/src/practice/gin/domain/user/repository"
	"gin_practice/src/practice/gin/global/redis"
	"gin_practice/src/practice/gin/global/security"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) error {
	user, err := security.GetUser(c)
	if err != nil {
		return err
	}
	repository.DeleteByEmail(user.Email)
	redis.DeleteValue("RefreshToken", user.Email)
	return nil
}
