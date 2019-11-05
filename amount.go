package util

import (
	"fmt"
	"strconv"
)

func AmountWithChineseUnit(amount int64) string {
	wan := amount / 1000000
	if wan > 0 {
		return strconv.FormatInt(wan, 10) + "万"
	}
	qian := amount / 100000
	if qian > 0 {
		return strconv.FormatInt(qian, 10) + "千"
	}
	bai := amount / 10000
	return strconv.FormatInt(bai, 10) + "百"
}

func Abs(amount int64) int64 {
	if amount < 0 {
		return -amount
	}
	return amount
}

func NonNegative(amount int64) int64 {
	if amount < 0 {
		return 0
	}
	return amount
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 500 变为5.00, 1 变为0.01, 10 变为0.10
func AmountToString(value int64) string {
	if value < 0 {
		return "0.00"
	}
	amountString := strconv.FormatInt(value, 10)
	switch len(amountString) {
	case 1:
		amountString = "00" + amountString
	case 2:
		amountString = "0" + amountString
	}
	return amountString[:len(amountString)-2] + "." + amountString[len(amountString)-2:]
}

func AmountToSignString(value int64) string {
	sign := "+"
	newValue := value
	if value < 0 {
		sign = "-"
		newValue = -value
	} else if value == 0 {
		sign = ""
	}
	return sign + AmountToString(newValue)
}

// 500 变为 5
func AmountToYuanString(value int64) string {
	return strconv.FormatInt((value+50)/100, 10)
}

// 500 变为 5
func AmountToYuanStringNoCarry(value int64) string {
	return strconv.FormatInt(value/100, 10)
}

//1000000 变成 10,000.00 增加了逗号来分割
func AmountToFormatString(value int64) string {
	base := AmountToString(value)
	ret := base[len(base)-3:]
	left := base[:len(base)-3]
	for len(left) > 3 {
		ret = "," + left[len(left)-3:] + ret
		left = left[:len(left)-3]
	}
	ret = left + ret
	return ret
}

// 小于一万元，显示xxx元，大于等于1万元时，则显示x.x万元（向下截取），如果小数点后是0，则显示x万元
// 5000 -> 50元，156000 -> 1.5万元
func AmountToShortString(value int64) string {
	amount := value / 100
	var amountStr string
	if amount >= 10000 {
		amount = amount / 1000 * 1000
		amountStr = fmt.Sprintf("%.1f", float64(amount)/10000.0)
		if amountStr[len(amountStr)-1:len(amountStr)] == "0" {
			amountStr = amountStr[0:len(amountStr)-2] + "万"
		} else {
			amountStr = amountStr + "万"
		}
	} else {
		amountStr = fmt.Sprintf("%d", amount)
	}

	return amountStr
}

func AmountCeilToYuan(value int64) int64 {
	if value <= 0 || value%100 == 0 {
		return value
	}
	return 100 + value - value%100
}

//0->0,100->1,10->0.1,1->0.01
func AmountToSimpleString(value int64) string {
	if value == 0 {
		return "0"
	}
	if value%100 == 0 {
		return fmt.Sprintf("%d", value/100)
	} else if value%10 == 0 {
		return fmt.Sprintf("%.1f", float64(value)/100.0)
	} else {
		return fmt.Sprintf("%.2f", float64(value)/100.0)
	}
}

//0->0,100->1,10->0.1,1->0.01,0.1->0 小数点后最多保留两位
func FloatAmountToSimpleString(value float64) string {
	return AmountToSimpleString(RoundInt64(value))
}

func BankPurchaseQuotaToString(mapp int64, mapd int64) string {
	var mappStr, mapdStr string
	if mapp >= 1000000 {
		mappStr = fmt.Sprintf("单笔限额%d万", mapp/1000000)
	} else {
		mappStr = fmt.Sprintf("单笔限额%d千", mapp/100000)
	}
	if mapd >= 1000000 {
		mapdStr = fmt.Sprintf("日限额%d万", mapd/1000000)
	} else {
		mapdStr = fmt.Sprintf("日限额%d千", mapd/100000)
	}
	if mapp == 0 {
		return ""
	} else if mapd == 0 {
		return mappStr
	}
	return fmt.Sprintf("%s，%s", mappStr, mapdStr)
}

// 忽略数据中的角和分
func AmountIgnoreJiaoAndFen(value int64) int64 {
	return value / 100 * 100
}
