package converter

import (
	"errors"
	"fmt"
)

const (
	millisecondsPerSecond = 1000.0
	secondsPerMinute      = 60.0
	minutesPerHour        = 60.0
	hoursPerDay           = 24.0
)

// Convert one event rate to another.
func DoConversion(source *EventRate, target *EventRate) (float32, error) {
	if target.Interval == source.Interval {
		result := source.Count * target.Count
		return result, nil
	}

	conversationRate, exists := conversionTable[source.Interval][target.Interval]
	if !exists {
		return 0, errors.New(fmt.Sprintf("unsupported conversion from %s to %s", source.Interval, target.Interval))
	}

	result := source.Count * conversationRate * target.Count

	return result, nil
}
