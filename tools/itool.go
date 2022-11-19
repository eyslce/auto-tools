package tools

// ITool Tool 工具接口
type ITool interface {

	// GetName 获取工具名
	GetName() string

	// Run 执行工具
	Run()
}
