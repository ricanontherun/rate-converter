package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/ricanontherun/rate-converter/internal/app/rate-converter"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
	"os"
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

	maybeFrequency := strings.Replace(parts[0], ",", "", -1)
	parsedNum, parseErr := strconv.ParseFloat(maybeFrequency, 32)
	if parseErr != nil {
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

// Attempt to print a nicely formatted result.
func printResult(result float32) {
	var printErr error

	// Fallback to fmt on package error.
	output := printer.Sprintf("%v", number.Decimal(result))
	if _, printErr = printer.Println(output); printErr != nil {
		fmt.Println(result)
	}
}

func main() {
	flag.Usage = func() {
		if convertErr != nil {
			_, _ = fmt.Fprintln(flag.CommandLine.Output(), fmt.Sprintf("Error: %s", convertErr.Error()))
		}

		_, _ = fmt.Fprintln(flag.CommandLine.Output(), fmt.Sprintf("Usage of %s: ", os.Args[0]))

		flag.PrintDefaults()

		os.Exit(1)
	}

	sourceRateFlag := flag.String("source", "", "Source rate, e.g 10/s. Available intervals: "+strings.Join(converter.AvailableIntervals, ","))
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

	if result, conversionErr := converter.DoConversion(source, target); conversionErr != nil {
		fmt.Println(conversionErr.Error())
		os.Exit(1)
	} else {
		printResult(result)
	}
}
