package tools

import (
	"auto-tools/logger"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

type BaseTool struct {
}

func (b *BaseTool) getBrowserPage() *rod.Page {
	path, _ := launcher.LookPath()
	u, err := launcher.New().
		//Headless(false).
		Bin(path).Launch()
	if err != nil {
		logger.Errorf("launch browser err:%s", err)
		return nil
	}

	browser := rod.New().
		//SlowMotion(time.Second).
		Trace(true).
		Logger(logger.GetLoggerFactory()).
		ControlURL(u)

	err = browser.Connect()
	if err != nil {
		logger.Errorf("launch browser err:%s", err)
		return nil
	}

	//防止机器人检测
	page, err := stealth.Page(browser)
	if err != nil {
		logger.Errorf("launch browser err:%v", err)
		return nil
	}

	return page
}
