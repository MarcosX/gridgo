package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

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

func FindEntries(gridInput io.Reader, input string) string {
	output := ""
	grid := buildGrid(gridInput)
	if len(grid) == 0 {
		return ""
	}
	for _, entryInput := range strings.Fields(input) {
		entryData := strings.Split(entryInput, "")
		gridRow := grid[entryData[1]]
		gridColumn, _ := strconv.ParseInt(entryData[2], 16, 32)
		output = output + gridRow[gridColumn-1]
	}
	return output
}
