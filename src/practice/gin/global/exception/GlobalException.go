package exception

import "strconv"

type GlobalException struct {
	Code int
	Msg  string
}

func (ex GlobalException) Error() string {
	return "에러코드: " + strconv.Itoa(ex.Code) + "\n에러메시지: " + ex.Msg
}

func InitGlobalException(code int, msg string) GlobalException {
	exception := GlobalException{}
	exception.Code = code
	exception.Msg = msg
	return exception
}

func (ex GlobalException) ToErrorResponse() ErrorResponse {
	errorResponse := ErrorResponse{}
	errorResponse.Code = ex.Code
	errorResponse.Msg = ex.Msg
	return errorResponse
}

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
