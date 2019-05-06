package main

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

const defaultFileName = "grid.txt"

func buildGrid(gridFile io.Reader) map[string][]string {
	scanner := bufio.NewScanner(gridFile)
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
func FindEntries(gridFile io.Reader, rawInput string) (string, error) {
	output := ""
	grid := buildGrid(gridFile)
	if len(grid) == 0 {
		return "", errors.New("Could not read grid values")
	}
	for _, singleInput := range strings.Fields(rawInput) {
		singleInputChars := strings.Split(singleInput, "")
		gridRow := grid[singleInputChars[1]]
		gridColumnIndex, err := strconv.Atoi(singleInputChars[2])
		if err != nil {
			return "", errors.New("Could not parse input " + singleInput + " to valid grid entry")
		}
		if len(gridRow) < gridColumnIndex {
			return "", errors.New("Input value " + singleInput + " is not valid")
		}
		output = output + gridRow[gridColumnIndex-1]
	}
	return output, nil
}
