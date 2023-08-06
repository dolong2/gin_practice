package exception

import "gin_practice/src/practice/gin/global/exception"

func JwtGenerateException() exception.GlobalException {
	return exception.GlobalException{Code: 500, Msg: "jwt 생성중 에러"}
}
