package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
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
			return "", errors.New("Could not parse input to valid integer")
		}
		if len(gridRow) < gridColumn {
			return "", errors.New("Input value is not valid")
		}
		output = output + gridRow[gridColumn-1]
	}
	return output, nil
}

func main() {
	if os.Args[1] == "configure" {
		gridFile := "grid.txt"
		if os.Args[2] == "-f" {
			gridFile = os.Args[3]
		}
		reader := bufio.NewReader(os.Stdin)
		re := regexp.MustCompile(".,.,.,.,.")
		fmt.Println("Values must be separated by a comma, like A: 1,2,3,4,5")
		gridRows := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
		file, err := os.Create(gridFile)
		if err != nil {
			fmt.Println("Cannot create file "+gridFile, err)
		}
		defer file.Close()

		for _, row := range gridRows {
			fmt.Print(row + ": ")
			text, _ := reader.ReadString('\n')
			match := re.FindStringSubmatch(text)
			if len(match) == 0 {
				fmt.Println("invalid input")
				return
			}
			fmt.Fprintf(file, row+","+text)
		}
		return
	}
	fileName := defaultFileName
	input := os.Args[1]
	if os.Args[1] == "-f" {
		fileName = os.Args[2]
		input = os.Args[3]
	}
	gridFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		// os.Exit(1) does not work well with Example type tests
	}
	output, err := FindEntries(gridFile, input)
	if err != nil {
		fmt.Println(err)
		//os.Exit(1) does not work well with Example typ tests
	}
	fmt.Println(output)
}
