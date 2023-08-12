package security

import (
	"gin_practice/src/practice/gin/domain/user/entity"
	"gin_practice/src/practice/gin/global/security/exception"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) (*entity.User, error) {
	value, exists := c.Get("currentUser")
	if !exists {
		return nil, exception.ContextEmptyException()
	}
	user := value.(entity.User)
	return &user, nil
}

func SetUser(c *gin.Context, user entity.User) {
	c.Set("currentUser", user)
}
