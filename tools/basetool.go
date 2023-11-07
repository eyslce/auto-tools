package tools

import (
	"auto-tools/config"
	"auto-tools/logger"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
	"github.com/samber/lo"
	"time"
)

type BaseTool struct {
}

func (b *BaseTool) getBrowserPage(debug bool) *rod.Page {
	path := config.GetBrowserPath()
	if lo.IsEmpty(path) {
		lookPath, _ := launcher.LookPath()
		path = lookPath
	}
	l := launcher.NewUserMode()
	if debug {
		l.Headless(false)
	}
	u, err := l.
		Bin(path).Launch()
	if err != nil {
		logger.Errorf("launch browser err:%s", err)
		return nil
	}

	browser := rod.New().
		Trace(true).
		Logger(logger.GetLoggerFactory()).
		ControlURL(u)
	if debug {
		browser.SlowMotion(time.Second)
	}

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
