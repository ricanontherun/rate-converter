package converter_test

import (
	"event-rate-converter/internal/app/rate-converter"
	"testing"
)

func TestConversions(t *testing.T) {
	var conversionTests = []struct {
		sourceCount    float64
		sourceInterval string

		expectedCount  float64
		targetInterval string
	}{
		{1.0, conversion.IntervalMS, 1.0, conversion.IntervalMS},
		{1.0, conversion.IntervalMS, 1000.0, conversion.IntervalS},
		{1.0, conversion.IntervalMS, 60000.00, conversion.IntervalM},
		{1.0, conversion.IntervalMS, 3600000, conversion.IntervalH},
		{1.0, conversion.IntervalMS, 86400000, conversion.IntervalD},

		{1.0, conversion.IntervalS, .001, conversion.IntervalMS},
		{1.0, conversion.IntervalS, 1.0, conversion.IntervalS},
		{1.0, conversion.IntervalS, 60.0, conversion.IntervalM},
		{1.0, conversion.IntervalS, 3600.0, conversion.IntervalH},
		{1.0, conversion.IntervalS, 86400.00, conversion.IntervalD},

		{15000.0, conversion.IntervalM, .25, conversion.IntervalMS},
	}

	source := &conversion.EventRate{}
	target := &conversion.EventRate{}

	for _, test := range conversionTests {
		source.Count = test.sourceCount
		source.Interval = test.sourceInterval

		target.Interval = test.targetInterval

		if conversionErr := conversion.DoConversion(source, target); conversionErr != nil {
			t.Errorf("Failed to convert: %s\n", conversionErr.Error())
		} else {
			if target.Count != test.expectedCount {
				t.Errorf("Received %f, expected %f", target.Count, test.expectedCount)
			}
		}
	}
}
