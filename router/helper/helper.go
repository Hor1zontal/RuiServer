package helper

import (
	"RuiServer/exception"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

//type ErrorCode int32

//const (
//	ErrorCodeNone ErrorCode = iota
//	WxLogin     	= 10001	//微信登录出错
//	EroorCodeWxData         = 10002
//)

func ResponseWithCode(c *gin.Context, code exception.ErrorCode) {
	c.JSON(http.StatusOK, gin.H{"errcode": code, "errmsg": exception.ErrorMapping[code]})

}

func ResponseWithData(c *gin.Context, obj interface{}) {
	res := structToMap(obj)
	res["errcode"] = exception.CodeSuccess
	res["errmsg"] = exception.ErrorMapping[exception.CodeSuccess]
	c.JSON(http.StatusOK, res)
}

func structToMap(obj interface{}) gin.H {
	obj_v := reflect.ValueOf(obj)
	v := obj_v.Elem()
	typeOfType := v.Type()
	var data = make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		data[typeOfType.Field(i).Tag.Get("json")] = field.Interface()
	}
	return data
}

func CheckReq(c *gin.Context, req interface{}) {
	if err := c.ShouldBind(req); err != nil {
		exception.ExceptionCustom("CheckReq", exception.InvalidParam, err)
	}
}

func FormatReq(c *gin.Context, req interface{}) {
	if err := c.BindJSON(req); err != nil {
		exception.ExceptionCustom("FormatReq", exception.InvalidParam, err)
	}
}
