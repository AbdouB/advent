package main

import (
	"embed"
	"encoding/csv"
	"errors"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

//go:embed sanitized_input.txt
var input embed.FS

type game struct {
	ID           int
	Combinations []struct {
		Red   int
		Blue  int
		Green int
	}
}

func (g *game) UnmarshalCSV(record []string) error {
	if len(record) < 1 {
		return errors.New("invalid record, minimum of 1 field required")
	}

	var err error

	g.ID, err = strconv.Atoi(record[0])
	if err != nil {
		return err
	}

	for _, draw := range record[1:] {
		for _, draw := range strings.Split(draw, ",") {
			var combination struct {
				Red   int
				Blue  int
				Green int
			}
			color := strings.Split(strings.TrimSpace(draw), " ")
			if len(color) != 2 {
				continue
			}

			count, err := strconv.Atoi(color[0])
			if err != nil {
				return err
			}

			switch color[1] {
			case "red":
				combination.Red = count
			case "blue":
				combination.Blue = count
			case "green":
				combination.Green = count
			default:
				return errors.New("invalid color")
			}

			g.Combinations = append(g.Combinations, combination)
		}
	}

	return nil
}

func (g *game) IsPossible(red, green, blue int) bool {
	for _, c := range g.Combinations {
		if c.Red > red || c.Green > green || c.Blue > blue {
			return false
		}
	}
	return true
}

func main() {
	data, err := input.Open("sanitized_input.txt")
	if err != nil {
		slog.Default().Error("failed to read input file", err)
		os.Exit(1)
	}

	games, err := parseGames(data)
	if err != nil {
		slog.Default().Error("failed to parse games", "message", err)
		os.Exit(1)
	}

	sumIDS := 0
	for _, g := range games {
		slog.Default().Info("game", "id", g.ID, "possible", g.IsPossible(12, 13, 14))
		if g.IsPossible(12, 13, 14) {
			sumIDS += g.ID
		}
	}
	slog.Default().Info("result", "sum", sumIDS)
}

func parseGames(data fs.File) ([]game, error) {
	reader := csv.NewReader(data)
	reader.Comma = ';'
	reader.TrimLeadingSpace = true

	var games []game
	for {
		record, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		g := game{}
		err = g.UnmarshalCSV(record)
		if err != nil {
			return nil, err
		}
		games = append(games, g)
	}
	return games, nil
}
