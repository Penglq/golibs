package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"git/miniTools/data-service/config"
	"github.com/penglq/QLog"
	"sync"
	"time"
)

var engine *xorm.Engine
var engineFund *xorm.Engine
var engineIndexValuation *xorm.Engine
var onceDo = new(sync.Once)
var onceDoFund = new(sync.Once)
var onceDoIndexValuation = new(sync.Once)

func InitEngine() {
	onceDo.Do(func() {
		var err error
		if dsn := config.GetGlobalConfig().Mysql.DSN; dsn != "" {
			engine, err = xorm.NewEngine("mysql", dsn)
			if err != nil {
				QLog.GetLogger().Alert("错误", "数据库连接错误")
				panic("数据库连接错误")
			}
			engine.SetMaxIdleConns(3)
			engine.SetMaxOpenConns(20)
			engine.SetConnMaxLifetime(0)
			engine.ShowExecTime(true)
			engine.ShowSQL(true)
			ping(0)
			go cyclePing()
		} else {
			QLog.GetLogger().Alert("配置", "数据库配置错误")
			panic("数据库配置错误")
		}
	})
}
func InitEngineFund() {
	onceDoFund.Do(func() {
		var err error
		if dsn := config.GetGlobalConfig().MysqlFund.DSN; dsn != "" {
			engineFund, err = xorm.NewEngine("mysql", dsn)
			if err != nil {
				QLog.GetLogger().Alert("错误", "数据库连接错误")
				panic("数据库连接错误")
			}
			engineFund.SetMaxIdleConns(3)
			engineFund.SetMaxOpenConns(20)
			engineFund.SetConnMaxLifetime(0)
			engineFund.ShowExecTime(true)
			engineFund.ShowSQL(true)
			pingFund(0)
			go cyclePingFund()
		} else {
			QLog.GetLogger().Alert("配置", "数据库配置错误")
			panic("数据库配置错误")
		}
	})
}
func InitEngineIndexValuation() {
	onceDoIndexValuation.Do(func() {
		var err error
		if dsn := config.GetGlobalConfig().MysqlIndexValuation.DSN; dsn != "" {
			engineIndexValuation, err = xorm.NewEngine("mysql", dsn)
			if err != nil {
				QLog.GetLogger().Alert("错误", "IndexValuation数据库连接错误")
				panic("IndexValuation数据库连接错误")
			}
			engineIndexValuation.SetMaxIdleConns(3)
			engineIndexValuation.SetMaxOpenConns(20)
			engineIndexValuation.SetConnMaxLifetime(0)
			engineIndexValuation.ShowExecTime(true)
			engineIndexValuation.ShowSQL(true)
			pingIndexValuation(0)
			go cyclePingIndexValuation()
		} else {
			QLog.GetLogger().Alert("配置", "IndexValuation数据库配置错误")
			panic("IndexValuation数据库配置错误")
		}
	})
}

func GetEngine() *xorm.Engine {
	InitEngine()
	return engine
}

func GetEngineFund() *xorm.Engine {
	InitEngineFund()
	return engineFund
}

func GetEngineIndexValuation() *xorm.Engine {
	InitEngineIndexValuation()
	return engineIndexValuation
}

func cyclePing() {
	for {
		time.Sleep(10 * time.Minute)
		ping(0)
	}
}
func cyclePingFund() {
	for {
		time.Sleep(10 * time.Minute)
		pingFund(0)
	}
}
func cyclePingIndexValuation() {
	for {
		time.Sleep(10 * time.Minute)
		pingIndexValuation(0)
	}
}

func ping(tryCount int8) {
	if tryCount > 3 {
		return
	}
	if err := engine.Ping(); err != nil {
		QLog.GetLogger().Alert("小工具数据库ping错误", err)
		time.Sleep(time.Second * 3)
		ping(tryCount + 1)
	}
}
func pingFund(tryCount int8) {
	if tryCount > 3 {
		return
	}
	if err := engineFund.Ping(); err != nil {
		QLog.GetLogger().Alert("万德数据库ping错误", err)
		time.Sleep(time.Second * 3)
		ping(tryCount + 1)
	}
}
func pingIndexValuation(tryCount int8) {
	if tryCount > 3 {
		return
	}
	if err := engineIndexValuation.Ping(); err != nil {
		QLog.GetLogger().Alert("ping错误", err)
		time.Sleep(time.Second * 3)
		ping(tryCount + 1)
	}
}
