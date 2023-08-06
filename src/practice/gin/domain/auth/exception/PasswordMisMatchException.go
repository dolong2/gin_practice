package exception

import (
	"gin_practice/src/practice/gin/global/exception"
	"net/http"
)

func PasswordMisMatchException() exception.GlobalException {
	exception := exception.GlobalException{}
	exception.Msg = "비밀번호가 일치하지 않음"
	exception.Code = http.StatusBadRequest
	return exception
}
