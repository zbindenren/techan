package techan

import (
	"github.com/sdcoffey/big"
)

// https://www.earnforex.com/guides/average-true-range/
// https://www.tradinformed.com/calculate-supertrend-indicator-using-excel/
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
	sum := big.ZERO
	for i := Max(1, index-t.window+1); i <= index; i++ {
		sum = sum.Add(max(t.series.Candles[i].MaxPrice, t.series.Candles[i-1].ClosePrice).Sub(min(t.series.Candles[i].MinPrice, t.series.Candles[i-1].ClosePrice)))
	}

	return sum.Div(big.NewDecimal(float64(t.window)))
}

func max(a, b big.Decimal) big.Decimal {
	if a.LT(b) {
		return b
	}

	return a
}

func min(a, b big.Decimal) big.Decimal {
	if a.GT(b) {
		return b
	}

	return a
}
