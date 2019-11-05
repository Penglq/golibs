package utils

import (
	"git/miniTools/data-service/common/value"
)

func GetErrCode(source, field, errType string) (code, msg string) {
	if source == value.SourceSuccess || source == value.SourceFailed {
		return source, value.ErrMsg[source]
	}
	return GetErrCodeKey(source, field, errType)
}

func GetErrCodeKey(source, field, errType string) (code string, msg string) {
	var ok bool
	var tmp string
	var msgTmp string
	// 错误源
	if tmp, ok = value.SourceOfErr[source]; !ok {
		tmp = value.SourceOfErr[value.SourceDataservice]
	}
	code += tmp

	// 错误字段
	if tmp, ok = value.Param[field]; !ok {
		tmp = value.Param[value.ParamNone]
	}
	code += tmp
	if msgTmp, ok = value.ParamsErrMsg[field]; !ok {
		msgTmp = value.ParamsErrMsg[tmp]
	}
	msg += msgTmp

	// 错误类型
	if tmp, ok = value.TypeErr[errType]; !ok {
		tmp = value.TypeErr[value.TypeIllegal]
	}
	code += tmp
	if msgTmp, ok = value.TypeErrMsg[errType]; !ok {
		msgTmp = value.TypeErrMsg[tmp]
	}
	msg += msgTmp

	return
}
