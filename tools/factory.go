package tools

var (
	toolsArr []ITool
)

// 向工厂注册tool
func registerTool(tool ITool) {
	toolsArr = append(toolsArr, tool)
}

// GetTools 获取所有注册的tool
func GetTools() []ITool {
	return toolsArr
}
