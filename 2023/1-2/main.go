package main

import (
	"embed"
	"log/slog"
	"os"
	"regexp"
	"strings"
)

//go:embed input.txt
var input embed.FS

var regNumbers = regexp.MustCompile(`([0-9]|zero|one|two|three|four|five|six|seven|eight|nine)`)
var numbersMap = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func main() {
	data, err := input.ReadFile("input.txt")
	if err != nil {
		slog.Default().Error("failed to read input file", err)
		os.Exit(1)
	}

	total := 0
	for lid, line := range strings.Split(string(data), "\n") {
		matches := regNumbers.FindStringIndex(line)
		if len(matches) == 0 {
			slog.Default().Info("no numbers found", "line", lid)
			continue
		}

		firstMatch, lastMatch := line[matches[0]:matches[1]], line[matches[0]:matches[1]]

		// since go doesn't support lookbehind/ahead, we need to do this the loopy way
		// we start at the end of the first match and look for the next match
		// this way if there were overlapping matches we'll be able to find them
		// example:
		// string to search: "sevenine"
		// first match: "seven"
		// leaving us with: "ine" which doesn't match, so we move one char back
		// new string to search: "nine"
		// second match: "nine"
		i := matches[1]
		for i <= len(line) {
			subString := line[i:]
			matches := regNumbers.FindStringIndex(subString)
			if len(matches) == 0 {
				break
			}
			lastMatch = subString[matches[0]:matches[1]]
			if len(lastMatch) > 1 {
				i += matches[1] - 1
			} else {
				i += matches[1]
			}
		}

		first, last := numbersMap[firstMatch], numbersMap[lastMatch]

		calibratedValue := (first * 10) + last
		slog.Default().Info("calibrated value",
			"id", lid,
			"line", line,
			"first_match", firstMatch,
			"first", first,
			"last_match", lastMatch,
			"last", last,
			"value", calibratedValue)

		total += calibratedValue
	}
	slog.Default().Info("result", "total", total)
}
