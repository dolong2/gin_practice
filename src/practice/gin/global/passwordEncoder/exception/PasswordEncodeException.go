package exception

import "gin_practice/src/practice/gin/global/exception"

func PasswordEncodeException() exception.GlobalException {
	return exception.GlobalException{Code: 500, Msg: "패스워드를 인코딩할 수 없음"}
}
