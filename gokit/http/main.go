package main

import (
	"github.com/penglq/QLog"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	var err error
	config.InitConfig()
	config.InitLogger(config.AppName)
	err = http.ListenAndServe(":8081", transport.NewTransport())
	QLog.GetLogger().Fatal("http error", err)
}
