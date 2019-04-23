package main

import (
	"os"
	"strings"
	"testing"
)

const sampleGridCard = `
A,1,2,3,4,5
B,1,2,3,4,5
C,1,2,3,4,5
D,1,2,3,4,5
E,1,2,3,4,5
F,1,2,3,4,5
G,1,2,3,4,5
H,1,2,3,4,5
I,1,2,3,4,5
J,1,2,3,4,5
`

func TestBuildGrid(t *testing.T) {
	gridInput := strings.NewReader(sampleGridCard)
	grid := buildGrid(gridInput)
	if len(grid) != 10 {
		t.Errorf("Grid size is not 10")
	}
	if len(grid["A"]) != 5 {
		t.Errorf("Entry size on grid is not 5")
	}
}

func TestFindEntries(t *testing.T) {
	grid := strings.NewReader(sampleGridCard)
	input := "[A3] [D5] [G4]"
	foundEntries, err := FindEntries(grid, input)
	if "354" == foundEntries && err == nil {
		return
	}
	t.Errorf("Entries %v did not match 354", foundEntries)
}

func TestFindEntriesForEmptyGrid(t *testing.T) {
	grid := strings.NewReader("")
	input := "[A3] [D1] [G4]"
	foundEntries, err := FindEntries(grid, input)
	if "" == foundEntries && err != nil {
		return
	}
	t.Errorf("Expected empty entries with an error")
}

func TestFindEntriesForOutOfBoundsInput(t *testing.T) {
	grid := strings.NewReader(sampleGridCard)
	input := "[A3] [D5] [G6]"
	foundEntries, err := FindEntries(grid, input)
	if "" == foundEntries && err != nil {
		return
	}
	t.Errorf("Expected empty entries with an error")
}

func TestFindEntriesForMalFormedInput(t *testing.T) {
	grid := strings.NewReader(sampleGridCard)
	input := "[AA] [D5] [G6]"
	foundEntries, err := FindEntries(grid, input)
	if "" == foundEntries && err != nil {
		return
	}
	t.Errorf("Expected empty entries with an error")
}

func TestFindEntriesForNotFoundINput(t *testing.T) {
	grid := strings.NewReader(sampleGridCard)
	input := "[X3] [D5] [G6]"
	foundEntries, err := FindEntries(grid, input)
	if "" == foundEntries && err != nil {
		return
	}
	t.Errorf("Expected empty entries with an error")
}

func TestConfigureGrid(t *testing.T) {
	input := []string{"1 2 3 4 5", "1 2 3 4 5",
		"1 2 3 4 5", "1 2 3 4 5", "1 2 3 4 5",
		"1 2 3 4 5", "1 2 3 4 5", "1 2 3 4 5",
		"1 2 3 4 5", "1 2 3 4 5"}

	grid := ConfigureGrid(input)
	if len(grid) != 10 {
		t.Errorf("Expected 10 entries for the grid")
	}
	if len(grid["A"]) != 5 {
		t.Errorf("Expected grid entries to have 5 chars")
	}
}

func ExampleReadingInputWithDefaultFile() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "[A3] [D5] [G2]"}
	main()
	// Output:
	// J0V
}

func ExampleReadingInputWithCustomFile() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "-f", "sample_grid.txt", "[A3] [D5] [G2]"}
	main()
	// Output:
	// 352
}

func ExampleMalformedInput() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "not", "valid", "arg"}
	main()
	// Output:
	// Could not parse input to valid integer
}
