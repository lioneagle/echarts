package option

type AxisPointerLinkOpts struct {
}

// AxisPointerOpts is the option set for a axisPointer component.
type AxisPointerOpts struct {
	// 是否显示
	Show bool `json:"show,omitempty"`
	// 联动轴
	Links []AxisPointerLinkOpts `json:"link,omitempty"`
}

func (*AxisPointerOpts) MarkGlobal() {}

func NewAxisPointerOpts() *AxisPointerOpts {
	return &AxisPointerOpts{}
}
