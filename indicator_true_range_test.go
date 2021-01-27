package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrueRangeIndicator(t *testing.T) {
	series := mockTimeSeriesOCHL(
		[]string{"0", "1.2919", "0", "0"},
		[]string{"0", "1.2884", "1.2942", "1.2842"},
		[]string{"0", "1.2881", "1.2929", "1.2846"},
		[]string{"0", "1.2836", "1.2889", "1.2796"},
		[]string{"0", "1.2881", "1.2900", "1.2819"},
		[]string{"0", "1.2905", "1.2933", "1.2840"},
		[]string{"0", "1.2857", "1.2997", "1.2833"},
		[]string{"0", "1.2932", "1.2956", "1.2821"},
		[]string{"0", "0", "1.2993", "1.2904"},
	)

	trInd := NewTrueRangeIndicator(series, 7)
	assert.EqualValues(t, "0.0107", trInd.Calculate(7).FormattedString(4))
}
