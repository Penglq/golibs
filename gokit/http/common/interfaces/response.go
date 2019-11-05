package interfaces

// 标准返回结构
type StandardReturn struct {
	ResCode string      `json:"resCode"`
	ResDesc string      `json:"resDesc"`
	ResData interface{} `json:"resData"`
}
