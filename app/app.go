package app

import (
	"auto-tools/config"
	"auto-tools/logger"
	"auto-tools/tools"
	"github.com/robfig/cron/v3"
)

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
