package service

import (
	"gin_practice/src/practice/gin/domain/auth/controller/dto/request"
	"gin_practice/src/practice/gin/domain/auth/controller/dto/response"
	"gin_practice/src/practice/gin/domain/auth/exception"
	userException "gin_practice/src/practice/gin/domain/user/exception"
	userRepository "gin_practice/src/practice/gin/domain/user/repository"
	"gin_practice/src/practice/gin/global/jwt"
	jwtException "gin_practice/src/practice/gin/global/jwt/exception"
)

func SignIn(request request.SignInRequest) (response.SignInResponse, error) {
	if !userRepository.ExistsByEmail(request.Email) {
		return response.SignInResponse{}, userException.UserNotFoundException()
	}
	user, _ := userRepository.GetByEmail(request.Email)
	if user.Password != request.Password {
		return response.SignInResponse{}, exception.PasswordMisMatchException()
	}
	accessToken, accessErr := jwt.CreateAccessToken(user.Email)
	refreshToken, refreshErr := jwt.CreateRefreshToken(user.Email)
	if accessErr != nil || refreshErr != nil {
		return response.SignInResponse{}, jwtException.JwtGenerateException()
	}
	return response.ToTokenResponse(accessToken, refreshToken), nil
}
