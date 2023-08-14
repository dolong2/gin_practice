package exception

import "gin_practice/src/practice/gin/global/exception"

func JwtExpireException() exception.GlobalException {
	return exception.GlobalException{Msg: "토큰이 만료되었습니다.", Code: 401}
}
