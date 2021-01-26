package techan

import "github.com/sdcoffey/big"

func NewTrueRangeIndicator(series *TimeSeries, window int) Indicator {
	return trueRangeIndicator{
		series: series,
		window: window,
	}
}

type trueRangeIndicator struct {
	series *TimeSeries
	window int
}

func (t trueRangeIndicator) Calculate(index int) big.Decimal {
	for i := index; i > index-t.window; i-- {
	}
	return big.ZERO
}
