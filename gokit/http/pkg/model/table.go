package model

type Count struct {
	S int64
}
type Collect struct {
	AnalystStocksdata `xorm:"extends"`
	AnalystCollect    `xorm:"extends"`
}
type AnalystCollect struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	YrdUid     string `xorm:"not null default '' comment('宜人贷用户ID') VARCHAR(20)"`
	StockId    int    `xorm:"not null default 0 comment('上市公司股票id') unique(yrd_uid_stock_id) INT(11)"`
	Collection string `xorm:"not null default '' comment('是否收藏过:yes/no') CHAR(5)"`
	CreatedAt  int    `xorm:"created not null INT(10)"`
	UpdatedAt  int    `xorm:"updated not null INT(10)"`
}

type AnalystStocksdata struct {
	Id            int     `xorm:"not null pk autoincr INT(11)"`
	SratingName   string  `xorm:"not null default '' comment('所属类型:强烈推荐(recommend)/买入(purchase)') unique(date_code_Industry_code_srating_name) VARCHAR(20)"`
	StockType     string  `xorm:"not null default '' comment('股票类型:个股(company)/板块(industry)') VARCHAR(10)"`
	Code          string  `xorm:"not null default '' comment('上市公司股票代码') unique(date_code_Industry_code_srating_name) VARCHAR(20)"`
	CodeName      string  `xorm:"not null default '' comment('上市公司股票名称') VARCHAR(20)"`
	IndustryName  string  `xorm:"not null default '' comment('行业名称') VARCHAR(20)"`
	IndustryCode  string  `xorm:"not null default '' comment('行业代码') unique(date_code_Industry_code_srating_name) VARCHAR(20)"`
	Actualdprice  float64 `xorm:"not null default 0.00 comment('目标股价') DECIMAL(18,2)"`
	Date          string  `xorm:"not null default '' comment('报告日期') unique(date_code_Industry_code_srating_name) CHAR(8)"`
	FileName      string  `xorm:"not null default '' comment('分析报告名称') VARCHAR(300)"`
	FileUrl       string  `xorm:"not null default '' comment('分析报告地址') VARCHAR(300)"`
	AuthorName    string  `xorm:"not null default '' comment('分析报告作者名') VARCHAR(20)"`
	AuthorInfo    string  `xorm:"not null comment('分析报告作者描述') TEXT"`
	PotentialGain int     `xorm:"not null default 0 comment('潜在涨幅/%') INT(10)"`
	FundName      string  `xorm:"not null default '' comment('基金名称') VARCHAR(20)"`
	FundCode      string  `xorm:"not null default '' comment('基金code') VARCHAR(20)"`
	CreatedAt     int     `xorm:"created not null default 0 INT(10)"`
	UpdatedAt     int     `xorm:"updated not null default 0 INT(10)"`
}

type AnalystUser struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	YrdUid     string `xorm:"not null default '' comment('宜人贷用户ID') VARCHAR(20)"`
	PassportId string `xorm:"not null default '' comment('宜人贷passportId') VARCHAR(60)"`
	Mobile     string `xorm:"not null default '' comment('手机号') VARCHAR(11)"`
	Name       string `xorm:"not null default '' comment('姓名') VARCHAR(100)"`
	CreatedAt  int    `xorm:"created not null default 0 INT(10)"`
	UpdatedAt  int    `xorm:"updated not null default 0 INT(10)"`
}