package converter_test

import (
	"github.com/ricanontherun/rate-converter/internal/app/rate-converter"
	"testing"
)

func TestConversions(t *testing.T) {
	var conversionTests = []struct {
		sourceCount    float32
		sourceInterval string

		targetCount    float32
		targetInterval string

		Result float32
	}{
		{1.0, converter.IntervalMS, 5.0, converter.IntervalMS, 5.0},
		{5.0, converter.IntervalS, 15.0, converter.IntervalM, 4500.00},
		{60.0, converter.IntervalH, 30, converter.IntervalD, 43200},
	}

	source := &converter.EventRate{}
	target := &converter.EventRate{}

	for _, test := range conversionTests {
		source.Count = test.sourceCount
		source.Interval = test.sourceInterval

		target.Count = test.targetCount
		target.Interval = test.targetInterval

		if result, conversionErr := converter.DoConversion(source, target); conversionErr != nil {
			t.Errorf("Failed to convert: %s\n", conversionErr.Error())
		} else {
			if result != test.Result {
				t.Errorf("Received %f, expected %f", result, test.Result)
			}
		}
	}
}
