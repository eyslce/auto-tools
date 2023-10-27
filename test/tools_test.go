package test

import (
	"auto-tools/config"
	"auto-tools/logger"
	"auto-tools/tools"
	"testing"
)

func initconfig() {
	config.InitConfig("../config.json")
	//初始化日志
	logfile := config.GetLogFile()
	logger.InitLogger(logfile)
}

func TestQtCrmTool(t *testing.T) {
	initconfig()
	tool := new(tools.QTCrmTool)
	tool.Run()
}
