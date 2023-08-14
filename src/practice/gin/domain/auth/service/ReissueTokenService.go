package service

import (
	"gin_practice/src/practice/gin/domain/auth/controller/dto/response"
	"gin_practice/src/practice/gin/global/jwt"
	jwtException "gin_practice/src/practice/gin/global/jwt/exception"
	"gin_practice/src/practice/gin/global/redis"
	"gin_practice/src/practice/gin/global/security"
	"github.com/gin-gonic/gin"
	"time"
)

func ReissueToken(context *gin.Context) (*response.ReissueTokenResponse, error) {
	user, err := security.GetUser(context)
	if err != nil {
		return nil, err
	}
	accessToken, accessErr := jwt.CreateAccessToken(user.Email)
	refreshToken, refreshErr := jwt.CreateRefreshToken(user.Email)
	redis.SaveValue("RefreshToken", user.Email, refreshToken, int(jwt.GetRefreshExp().Microseconds()))
	if accessErr != nil || refreshErr != nil {
		return nil, jwtException.JwtGenerateException()
	}
	accessExp := jwt.GetAccessExp()
	refreshExp := jwt.GetRefreshExp()
	tokenResponse := response.ToReissueTokenResponse(accessToken, refreshToken, time.Now().Add(accessExp), time.Now().Add(refreshExp))
	return &tokenResponse, nil
}
