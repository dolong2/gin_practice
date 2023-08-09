package exception

import "gin_practice/src/practice/gin/global/exception"

func ConnectionNullException() exception.GlobalException {
	return exception.GlobalException{Code: 500, Msg: "redis에 연결되지 않음"}
}
