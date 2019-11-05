package value

var ErrMsg = map[string]string{
	"0":  "成功",
	"-1": "失败",
}

var ParamsErrMsg = map[string]string{
	"0000": "未知",
	"0001": "cid1",
	"0002": "cid2",
	"0003": "期限",
	"0004": "月份",
	"0005": "日期",
	"0006": "用户id",
	"0007": "公司",
	"0008": "保单类型",
	"5000": "请求后端",
	"5001": "组合数据",
	"5002": "后端响应",
	"5003": "更新数据",
	"5004": "删除数据",
	"5005": "预约时间",
	"5010": "股票数据",
}

var TypeErrMsg = map[string]string{
	"001": "不为空",
	"002": "不存在",
	"003": "格式错误",
	"004": "错误（不合法）",
	"005": "不小于",
	"006": "失败",
	"007": "服务异常",
	"008": "不符合条件",
	"009": "非成功态",
	"010": "长度不符合条件",
	"030": "冲突",
}
