package service

import (
	"gin_practice/src/practice/gin/domain/auth/controller/dto/request"
	"gin_practice/src/practice/gin/domain/auth/controller/dto/response"
	"gin_practice/src/practice/gin/domain/auth/exception"
	userException "gin_practice/src/practice/gin/domain/user/exception"
	userRepository "gin_practice/src/practice/gin/domain/user/repository"
	"gin_practice/src/practice/gin/global/jwt"
	jwtException "gin_practice/src/practice/gin/global/jwt/exception"
	"gin_practice/src/practice/gin/global/passwordEncoder"
	"gin_practice/src/practice/gin/global/redis"
	"time"
)

func SignIn(request request.SignInRequest) (*response.SignInResponse, error) {
	if !userRepository.ExistsByEmail(request.Email) {
		return nil, userException.UserNotFoundException()
	}
	user, _ := userRepository.GetByEmail(request.Email)
	if !passwordEncoder.MatchPassword(request.Password, user.Password) {
		return nil, exception.PasswordMisMatchException()
	}
	accessToken, accessErr := jwt.CreateAccessToken(user.Email)
	refreshToken, refreshErr := jwt.CreateRefreshToken(user.Email)
	redis.SaveValue("RefreshToken", user.Email, refreshToken, int(jwt.GetRefreshExp().Microseconds()))
	if accessErr != nil || refreshErr != nil {
		return nil, jwtException.JwtGenerateException()
	}
	accessExp := jwt.GetAccessExp()
	refreshExp := jwt.GetRefreshExp()
	tokenResponse := response.ToTokenResponse(accessToken, refreshToken, time.Now().Add(accessExp), time.Now().Add(refreshExp))
	return &tokenResponse, nil
}
