package tools

import (
	"auto-tools/config"
	"auto-tools/logger"
	"context"
	"time"
)

func init() {
	tool := new(QTCrmTool)
	registerTool(tool)
}

type QTCrmTool struct {
	BaseTool
}

func (t *QTCrmTool) GetName() string {
	return "qt-crm-tool"
}

func (t *QTCrmTool) Run() {
	browserPage := t.getBrowserPage(true)
	if browserPage == nil {
		return
	}
	//设置下面发起的所有操作10秒超时，防止长时间等待卡住
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	pageWithCancel := browserPage.Context(ctx)

	crmUrl := config.GetCRMUrl()
	err := pageWithCancel.Navigate(crmUrl)
	if err != nil {
		logger.Errorf("qt-crm-tool run err:%s", err)
		return
	}

	err = pageWithCancel.WaitLoad()
	if err != nil {
		logger.Errorf("qt-crm-tool run err:%s", err)
		return
	}

	pageWithCancel.MustElement("a.icon.issueaction-create-subtask.jira-icon-add").MustClick()

	pageWithCancel.MustElement("#customfield_11607").MustSelect("研发类")
	//
	pageWithCancel.MustElement("div[data-customfieldid=\"customfield_11045\"]").MustClick()
	pageWithCancel.MustElement("#react-select-6-option-2").MustClick()
	//
	pageWithCancel.MustElement("div[data-customfieldid=\"customfield_11046\"]").MustClick()
	pageWithCancel.MustElement("#react-select-8-option-0").MustClick()
	//
	pageWithCancel.MustElement("div[data-customfieldid=\"customfield_11100\"]").MustClick()
	pageWithCancel.MustElement("#react-select-10-option-9").MustClick()
	//e := pageWithCancel.MustElement("div.sr-rs__menu.css-9g7j0m-menu")
	//fmt.Print(e.HTML())

	date := time.Now().Format("2006-01-02")
	pageWithCancel.MustElement("#customfield_11108").MustInput(date)
	pageWithCancel.MustElement("#customfield_11109").MustInput("8")
	pageWithCancel.MustElement("#customfield_11720").MustInput("风险app代码开发")
	// 点击提交按钮
	pageWithCancel.MustElement("#create-issue-submit").MustClick()
}
