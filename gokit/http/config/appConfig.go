package config

import (
	"fmt"
	"os"
	"sync"
)

var appConfigs *appConfig
var mux = new(sync.RWMutex)
var consulUrl = ""
var consulKey = ""
var consulQuestionKey = ""
var consulToken = ""
var Company map[string]string = map[string]string{}
var InsuranceType map[string]string = map[string]string{}

func InitConfig() {
	appConfigs = &appConfig{}
	// 用于本地调试
	consulToken = os.Getenv("POD_CONSUL_TOKEN")
	if consulToken == "" {
		setGlobalConfig(&config_local)
		return
	}
	consulKey = os.Getenv("POD_CONSUL_KEY")
	if consulKey == "" {
		consulKey = "operations.miniTools/data-serviceDev"
	}
	consulUrl = os.Getenv("POD_CONSUL_URL")
	if consulUrl == "" {
		consulUrl = "http://consul.yixinonline.org:8500"
		// addr = "http://consul.default:8500" // 线上
	}
	consulQuestionKey = os.Getenv("POD_CONSUL_QUESTIONS_KEY")
	if consulQuestionKey == "" {
		consulQuestionKey = "operations.miniTools/questions"
	}

	InitConsul()
}

func GetGlobalConfig() appConfig {
	mux.RLock()
	defer mux.RUnlock()
	return *appConfigs
}

func setGlobalConfig(c *appConfig) {
	mux.Lock()
	defer mux.Unlock()
	appConfigs = c
	for _, v := range appConfigs.Company {
		Company[v] = ""
	}
	for _, v := range appConfigs.InsuranceType {
		InsuranceType[v] = ""
	}
	fmt.Printf("配置环境:%+v\n", (*appConfigs).Env)
}

type appConfig struct {
	Env   string `json:"env"`
	Mysql struct {
		DSN string `json:"dsn"`
	} `json:"mysql"`
	MysqlIndexValuation struct {
		DSN string `json:"dsn"`
	} `json:"mysql_index_valuation"`
	MysqlFund struct {
		DSN string `json:"dsn"`
	} `json:"mysql_fund"`
	AnalystUrl AnalystUrl `json:"analyst_url"`
	Redis      struct {
		Cluster  bool     `json:"cluster"`
		Addr     []string `json:"addr"`
		Password string   `json:"password"`
	} `json:"redis"`
	Rabbitmq struct {
		Addrs []string  `json:"addrs"`
		Queue QueueName `json:"queue"`
	} `json:"rabbitmq"`
	Alert                   Alert       `json:"alert"`
	Email                   Email       `json:"email"`
	ConsultantNotifyEmail   []string    `json:"consultantNotifyEmail"`
	ConsultantNotifyCCEmail []string    `json:"consultantNotifyCCEmail"`
	Company                 []string    `json:"company"`
	InsuranceType           []string    `json:"insurance_type"`
	Questions               []Questions `json:"questions"`
	UnTradingDay            []string    `json:"un_trading_day"`
	Hei                     float64     `json:"hei"`
}

type Email struct {
	ApiUrl   string `json:"apiUrl"`
	OrgNo    string `json:"orgNo"`
	AuthCode string `json:"authCode"`
	TplId    string `json:"tplId"`
}
type AnalystUrl struct {
	StockCompany       string `json:"stock_company"`
	StockIndustry      string `json:"stock_industry"`
	StockIndustryKLine string `json:"stock_industry_kline"`
	StockKLine         string `json:"stock_kline"`
	FundAppBaseURL     string `json:"fund_app_base_url"`
}
type Alert struct {
	AppId string `json:"appId"`
	URL   string `json:"url"`
}

type URL struct {
	Ipo             string `json:"ipo"`
	IpoDetail       string `json:"ipoDetail"`
	IndexEvaluation string `json:"indexEvaluation"`
	ExchangeRate    string `json:"exchangeRate"`
}

type QueueName struct {
	Withdraw struct {
		Exchange   string `json:"exchange"`
		RoutingKey string `json:"routing_key"`
	} `json:"withdraw"`
}

type Questions struct {
	QuestionId string            `json:"questionId"`
	Question   string            `json:"question"`
	Answers    map[string]string `json:"answers"`
	Weight     int               `json:"weight"`
	Attr       int               `json:"attr"`
}
