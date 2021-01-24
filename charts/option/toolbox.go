package option

// TBFeature is a feature component under toolbox.
type TBFeature struct {
	// 保存为图片
	SaveAsImage struct{} `json:"saveAsImage"`
	// 数据区域缩放。目前只支持直角坐标系的缩放
	DataZoom struct{} `json:"dataZoom"`
	// 数据视图工具，可以展现当前图表所用的数据，编辑后可以动态更新
	DataView struct{} `json:"dataView"`
	// 配置项还原
	Restore struct{} `json:"restore"`
}

// ToolboxOpts is the option set for a toolbox component.
type ToolboxOpts struct {
	// 是否显示工具栏组件
	Show bool `json:"show"`
	// 工具箱功能种类，不支持自定义
	TBFeature `json:"feature"`
}

func (*ToolboxOpts) MarkGlobal() {}

func NewToolboxOpts() *ToolboxOpts {
	return &ToolboxOpts{}
}
