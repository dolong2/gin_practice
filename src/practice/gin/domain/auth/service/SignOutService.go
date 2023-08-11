package service

import (
	"gin_practice/src/practice/gin/global/redis"
	"gin_practice/src/practice/gin/global/security"
)

func SignOut() error {
	user, err := security.GetUser()
	if err != nil {
		return err
	}
	redis.DeleteValue("RefreshToken", user.Email)
	return nil
}
