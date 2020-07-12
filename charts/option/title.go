package option

// TitleOpts is the option set for a title component.
type TitleOpts struct {
	// 主标题
	Title string `json:"text,omitempty"`
	// 主标题样式配置项
	TitleStyle TextStyleOpts `json:"textStyle,omitempty"`
	// 副标题
	Subtitle string `json:"subtext,omitempty"`
	// 副标题样式配置项
	SubtitleStyle TextStyleOpts `json:"subtextStyle,omitempty"`
	// 主标题文本超链接
	Link string `json:"link,omitempty"`
	// 指定窗口打开主标题超链接
	// 'self' 当前窗口打开
	// 'blank' 新窗口打开
	Target string `json:"target,omitempty"`
	// grid 组件离容器上侧的距离。
	// top 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比
	// 也可以是 'top', 'middle', 'bottom'。
	// 如果 top 的值为'top', 'middle', 'bottom'，组件会根据相应的位置自动对齐
	Top string `json:"top,omitempty"`
	// 图例组件离容器下侧的距离。
	// bottom 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Bottom string `json:"bottom,omitempty"`
	// 图例组件离容器左侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比
	// 也可以是 'left', 'center', 'right'。
	// 如果 left 的值为'left', 'center', 'right'，组件会根据相应的位置自动对齐。
	Left string `json:"left,omitempty"`
	// 图例组件离容器右侧的距离。
	// right 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Right string `json:"right,omitempty"`
}

func (TitleOpts) MarkGlobal() {}
