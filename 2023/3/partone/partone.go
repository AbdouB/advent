package partone

import (
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numRegex = regexp.MustCompile(`\d+`)

func Solve(data []byte) {
	total := 0
	lines := strings.Split(string(data), "\n")
	for lid, line := range lines {
		for _, num := range numRegex.FindAllStringIndex(line, -1) {
			if isSymbolAround(lines, lid, num[0], num[1]) {
				number, err := strconv.Atoi(line[num[0]:num[1]])
				if err != nil {
					slog.Default().Error("failed to parse number", err)
					os.Exit(1)
				}
				total += number
			}
		}
	}
	slog.Default().Info("result", "total", total)
}

func isSymbolAround(lines []string, lid, x, y int) bool {
	for i := lid - 1; i <= lid+1; i++ {
		if i < 0 || i >= len(lines) {
			continue
		}
		for j := x - 1; j <= y; j++ {
			if j < 0 || j >= len(lines[i]) {
				continue
			}
			if _, ok := symbols[rune(lines[i][j])]; ok {
				return true
			}
		}
	}
	return false
}

// symbols map
// @, *, #, %, &, +, -, =, |, ~, ^, <, >, /, \, !, ?, $, :, [, ], {, }, (, )
var symbols = map[rune]bool{
	'@':  true,
	'*':  true,
	'#':  true,
	'%':  true,
	'&':  true,
	'+':  true,
	'-':  true,
	'=':  true,
	'|':  true,
	'~':  true,
	'^':  true,
	'<':  true,
	'>':  true,
	'/':  true,
	'\\': true,
	'!':  true,
	'?':  true,
	'$':  true,
	':':  true,
	'[':  true,
	']':  true,
	'{':  true,
	'}':  true,
	'(':  true,
	')':  true,
}
