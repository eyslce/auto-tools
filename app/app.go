package app

import (
	"auto-tools/config"
	"auto-tools/logger"
	"auto-tools/tools"
	"flag"
	"github.com/robfig/cron/v3"
)

// Start 启动应用
func Start() {
	//初始化应用
	initApp()
	defer logger.Sync()
	//
	run()
}

func run() {
	logger.Info("app started")
	c := cron.New()
	toolsArr := tools.GetTools()
	for _, tool := range toolsArr {
		name := tool.GetName()

		if !config.IsToolActive(name) {
			logger.Warnf("%s is not active", name)
			continue
		}

		schedule := config.GetToolSchedule(name)
		_, err := c.AddJob(schedule, tool)
		if err != nil {
			logger.Errorf("add %s tool err %s", name, err)
			continue
		}
	}

	c.Start()

	select {}

}

func initApp() {
	//初始化配置
	configFile := flag.String("c", "config.json", "config file path")
	flag.Parse()
	config.InitConfig(*configFile)
	//初始化日志
	logfile := config.GetLogFile()
	logger.InitLogger(logfile)
}
