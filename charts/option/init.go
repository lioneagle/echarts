package option

// InitOpts contains options for the canvas.
type InitOpts struct {
	// 生成的 HTML 页面标题
	PageTitle string `default:"Awesome go-echarts"`
	// 画布宽度
	Width string `default:"900px"`
	// 画布高度
	Height string `default:"500px"`
	// 画布背景颜色
	BackgroundColor string `json:"backgroundColor,omitempty"`
	// 图表 ID，是图表唯一标识
	ChartID string
	// 静态资源 host 地址
	AssetsHost string `default:"https://go-echarts.github.io/go-echarts-assets/assets/"`
	// 图表主题
	Theme string `default:"white"`
}

func (InitOpts) MarkGlobal() {}
