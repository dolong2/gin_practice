package service

import (
	"gin_practice/src/practice/gin/domain/example/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GetUser(c *gin.Context) {
	user := c.Params.ByName("name")
	value, ok := repository.Get(user)
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}

func AddAdmin(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	// Parse JSON
	var json struct {
		Value string `json:"value" binding:"required"`
	}

	if c.Bind(&json) == nil {
		repository.Add(user, json.Value)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
