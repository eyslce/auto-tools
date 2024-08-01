package tools

import (
	"auto-tools/config"
	"auto-tools/logger"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
	"github.com/samber/lo"
	"math/rand"
	"strings"
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
	l := launcher.NewUserMode().Headless(!debug)
	u, err := l.Bin(path).
		Leakless(true).Launch()
	if err != nil {
		logger.Errorf("launch browser err:%s", err)
		return nil
	}

	browser := rod.New().
		Trace(true).
		Logger(logger.GetLoggerFactory()).
		ControlURL(u)
	browser.SlowMotion(time.Millisecond * 100)

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

func (b *BaseTool) EmulateDevice(pg *rod.Page) {
	devs := []devices.Device{
		devices.Pixel2, devices.Pixel2XL, devices.IPad, devices.IPadMini, devices.IPadPro,
		devices.IPhoneX, devices.IPhone4, devices.IPhone5orSE, devices.IPhone6or7or8,
		devices.IPhone6or7or8Plus, devices.Nexus6, devices.GalaxyNote3, devices.Nexus7,
		devices.Nexus6P, devices.Nexus10,
	}
	index := rand.Intn(15)
	err := pg.Emulate(devs[index])
	if err != nil {
		logger.Errorf(fmt.Sprintf("scan cnvd err:%s", err))
	}
}

func (b *BaseTool) getDocument(pg *rod.Page, url string, useMobile bool) *goquery.Document {
	if useMobile {
		b.EmulateDevice(pg)
	}
	err := pg.Navigate(url)
	if err != nil {
		logger.Errorf("nav to url:%s", err)
		return nil
	}
	err = pg.WaitLoad()
	if err != nil {
		logger.Errorf("wait page load err:%s", err)
		return nil
	}
	html, err := pg.HTML()
	if err != nil {
		logger.Errorf("get page html err:%s", err)
		return nil
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		logger.Errorf("get document err:%s", err)
		return nil
	}
	return doc
}
