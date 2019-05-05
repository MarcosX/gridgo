package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func configure() {
	gridRows := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}

	gridFile := defaultFileName
	if len(os.Args) == 4 && os.Args[2] == "-f" {
		gridFile = os.Args[3]
	}
	fmt.Println("Saving grid info to file " + gridFile)
	file, err := os.Create(gridFile)
	if err != nil {
		fmt.Println("Cannot create file "+gridFile, err)
		return
	}
	defer file.Close()

	re := regexp.MustCompile(".,.,.,.,.")
	reader := bufio.NewReader(os.Stdin)
	for _, row := range gridRows {
		fmt.Print(row + ": ")
		text, _ := reader.ReadString('\n')
		if len(re.FindStringSubmatch(text)) == 0 {
			fmt.Println("Values must be separated by a comma, like A: 1,2,3,4,5")
			fmt.Println("invalid input " + text)
			return
		}
		fmt.Fprintf(file, row+","+text)
	}
}

func read() {
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
		return
	}
	output, err := FindEntries(gridFile, input)
	if err != nil {
		fmt.Println(err)
		//os.Exit(1) does not work well with Example typ tests
		return
	}
	fmt.Println(output)
}

func main() {
	if os.Args[1] == "configure" {
		configure()
		return
	}
	read()
}
