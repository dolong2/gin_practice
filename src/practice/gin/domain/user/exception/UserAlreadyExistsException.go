package exception

import (
	"gin_practice/src/practice/gin/global/exception"
	"net/http"
)

func UserAlreadyExistsException() exception.GlobalException {
	return exception.InitGlobalException(http.StatusBadRequest, "이미 해당 유저가 존재함")
}
