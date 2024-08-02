package tools

import (
	"auto-tools/logger"
	"auto-tools/utils"
	"context"
	"fmt"
	"github.com/go-rod/rod/lib/proto"
	"time"
)

func init() {
	tool := new(BingTools)
	registerTool(tool)
}

type BingTools struct {
	BaseTool
}

func (b *BingTools) GetName() string {
	return "bing-tool"
}

func (b *BingTools) Run() {

}

func (b *BingTools) RunE(useMobile bool) {
	browser := b.getBrowser(true)
	if browser == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	browserPage, err := browser.Page(proto.TargetCreateTarget{URL: "about:blank"})
	if err != nil {
		logger.Errorf("bing-tool run err:%s", err)
		return
	}
	pg := browserPage.Context(ctx)
	if useMobile {
		b.EmulateDevice(pg)
	}
	baseurl := fmt.Sprintf("https://cn.bing.com/search?q=%s", utils.RandomHanZi(2))
	err = pg.Navigate(baseurl)
	if err != nil {
		logger.Errorf("bing-tool run err:%s", err)
		return
	}
	err = pg.WaitLoad()
	if err != nil {
		logger.Errorf("bing-tool run err:%s", err)
		return
	}
	time.Sleep(1 * time.Second)
	if useMobile {
		for i := 0; i < 20; i++ {
			word := utils.RandomHanZi(2)
			pg.MustElement("#sb_form_q").MustSelectAllText().MustInput("")
			pg.MustElement("#sb_form_q").MustInput(word)
			pg.MustElement("#sb_form_go").MustClick()
			time.Sleep(5 * time.Second)
		}
	} else {
		for i := 0; i < 30; i++ {
			word := utils.RandomHanZi(2)
			pg.MustElement("#sb_form_q").MustSelectAllText().MustInput("")
			pg.MustElement("#sb_form_q").MustInput(word)
			pg.MustElement("#sb_search").MustClick()
			time.Sleep(5 * time.Second)
		}
	}
	err = pg.Close()
	if err != nil {
		logger.Errorf("bing-tool run err:%s", err)
		return
	}
}
