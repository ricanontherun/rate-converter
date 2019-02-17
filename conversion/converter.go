package conversion

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

func DoConversion(source *EventRate, target *EventRate) error {
	if target.Interval == source.Interval {
		target.Count = source.Count * target.Count
		return nil
	}

	conversationRate, exists := conversionTable[source.Interval][target.Interval]

	if !exists {
		return errors.New(fmt.Sprintf("unsupported conversion from %s to %s", source.Interval, target.Interval))
	}

	target.Count = source.Count * conversationRate * target.Count

	return nil
}
