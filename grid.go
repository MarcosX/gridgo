package main

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

const defaultFileName = "grid.txt"

func buildGrid(gridInput io.Reader) map[string][]string {
	scanner := bufio.NewScanner(gridInput)
	grid := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		lineData := strings.Split(line, ",")
		grid[lineData[0]] = lineData[1:]
	}
	return grid
}

// Searches on the reader for entries as in a two dimensional matrix
func FindEntries(gridInput io.Reader, input string) (string, error) {
	output := ""
	grid := buildGrid(gridInput)
	if len(grid) == 0 {
		return "", errors.New("Could not read grid values")
	}
	for _, entryInput := range strings.Fields(input) {
		entryData := strings.Split(entryInput, "")
		gridRow := grid[entryData[1]]
		gridColumn, err := strconv.Atoi(entryData[2])
		if err != nil {
			return "", errors.New("Could not parse input " + entryData[2] + " to valid integer")
		}
		if len(gridRow) < gridColumn {
			return "", errors.New("Input value is not valid")
		}
		output = output + gridRow[gridColumn-1]
	}
	return output, nil
}
