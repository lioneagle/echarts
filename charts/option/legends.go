package option

// LegendOpts is the option set for a legend component.
type LegendOpts struct {
	// 图例组件离容器左侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比
	// 也可以是 'left', 'center', 'right'。
	// 如果 left 的值为'left', 'center', 'right'，组件会根据相应的位置自动对齐。
	Left string `json:"left,omitempty"`
	// 图例组件离容器上侧的距离。
	// top 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比
	// 也可以是 'top', 'middle', 'bottom'。
	// 如果 top 的值为'top', 'middle', 'bottom'，组件会根据相应的位置自动对齐。
	Top string `json:"top,omitempty"`
	// 图例组件离容器右侧的距离。
	// right 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Right string `json:"right,omitempty"`
	// 图例组件离容器下侧的距离。
	// bottom 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Bottom string `json:"bottom,omitempty"`
	// Legend 数据项
	// 如果需要隐藏 Legend 则把 Data 设置为 []string{}
	Data interface{} `json:"data,omitempty"`
	// 除此之外也可以设成 "single" 或者 "multiple" 使用单选或者多选模式。默认 "multiple"
	SelectedMode string `json:"selectedMode,omitempty"`
	// 图例的公用文本样式
	TextStyle TextStyleOpts `json:"textStyle,omitempty"`
}

func (*LegendOpts) markGlobal() {}

func NewLegendOpts() *LegendOpts {
	return &LegendOpts{}
}
