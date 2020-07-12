package option

type IGlobalOption interface {
	MarkGlobal()
}

type ISeriesOption interface {
	MarkSeries()
}
