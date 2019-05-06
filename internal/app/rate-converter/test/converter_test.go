package converter_test

import (
	"github.com/ricanontherun/rate-converter/internal/app/rate-converter"
	"testing"
)

func TestConversions(t *testing.T) {
	var conversionTests = []struct {
		sourceCount    float32
		sourceInterval string

		expectedCount  float32
		targetInterval string
	}{
		{1.0, converter.IntervalMS, 1.0, converter.IntervalMS},
		{1.0, converter.IntervalMS, 1000.0, converter.IntervalS},
		{1.0, converter.IntervalMS, 60000.00, converter.IntervalM},
		{1.0, converter.IntervalMS, 3600000, converter.IntervalH},
		{1.0, converter.IntervalMS, 86400000, converter.IntervalD},

		{1.0, converter.IntervalS, .001, converter.IntervalMS},
		{1.0, converter.IntervalS, 1.0, converter.IntervalS},
		{1.0, converter.IntervalS, 60.0, converter.IntervalM},
		{1.0, converter.IntervalS, 3600.0, converter.IntervalH},
		{1.0, converter.IntervalS, 86400.00, converter.IntervalD},

		{15000.0, converter.IntervalM, .25, converter.IntervalMS},
	}

	source := &converter.EventRate{}
	target := &converter.EventRate{}

	for _, test := range conversionTests {
		source.Count = test.sourceCount
		source.Interval = test.sourceInterval

		target.Interval = test.targetInterval

		if conversionErr := converter.DoConversion(source, target); conversionErr != nil {
			t.Errorf("Failed to convert: %s\n", conversionErr.Error())
		} else {
			if target.Count != test.expectedCount {
				t.Errorf("Received %f, expected %f", target.Count, test.expectedCount)
			}
		}
	}
}
