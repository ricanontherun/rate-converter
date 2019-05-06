package main

import (
	"errors"
	"flag"
	"fmt"
	"golang.org/x/text/message"
	"os"
	"rate-converter/internal/app/rate-converter"
	"regexp"
	"strconv"
	"strings"
)

// TODO: Would be very nice to query the system language.
var printer = message.NewPrinter(message.MatchLanguage("en"))
var convertErr error

// Parse source rate string into source rate structure.
func parseSourceRate(argument string) (*converter.EventRate, error) {
	parts := strings.Split(argument, "/")

	if len(parts) != 2 {
		return nil, errors.New("source rate should be in the form FREQUENCY/INTERVAL")
	}

	parsedNum, atoiErr := strconv.ParseFloat(parts[0], 32)
	if atoiErr != nil {
		return nil, errors.New("source rate should be in the form FREQUENCY/INTERVAL")
	}

	interval := strings.ToLower(parts[1])
	if !converter.IsAvailableInterval(interval) {
		return nil, errors.New(fmt.Sprintf("%s is not a valid interval", interval))
	}

	source := &converter.EventRate{
		Count:    float32(parsedNum),
		Interval: interval,
	}

	return source, nil
}

// Parse a target rate string into a target rate structure.
func parseTargetRate(argument string) (*converter.EventRate, error) {
	regexPattern := "^(?P<Count>\\d+)?(?P<Interval>%s)$"
	re, reErr := regexp.Compile(fmt.Sprintf(regexPattern, strings.Join(converter.AvailableIntervals, "|")))

	if reErr != nil {
		return nil, reErr
	}

	targetRate := &converter.EventRate{}

	matches := re.FindStringSubmatch(argument)
	matchesLen := len(matches)

	if matchesLen == 0 {
		return nil, errors.New("invalid target rate")
	}

	for i, name := range re.SubexpNames() {
		if i > 0 && i <= matchesLen {
			match := matches[i]

			switch name {
			case "Count":
				if match == "" {
					targetRate.Count = 1
				} else {
					count, atoiErr := strconv.Atoi(matches[i])

					if atoiErr != nil {
						return nil, atoiErr
					}

					targetRate.Count = float32(count)
				}
			case "Interval":
				targetRate.Interval = matches[i]
			}
		}
	}

	return targetRate, nil
}

// Attempt to print a nicely formatted result. If the required package isn't installed
// fallback to system fmt
func printResult(result *converter.EventRate, precision int) {
	if _, printErr := printer.Println(fmt.Sprintf("%.*f", precision, result.Count)); printErr != nil {
		fmt.Println(result.Count)
	}
}

func main() {
	flag.Usage = func() {
		if convertErr != nil {
			_,_ = fmt.Fprintln(flag.CommandLine.Output(), fmt.Sprintf("Error: %s", convertErr.Error()))
		}

		_,_ = fmt.Fprintln(flag.CommandLine.Output(), fmt.Sprintf("Usage of %s: ", os.Args[0]))

		flag.PrintDefaults()

		os.Exit(1)
	}

	precisionFlag := flag.Int("precision", 2, "Display precision (think %.Nf)")
	sourceRateFlag := flag.String("source", "", "Source rate, e.g 10/s. Available intervals: " + strings.Join(converter.AvailableIntervals, ","))
	targetRateFlag := flag.String("target", "", "Target rate, e.g 30h")

	flag.Parse()

	sourceRate := *sourceRateFlag
	targetRate := *targetRateFlag

	if sourceRate == "" || targetRate == "" {
		flag.Usage()
	}

	source, convertErr := parseSourceRate(sourceRate)
	if convertErr != nil {
		flag.Usage()
	}

	target, convertErr := parseTargetRate(targetRate)
	if convertErr != nil {
		flag.Usage()
	}

	if conversionErr := converter.DoConversion(source, target); conversionErr != nil {
		fmt.Println(conversionErr.Error())
		os.Exit(1)
	}

	printResult(target, *precisionFlag)
}
