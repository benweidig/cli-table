package clitable

import (
	"testing"
)

func TestCellContentInt(t *testing.T) {
	// ARRANGE
	c := newCell(5)
	need := "5"

	// ACT
	got := c.content

	// ASSERT
	if got != need {
		t.Fatal("\nNeed:", need, "\nGot: ", got)
	}
}

func TestCellContentString(t *testing.T) {
	// ARRANGE
	s := "This is a test"
	c := newCell(s)
	need := s

	// ACT
	got := c.content

	// ASSERT
	if got != need {
		t.Fatal("\nNeed:", need, "\nGot: ", got)
	}
}

func TestCellWidth(t *testing.T) {
	// ARRANGE
	s := "This is a test"
	c := newCell(s)
	need := len(s)

	// ACT
	got := c.width

	// ASSERT
	if got != need {
		t.Fatal("\nNeed:", need, "\nGot: ", got)
	}
}

func TestCellWidthHalfWidthKatakana(t *testing.T) {
	// ARRANGE
	s := "ｨｨ 4"
	c := newCell(s)
	need := 4 // Byte length is actually 8

	// ACT
	got := c.width

	// ASSERT
	if got != need {
		t.Fatal("\nNeed:", need, "\nGot: ", got)
	}
}

func TestCellPaddedContent(t *testing.T) {
	// ARRANGE
	s := "Test"
	c := newCell(s)
	need := "Test  "

	// ACT
	got := c.paddedContent(6)

	// ASSERT
	if got != need {
		t.Fatal("\nNeed:", need, "\nGot: ", got)
	}
}
