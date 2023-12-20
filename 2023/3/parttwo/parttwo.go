package parttwo

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var gearRegex = regexp.MustCompile(`\*`)

func Solve(data []byte) {
	total := 0
	lines := strings.Split(string(data), "\n")
	for lid, line := range lines {
		for _, num := range gearRegex.FindAllStringIndex(line, -1) {
			if nmbrs := numbersAround(lines, lid, num[0], num[1]); len(nmbrs) > 1 {
\				total += multiply(nmbrs)
			}
		}
	}
	slog.Default().Info("result", "total", total)
}

func multiply(nmbrs []int) int {
	total := 1
	for _, nmbr := range nmbrs {
		total *= nmbr
	}
	return total
}

func numbersAround(lines []string, lid, x, y int) []int {
	var nmbrs []int
	for i := lid - 1; i <= lid+1; i++ { // lines
		if i < 0 || i >= len(lines) {
			continue
		}
		for j := x - 1; j <= y; j++ { //columns
			if j < 0 || j >= len(lines[i]) {
				continue
			}
			if unicode.IsDigit(rune(lines[i][j])) {
				nmbr := captureNumber(lines[i], j)
				if len(nmbrs) >= 1 && nmbrs[len(nmbrs)-1] == nmbr { // lazy way to avoid duplicates
					continue
				}
				nmbrs = append(nmbrs, nmbr)
			}
		}
	}
	return nmbrs
}

func captureNumber(line string, pos int) int {
	var number string
	for i := pos; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			number += string(line[i])
		} else {
			break
		}
	}
	for i := pos - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			number = string(line[i]) + number
		} else {
			break
		}
	}
	nmbr, err := strconv.Atoi(number)
	if err != nil {
		slog.Default().Error("failed to parse number", err)
		os.Exit(1)
	}
	return nmbr
}
