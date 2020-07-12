package option

// TextStyle is the option set for a text style component.
type TextStyleOpts struct {
	// 文字字体颜色
	Color string `json:"color,omitempty"`
	// 文字字体的风格
	// 可选  'normal', 'italic', 'oblique'
	FontStyle string `json:"fontStyle,omitempty"`
	// 字体大小
	FontSize int `json:"fontSize,omitempty"`
	// 递归结构，为了兼容 wordCloud
	Normal *TextStyleOpts `json:"normal,omitempty"`
}
