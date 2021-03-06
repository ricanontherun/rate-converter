package converter

const (
	IntervalMS = "ms"
	IntervalS  = "s"
	IntervalM  = "m"
	IntervalH  = "h"
	IntervalD  = "d"
)

var AvailableIntervals = []string{
	IntervalMS, // milliseconds
	IntervalS,  // seconds
	IntervalM,  // minutes
	IntervalH,  // hours
	IntervalD,  // days
}

func IsAvailableInterval(interval string) bool {
	for _, availableInterval := range AvailableIntervals {
		if interval == availableInterval {
			return true
		}
	}

	return false
}

// struct to represent an event rate, e.g something happened Count times per Interval
type EventRate struct {
	Count    float32
	Interval string
}

type conversionRateTable = map[string]float32
