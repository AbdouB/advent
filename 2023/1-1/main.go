package main

import (
	"embed"
	"log/slog"
	"os"
	"strings"
	"unicode"
)

//go:embed input.txt
var input embed.FS

func main() {
	data, err := input.ReadFile("input.txt")
	if err != nil {
		slog.Default().Error("failed to read input file", err)
		os.Exit(1)
	}

	total := 0
	for lid, line := range strings.Split(string(data), "\n") {
		first, last := 0, 0
		for _, c := range []rune(line) {
			if unicode.IsDigit(c) {
				if first == 0 {
					first, last = int(c-'0'), int(c-'0')
					continue
				}
				last = int(c - '0')
			}
		}
		calibratedValue := (first * 10) + last

		slog.Default().Info("calibrated value",
			"line", lid,
			"first", first,
			"last", last,
			"value", calibratedValue)
		total += calibratedValue
	}
	slog.Default().Info("result", "total", total)
}
