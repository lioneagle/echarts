package option

// BaseOpts represents a option set needed by all chart types.
type BaseOpts struct {
	*InitOpts    // 图形初始化配置项
	*AssetsOpts  // 静态资源配置项
	*TitleOpts   // 标题组件配置项
	*LegendOpts  // 图例组件配置项
	*TooltipOpts // 提示框组件配置项
	*ToolboxOpts // 工具箱组件配置项
}

func NewBaseOpts() *BaseOpts {
	return &BaseOpts{
		InitOpts:    NewInitOpts(),
		AssetsOpts:  NewAssetsOpts(),
		TitleOpts:   NewTitleOpts(),
		LegendOpts:  NewLegendOpts(),
		TooltipOpts: NewTooltipOpts(),
		ToolboxOpts: NewToolboxOpts(),
	}
}
