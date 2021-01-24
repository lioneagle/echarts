package option

// InitOpts contains options for the canvas.
type InitOpts struct {
	// 生成的 HTML 页面标题
	PageTitle string
	// 画布宽度
	Width string
	// 画布高度
	Height string
	// 画布背景颜色
	BackgroundColor string `json:"backgroundColor,omitempty"`
	// 图表 ID，是图表唯一标识
	ChartID string
	// 静态资源 host 地址
	AssetsHost string
	// 图表主题
	Theme string
}

func (*InitOpts) MarkGlobal() {}

func NewInitOpts() *InitOpts {
	return &InitOpts{
		PageTitle:  "Awesome go-echarts",
		Width:      "900px",
		Height:     "500px",
		AssetsHost: "https://go-echarts.github.io/go-echarts-assets/assets/",
		Theme:      "white",
	}
}
