package charts

type IChart interface {
	Type() string
	ValidateOpts()
	GetAssets() (jsAssets, cssAssets []string)
}
