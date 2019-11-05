package utils

import (
	"git/miniTools/data-service/common/value"
	"testing"
)

func TestGetErrCode(t *testing.T) {
	codeWish := []string{
		"100001001",
		"100000004",
	}

	code, msg := GetErrCode(value.SourceDataservice, "Mobile", "required")
	if code != codeWish[0] {
		t.Fatalf("输出值：%+v  期望值：%+v\n", code, codeWish[0])
	}
	t.Log(code, msg)

	code, msg = GetErrCode("", "", "Mobiles")
	if code != codeWish[1] {
		t.Fatalf("输出值：%+v  期望值：%+v\n", code, codeWish[1])
	}
	t.Log(code, msg)
}
