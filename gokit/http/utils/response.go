package utils

import (
	"encoding/json"
	"git/miniTools/data-service/common/interfaces"
	"git/miniTools/data-service/common/value"
	"github.com/penglq/QLog"
	"net/http"
)

type OutResponse struct{}

// 成功响应时使用
func (r *OutResponse) ResponseSuccService(data interface{}) interfaces.StandardReturn {
	code, msg := GetErrCode(value.SourceSuccess, "", "")
	return responseReturn(code, msg, data)
}

// 错误响应
func (r *OutResponse) ResponseErrorService(source, field, errType string) interfaces.StandardReturn {
	code, msg := GetErrCode(source, field, errType)
	return responseReturn(code, msg, struct{}{})
}

// 请求后端服务失败
func (r *OutResponse) ResponseRequestErrorService(source string) interfaces.StandardReturn {
	code, msg := GetErrCode(source, value.ParamRequest, value.TypeFail)
	return responseReturn(code, msg, struct{}{})
}

// 用在全局中间件中的响应
func (r *OutResponse) WriteResponse(w http.ResponseWriter, req *http.Request, source, module, field, errType string) {
	res := r.ResponseErrorService(source, field, errType)
	b, _ := json.Marshal(&res)
	_, err := w.Write(b)
	if err != nil {
		QLog.GetLogger().Alert(TraceKey, GetTraceIdFromCTX(req.Context()), "method", "WriteResponse", "action", "http写响应", "error", err)
	}
}

// 响应
func responseReturn(code, msg string, data interface{}) interfaces.StandardReturn {
	return interfaces.StandardReturn{
		ResCode: code, ResDesc: msg, ResData: data,
	}
}
