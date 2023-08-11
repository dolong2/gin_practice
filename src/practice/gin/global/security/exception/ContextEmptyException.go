package exception

import "gin_practice/src/practice/gin/global/exception"

func ContextEmptyException() exception.GlobalException {
	return exception.GlobalException{Code: 500, Msg: "컨텍스트가 비었음"}
}
