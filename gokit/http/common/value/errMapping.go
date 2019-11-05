package value

// 错误源
var SourceOfErr = map[string]string{
	"gateway":       "10",
	"micro-service": "11",
	"data-service":  "12",
}

// 参数标识
var Param = map[string]string{
	"None":          "0000",
	"Cid1":          "0001",
	"Cid2":          "0002",
	"Period":        "0003",
	"Month":         "0004",
	"Dt":            "0005",
	"YrdUid":        "0006",
	"Company":       "0007",
	"InsuranceType": "0008",
	"Request":       "5000",
	"Combine":       "5001",
	"Response":      "5002",
	"Update":        "5003",
	"Delete":        "5004",
	"Reservation":   "5005",
	"Stock":         "5010",
}

// 错误定义
var TypeErr = map[string]string{
	"required":          "001", // 不为空
	"notExist":          "002", // 不存在
	"format":            "003", // 格式错误
	"illegal":           "004", // 错误（不合法）
	"notLess":           "005", // 不小于
	"fail":              "006", // 失败
	"serviceExceptions": "007", // 服务异常
	"notCondition":      "008", // 不符合条件
	"notSuccStatus":     "009", // 非成功态
	"len":               "010", // 长度不符合条件
	"conflict":          "030", // 冲突
}
