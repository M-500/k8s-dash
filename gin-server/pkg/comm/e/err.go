package e

var MessageFlag = map[int]string{
	SUCCESS:          "ok",
	ERROR:            "fail",
	BAD_INPUT_PARAMS: "请求参数错误",
	NOT_LOGIN:        "用户未登录",
	TokenExpired:     "Token已经过期，请重新登录",
	TokenNotValidYet: "Token not active yet",
	TokenMalformed:   "Token格式错误！",
	TokenInvalid:     "Couldn't handle this token:",
}

func GetMessage(code int) string {
	msg, ok := MessageFlag[code]
	if ok {
		return msg
	}
	return MessageFlag[ERROR]
}
