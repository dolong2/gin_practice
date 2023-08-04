package exception

import "strconv"

type UserAlreadyExistsException struct {
	Code int
	Msg  string
}

func (ex UserAlreadyExistsException) Error() string {
	return "에러코드: " + strconv.Itoa(ex.Code) + "\n에러메시지: " + ex.Msg
}

func InitUserAlreadyExistsException() UserNotFoundException {
	exception := UserNotFoundException{}
	exception.Code = 404
	exception.Msg = "해당 유저를 찾을 수 없음"
	return exception
}
