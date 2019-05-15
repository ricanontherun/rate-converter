package converter

var conversionTable = map[string]conversionRateTable{
	IntervalMS: {
		IntervalS: millisecondsPerSecond,
		IntervalM: millisecondsPerSecond * secondsPerMinute,
		IntervalH: millisecondsPerSecond * secondsPerMinute * minutesPerHour,
		IntervalD: millisecondsPerSecond * secondsPerMinute * minutesPerHour * hoursPerDay,
	},

	IntervalS: {
		IntervalMS: 1.0 / millisecondsPerSecond,
		IntervalM:  secondsPerMinute,
		IntervalH:  secondsPerMinute * minutesPerHour,
		IntervalD:  secondsPerMinute * minutesPerHour * hoursPerDay,
	},

	IntervalM: {
		IntervalMS: 1.0 / (secondsPerMinute * millisecondsPerSecond),
		IntervalS:  1.0 / secondsPerMinute,
		IntervalH:  minutesPerHour,
		IntervalD:  minutesPerHour * hoursPerDay,
	},

	IntervalH: {
		IntervalMS: 1.0 / (minutesPerHour * secondsPerMinute * millisecondsPerSecond),
		IntervalS:  1.0 / (minutesPerHour * secondsPerMinute),
		IntervalM:  1.0 / (minutesPerHour),
		IntervalD:  hoursPerDay,
	},

	IntervalD: {
		IntervalMS: 1.0 / (hoursPerDay * minutesPerHour * secondsPerMinute * millisecondsPerSecond),
		IntervalS:  1.0 / (hoursPerDay * minutesPerHour * secondsPerMinute),
		IntervalM:  1.0 / (hoursPerDay * minutesPerHour),
		IntervalH:  1.0 / hoursPerDay,
	},
}
