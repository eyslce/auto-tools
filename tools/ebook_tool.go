package tools

import (
	"context"
	"time"
)

func init() {

}

type EbookTool struct {
	BaseTool
}

func (t *EbookTool) GetName() string {
	return "ebook-tool"
}

func (t *EbookTool) Run() {
	browserPage := t.getBrowserPage(true)
	if browserPage == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	pg := browserPage.Context(ctx)
	doc := t.getDocument(pg, "https://m.31xs.com/9/9183/7336306.html")
	if doc == nil {
		return
	}

}
