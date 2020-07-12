package option

// BaseOpts represents a option set needed by all chart types.
type BaseOpts struct {
	InitOpts   // 图形初始化配置项
	AssetsOpts // 静态资源配置项
	TitleOpts  // 标题组件配置项

}
