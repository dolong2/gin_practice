package exception

import (
	"gin_practice/src/practice/gin/global/exception"
	"net/http"
)

func UserNotFoundException() exception.GlobalException {
	return exception.InitGlobalException(http.StatusNotFound, "해당 유저를 찾을 수 없음")
}
