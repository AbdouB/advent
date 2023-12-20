package main

import (
	"embed"
	"flag"
	"log/slog"
	"os"

	"github.com/abdoub/advent/2023/3/partone"
	"github.com/abdoub/advent/2023/3/parttwo"
)

//go:embed input.txt
var input embed.FS

func main() {
	data, err := input.ReadFile("input.txt")
	if err != nil {
		slog.Default().Error("failed to read input file", err)
		os.Exit(1)
	}

	p := flag.Int("p", 1, "puzzle part to solve")
	flag.Parse()

	switch *p {
	case 1:
		partone.Solve(data)
	case 2:
		parttwo.Solve(data)
	}
}
