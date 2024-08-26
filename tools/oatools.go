package tools

import (
	"auto-tools/config"
	"auto-tools/logger"
	"context"
	"strconv"
	"strings"
	"time"
)

func init() {
	tool := new(OaTools)
	registerTool(tool)
}

type OaTools struct {
	BaseTool
}

func (o *OaTools) GetName() string {
	return "oa-tool"
}

func (o *OaTools) Run() {
	browser := o.getBrowser(true)
	if browser == nil {
		return
	}
	//设置下面发起的所有操作10秒超时，防止长时间等待卡住
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	oaUrl := config.GetOAUrl()
	browserPage := browser.MustPage(oaUrl)
	pageWithCancel := browserPage.Context(ctx)

	err := pageWithCancel.WaitLoad()
	if err != nil {
		logger.Errorf("oa-tool run err:%s", err)
		return
	}

	pageWithCancel.MustElement("#addbutton1").MustClick()
	pageWithCancel.MustElement("div.ant-select-selection__rendered").MustClick()
	elements := pageWithCancel.MustElements("li.ant-select-dropdown-menu-item")
	for _, element := range elements {
		if element.MustText() == "研发类" {
			element.MustClick()
			break
		}
	}
	searchElements := pageWithCancel.MustElements("i.anticon-search")
	for i, element := range searchElements {
		if i == 3 {
			element.MustClick()
			time.Sleep(1 * time.Second)
			es := pageWithCancel.MustElements("input.ant-input")
			es[1].MustInput("cwpp")
			buttons := pageWithCancel.MustElements("button.ant-btn.ant-btn-ghost.ant-btn-icon-only.wea-input-focus-btn")
			buttons[1].MustClick()
			time.Sleep(1 * time.Second)
			firstTasks := pageWithCancel.MustElements("div.wea-url")
			for _, e := range firstTasks {
				if strings.Contains(e.MustText(), "CWPP") {
					e.MustClick()
					break
				}
			}
		}
		if i == 4 {
			element.MustClick()
			time.Sleep(1 * time.Second)
			firstTasks := pageWithCancel.MustElements("div.wea-url")
			for _, e := range firstTasks {
				if strings.Contains(e.MustText(), "资配漏补") {
					e.MustClick()
					break
				}
			}
		}
		if i == 5 {
			element.MustClick()
			time.Sleep(1 * time.Second)
			firstTasks := pageWithCancel.MustElements("div.wea-url")
			for _, e := range firstTasks {
				if strings.Contains(e.MustText(), "代码开发") {
					e.MustClick()
					break
				}
			}
		}
	}
	pageWithCancel.MustElement("#field18643_0").MustInput("风险app代码开发")
	pageWithCancel.MustElement("div.wea-date-picker").MustClick()
	t := time.Now()
	date := t.Format("2006-01-02")
	pageWithCancel.MustElement("input.ant-calendar-input").MustInput(date)
	dateElements := pageWithCancel.MustElements("div.ant-calendar-date")
	for _, element := range dateElements {
		attr := element.MustAttribute("aria-disabled")
		if attr == nil || *attr != "false" {
			continue
		}
		attr = element.MustAttribute("aria-selected")
		if attr == nil || *attr != "true" {
			continue
		}
		day := t.Day()
		if element.MustText() == strconv.Itoa(day) {
			element.MustClick()
			break
		}
	}
	pageWithCancel.MustElement("#field13396_0").MustInput("8")
	buttons := pageWithCancel.MustElements("div.wf-req-top-button")
	for _, button := range buttons {
		if button.MustText() == "提 交" {
			button.MustClick()
			break
		}
	}
	time.Sleep(5 * time.Second)
}
