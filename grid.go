package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
func FindEntries(gridInput io.Reader, input string) string {
	output := ""
	grid := buildGrid(gridInput)
	if len(grid) == 0 {
		return ""
	}
	for _, entryInput := range strings.Fields(input) {
		entryData := strings.Split(entryInput, "")
		gridRow := grid[entryData[1]]
		gridColumn, err := strconv.Atoi(entryData[2])
		if err != nil {
			return ""
		}
		if len(gridRow) < gridColumn {
			return ""
		}
		output = output + gridRow[gridColumn-1]
	}
	return output
}

// Saves the grid configuration into a known formated file
func ConfigureGrid(input []string) map[string][]string {
	grid := make(map[string][]string)
	gridRows := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for rowIndex, row := range gridRows {
		grid[row] = strings.Split(input[rowIndex], " ")
	}
	return grid
}

func main() {
	fileName := defaultFileName
	input := os.Args[1]
	if os.Args[1] == "-f" {
		fileName = os.Args[2]
		input = os.Args[3]
	}
	gridFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(FindEntries(gridFile, input))
	// reader := bufio.NewReader(os.Stdin)
	// text, _ := reader.ReadString('\n')
}
