package consts

const SYSTEM_ERROR = 10001

var arrErrMsgMap = map[int]string{
	SYSTEM_ERROR: "系统错误",
}

func GetErrorMsg(code int) string {
	return arrErrMsgMap[code]
}
