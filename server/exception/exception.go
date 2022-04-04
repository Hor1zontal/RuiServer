package exception

import (
	"log"
)

type ErrorCode int32

const (
	CodeSuccess    ErrorCode = iota
	InvalidParam             = 10001 //参数错误
	DatabaseError            = 10002 //数据库操作异常
	InternalError            = 10003 //内部异常
	TimeParseError           = 10004 //时间格式错误
)

var ErrorMapping = map[ErrorCode]string{
	CodeSuccess:    "success",
	InvalidParam:   "invalid param",
	DatabaseError:  "database error",
	InternalError:  "internal error",
	TimeParseError: "time parse error",
}

func ExceptionCustom(api string, code ErrorCode, err error) {
	log.Println(api + " --- code:" + string(code) + "--- error:" + err.Error())
	exception(code)
}

func exception(this ErrorCode) {
	panic(this)
}
