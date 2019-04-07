package main

import (
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
	input := "[A3] [D1] [G4]"
	foundEntries := FindEntries(grid, input)
	if "314" != foundEntries {
		t.Errorf("Entries %v did not match 314", foundEntries)
	}
}

func TestFindEntriesForEmptyGrid(t *testing.T) {
	grid := strings.NewReader("")
	input := "[A3] [D1] [G4]"
	foundEntries := FindEntries(grid, input)
	if "" != foundEntries {
		t.Errorf("Expected empty found entries")
	}
}

func TestFindEntriesForOutOfBoundsInput(t *testing.T) {
	grid := strings.NewReader(sampleGridCard)
	input := "[A3] [D5] [G6]"
	foundEntries := FindEntries(grid, input)
	if "" != foundEntries {
		t.Errorf("Expected empty found entries")
	}
}
