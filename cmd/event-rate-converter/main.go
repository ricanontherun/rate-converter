package main

import (
	"errors"
	"event-rate-converter/internal/app/event-rate-converter"
	"fmt"
	"golang.org/x/text/message"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseSourceRate(argument string) (*converter.EventRate, error) {
	parts := strings.Split(argument, "/")

	if len(parts) != 2 {
		return nil, errors.New("source rate should be in the form FREQUENCY/INTERVAL")
	}

	parsedNum, atoiErr := strconv.ParseFloat(parts[0], 64)
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

func help(err error) {
	helpString := ""

	if err != nil {
		helpString += fmt.Sprintf("Error: %s\n\n", err.Error())
	}

	helpString += "usage: event-rate-converter SOURCE_RATE TARGET_RATE\n\n"
	helpString += "Available intervals: " + strings.Join(converter.AvailableIntervals, ",") + "\n"

	helpString += "\nExamples:\n"
	helpString += "\tevent-rate-converter 1000/s ms\n"
	helpString += "\t1\n"
	helpString += "\tevent-rate-converter 1/ms s\n"
	helpString += "\t1,000\n"

	fmt.Println(helpString)
	os.Exit(1)
}

func main() {
	args := os.Args[1:]
	argLen := len(args)

	if argLen == 1 && args[0] == "--help" {
		help(nil)
	}

	if len(args) != 2 {
		help(errors.New("incorrect amount of arguments"))
	}

	source, sourceErr := parseSourceRate(args[0])
	if sourceErr != nil {
		help(sourceErr)
	}

	target, targetErr := parseTargetRate(args[1])
	if targetErr != nil {
		help(targetErr)
	}

	if conversionErr := converter.DoConversion(source, target); conversionErr != nil {
		fmt.Println(conversionErr.Error())
		os.Exit(1)
	}

	printer := message.NewPrinter(message.MatchLanguage("en"))
	if _, printErr := printer.Println(target.Count); printErr != nil { // Fallback to fmt
		fmt.Println(target.Count)
	}
}
