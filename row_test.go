package clitable

import (
	"testing"
)

func TestRowEmpty(t *testing.T) {
	// ARRANGE
	need := 0

	// ACT
	row := newRow()
	got := len(row.cells)

	// ASSERT
	if got != need {
		t.Fatal("\nNeed:", need, "\nGot: ", got)
	}
}

func TestRowOneCell(t *testing.T) {
	// ARRANGE
	need := 1

	// ACT
	row := newRow("1")
	got := len(row.cells)

	// ASSERT
	if got != need {
		t.Fatal("\nNeed:", need, "\nGot: ", got)
	}
}
