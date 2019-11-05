package model

import (
	"testing"
)

func init() {
	config.InitConfig()
	config.InitLogger(config.AppName)
}
func TestGetCollectCount(t *testing.T) {
	total, err := GetCollectCount("purchase", "1233456")
	t.Logf("%+v %+v", total, err)
}

func TestGetStockByDateYrdUid(t *testing.T) {
	c, err := GetStockByDateYrdUid("20190604", "123456", "purchase", 1, 10)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(c); i++ {
		t.Logf("%+v\n", c[i])
	}
}

func TestUpdateCollectByYrdUidStockId(t *testing.T) {
	data := new(AnalystCollect)
	data.Collection = "no"
	affected, err := UpdateCollectByYrdUidStockId("123456", 1, data)
	t.Logf("%+v %+v", affected, err)
}

func TestGetCollectStockByDateYrdUid(t *testing.T) {
	affected, err := GetCollectStockByDateYrdUid("recommend", "1000060150", 1, 10)
	t.Logf("%+v %+v", affected, err)
}
func TestGetStockByDate(t *testing.T) {
	t.Log(config.GetGlobalConfig().Mysql)
	stock, err := GetStockByDate("20190524", "purchase")
	t.Logf("%+v %+v", stock, err)
}
