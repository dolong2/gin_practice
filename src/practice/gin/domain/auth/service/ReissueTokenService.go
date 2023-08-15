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
	_, err := jwt.ValidTokenType(context.Request, jwt.RefreshType)
	if err != nil {
		return nil, err
	}
	user, err := security.GetUser(context)
	if err != nil {
		return nil, err
	}
	refreshToken, err := redis.GetValue("RefreshToken", user.Email)
	if err != nil {
		return nil, jwtException.JwtExpireException()
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
