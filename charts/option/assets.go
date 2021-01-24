package option

import (
	"github.com/lioneagle/echarts/types"
)

type AssetsOpts struct {
	JSAssets  *types.StringSet
	CSSAssets *types.StringSet
}

func NewAssetsOpts() *AssetsOpts {
	return &AssetsOpts{
		JSAssets:  types.NewStringSet(),
		CSSAssets: types.NewStringSet(),
	}
}
