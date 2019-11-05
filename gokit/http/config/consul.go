package config

import (
	"bytes"
	"encoding/json"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/watch"
	"github.com/penglq/QLog"
	"log"
	"sync"
)

var once = new(sync.Once)

func InitConsul() {
	once.Do(func() {
		consulPlan(consulKey, watchDataserviceConfig)
		consulPlan(consulQuestionKey, watchQuestionConfig)
		// 先获取一次
		KVGetConfig(consulKey, parseConsulConfig)
		KVGetConfig(consulQuestionKey, parseConsulQuestionConfig)
	})
}

func consulPlan(key string, f func(idx uint64, result interface{})) {
	plan := mustParse(`{"type":"key","key":"` + key + `"}`)
	// 环境变量只有在容器里有，如果没有获取到手动赋值 方便本地开发调试
	plan.Token = consulToken
	plan.Handler = f

	go func() {
		log.Println("consul error", plan.Run(consulUrl))
	}()
}

func watchDataserviceConfig(idx uint64, result interface{}) {
	if entries, ok := result.(*api.KVPair); ok {
		// log.Println("index", idx, "value", string(entries.Value))
		parseConsulConfig(entries.Value)
	} else {
		log.Println("watchDataserviceConfig error", "类型不为*api.KVPair")
	}
}

func watchQuestionConfig(idx uint64, result interface{}) {
	if entries, ok := result.(*api.KVPair); ok {
		// log.Println("index", idx, "value", string(entries.Value))
		parseConsulQuestionConfig(entries.Value)
	} else {
		log.Println("watchQuestionConfig error", "类型不为*api.KVPair")
	}
}

func KVGetConfig(key string, f func([]byte)) {
	// 先获取一次
	conf := api.DefaultConfig()
	conf.Address = consulUrl
	conf.Token = consulToken
	c, err := api.NewClient(conf)
	if err != nil {
		panic("连接consul出错")
	}
	// fmt.Println(">>>>consulkey", key, ">>>>consulUrl", consulUrl, ">>>>consulToken", consulToken)
	kv, _, err := c.KV().Get(key, nil)
	if err != nil {
		panic("consul获取KV失败-" + err.Error())
	}
	// fmt.Println(string(kv.Value))
	f(kv.Value)
}

func parseConsulConfig(value []byte) {
	co := appConfig{}
	// 解析json
	err := json.Unmarshal(value, &co)
	if err != nil {
		QLog.GetLogger().Alert("parseConsulConfig配置文件json解析错误" + err.Error())
		return
	}
	// fmt.Println(">>>>", co)
	setGlobalConfig(&co)
}
func parseConsulQuestionConfig(value []byte) {
	co := GetGlobalConfig()
	question := make([]Questions, 0, 0)
	// 解析json
	err := json.Unmarshal(value, &question)
	if err != nil {
		QLog.GetLogger().Alert("parseConsulQuestionConfig配置文件json解析错误" + err.Error())
		return
	}
	co.Questions = question
	// fmt.Println(">>>", co)
	setGlobalConfig(&co)
}

func mustParse(q string) *watch.Plan {
	params := makeParams(q)
	plan, err := watch.Parse(params)
	if err != nil {
		log.Fatalf("plan err: %v", err)
	}
	return plan
}

func makeParams(s string) map[string]interface{} {
	var out map[string]interface{}
	dec := json.NewDecoder(bytes.NewReader([]byte(s)))
	if err := dec.Decode(&out); err != nil {
		log.Fatalf("err: %v", err)
	}
	return out
}
