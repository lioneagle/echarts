package option

type TooltipAxisPointerOpts struct {
	Type string `json:"type,omitempty`
}

// TooltipOpts is the option set for a tooltip component.
type TooltipOpts struct {
	// 是否显示提示框
	Show bool `json:"show,omitempty"`
	// 触发类型。
	// "item": 数据项图形触发，主要在散点图，饼图等无类目轴的图表中使用。
	// "axis": 坐标轴触发，主要在柱状图，折线图等会使用类目轴的图表中使用。
	// "none": 什么都不触发。
	Trigger string `json:"trigger,omitempty"`
	// 提示框触发的条件，可选：
	// "mousemove": 鼠标移动时触发。
	// "click": 鼠标点击时触发。
	// "mousemove|click": 同时鼠标移动和点击时触发。
	// "none": 不在 "mousemove" 或 "click" 时触发
	TriggerOn string `json:"triggerOn,omitempty"`
	// 1, 字符串模板
	// 模板变量有 {a}, {b}，{c}，{d}，{e}，分别表示系列名，数据名，数据值等。
	// 在 trigger 为 'axis' 的时候，会有多个系列的数据，此时可以通过 {a0}, {a1}, {a2}
	// 这种后面加索引的方式表示系列的索引。 不同图表类型下的 {a}，{b}，{c}，{d} 含义不一样。
	// 其中变量{a}, {b}, {c}, {d} 在不同图表类型下代表数据含义为：
	// 折线（区域）图、柱状（条形）图、K 线图 : {a}（系列名称），{b}（类目值），{c}（数值）, {d}（无）
	// 散点图（气泡）图 : {a}（系列名称），{b}（数据名称），{c}（数值数组）, {d}（无）
	// 地图 : {a}（系列名称），{b}（区域名称），{c}（合并数值）, {d}（无）
	// 饼图、仪表盘、漏斗图: {a}（系列名称），{b}（数据项名称），{c}（数值）, {d}（百分比）
	//
	// 2, 回调函数
	// 回调函数格式：
	// (params: Object|Array, ticket: string, callback: (ticket: string, html: string)) => string
	// 第一个参数 params 是 formatter 需要的数据集。格式如下：
	// {
	//    componentType: 'series',
	//    seriesType: string,	// 系列类型
	//    seriesIndex: number,	// 系列在传入的 option.series 中的 index
	//    seriesName: string,	// 系列名称
	//    name: string,			// 数据名，类目名
	//    dataIndex: number,	// 数据在传入的 data 数组中的 index
	//    data: Object,			// 传入的原始数据项
	//    value: number|Array,	// 传入的数据值
	//    color: string,		// 数据图形的颜色
	//    percent: number,		// 饼图的百分比
	// }
	Formatter string `json:"formatter,omitempty"`

	AxisPointer TooltipAxisPointerOpts `json:"axisPointer,omitempty"`
}

func (*TooltipOpts) markGlobal() {}

func NewTooltipOpts() *TooltipOpts {
	return &TooltipOpts{}
}
