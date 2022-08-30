package golibs

import (
	"encoding/json"
	"testing"
)

func TestCeil(t *testing.T) {
	t.Log(Ceil(10, 3))
}

func TestGetMonthZero(t *testing.T) {
	t.Log(GetMonthZero())
}

func TestGetNextMonthZero(t *testing.T) {
	t.Log(GetNextMonthZero())
}

func TestJsonMap(t *testing.T) {
	b := make(map[string][]string)
	b["a"] = []string{"aa", "bb"}
	b["b"] = []string{"aa1", "bb1"}
	b["c"] = []string{"aa1", "bb1"}

	d := []data{}
	d = append(d, data{Questions: b})
	a := s{
		Rescode: "0",
		ResDesc: "成功",
		ResData: d,
	}

	c, _ := json.Marshal(&a)

	t.Log("json", string(c))
}

type s struct {
	Rescode string `json:"resCode"`
	ResDesc string `json:"resDesc"`
	ResData []data `json:"resData"`
}

type data struct {
	Name            string              `json:"name"`
	Mobile          string              `json:"mobile"`
	ReservationTime string              `json:"reservationTime"`
	Questions       map[string][]string `json:"questions"`
}
