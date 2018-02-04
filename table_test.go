package clitable

import (
	"strings"
	"testing"
)

func TestTableEmpty(t *testing.T) {
	// ARRANGE
	table := New()
	need := ""

	// ACT
	got := table.String()

	// ASSERT
	if got != need {
		t.Fatal("\nNeed:", need, "\nGot: ", got)
	}
}

func TestTableSingleRow(t *testing.T) {
	// ARRANGE
	table := New()
	need := 1

	// ACT
	table.AddRow("Test 1", "Test 2")
	s := table.String()
	got := strings.Count(s, "\n")

	// ASSERT
	if got != need {
		t.Fatal("\nNeed:", need, "\nGot: ", got)
	}
}

func TestTableMultipleRows(t *testing.T) {
	// ARRANGE
	table := New()
	need := 3

	// ACT
	table.AddRow("Test 1", "Test 2")
	table.AddRow("Test 3", "Test 4")
	table.AddRow("Test 5", "Test 6")
	s := table.String()
	got := strings.Count(s, "\n")

	// ASSERT
	if got != need {
		t.Fatal("\nNeed:", need, "\nGot: ", got)
	}
}

func TestTableOuterRightBorderTrue(t *testing.T) {
	// ARRANGE
	table := New()
	table.AddRow("1")
	need := true

	// ACT
	table.RightBorder = true
	s := table.String()
	got := strings.HasSuffix(s, table.ColSeparator+"\n")

	// ASSERT
	if got != need {
		t.Fatal("\nNeed:", need, "\nGot: ", got)
	}
}

func TestTableOuterRightBorderFalse(t *testing.T) {
	// ARRANGE
	table := New()
	table.AddRow("1")
	need := false

	// ACT
	table.RightBorder = false
	s := table.String()
	got := strings.HasSuffix(s, table.ColSeparator+"\n")

	// ASSERT
	if got != need {
		t.Fatal("\nNeed:", need, "\nGot: ", got)
	}
}
