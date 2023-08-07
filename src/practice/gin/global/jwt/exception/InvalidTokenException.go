package exception

import "gin_practice/src/practice/gin/global/exception"

func InvalidTokenException() exception.GlobalException {
	return exception.GlobalException{Code: 401, Msg: "옳바르지 않은 토큰"}
}
