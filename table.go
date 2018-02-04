package clitable

import (
	"bytes"
	"sync"
)

type Table struct {
	rows []*row
	mtx  *sync.RWMutex

	// The column separator
	ColSeparator string

	// Add an additional horizontal padding to the right of the content to cells
	AdditionalRightPadding int

	// Print the border at the right outer side of the table
	OuterRightBorder bool

	// Add a border under the first row
	HeaderBorder bool

	// Header separator for border under the first row
	HeaderSeparator byte
}

// Creates a new Table with sensible defaults
func New() *Table {
	return &Table{
		ColSeparator:           " | ",
		AdditionalRightPadding: 0,
		OuterRightBorder:       false,
		HeaderBorder:           false,
		HeaderSeparator:        '-',
		mtx:                    new(sync.RWMutex),
	}
}

// Adds a row of data to the table. Col count doesn't matter.
func (t *Table) AddRow(contents ...interface{}) {
	// We don't want to have a half-build table so we need a lock for updating content
	t.mtx.Lock()
	defer t.mtx.Unlock()

	row := newRow(contents...)
	t.rows = append(t.rows, row)
}

// Returns string representation of the table
func (t *Table) String() string {
	// We want to make sure the data won't change mid string building
	t.mtx.RLock()
	defer t.mtx.RUnlock()

	// Empty table == empty string
	if len(t.rows) == 0 {
		return ""
	}

	// Determinate the width of each column
	var colWidths []int
	for _, row := range t.rows {
		for i, cell := range row.cells {
			if i+1 > len(colWidths) {
				colWidths = append(colWidths, 0)
			}

			if cell.width > colWidths[i] {
				colWidths[i] = cell.width
			}
		}
	}

	// Check if we want the outer right border
	cols := len(colWidths)
	borderedCols := cols
	if t.OuterRightBorder == false {
		borderedCols -= 1
	}

	// Holds the string representation of the table
	var buf bytes.Buffer

	// Build table data
	for rowIdx, row := range t.rows {
		for colIdx := 0; colIdx < cols; colIdx++ {
			colWidth := colWidths[colIdx]

			// Rows don't need to have the same amount of cells so we might need to fill up
			// the empty cells with spaces
			if colIdx < len(row.cells) {
				cell := row.cells[colIdx]
				buf.WriteString(cell.paddedContent(colWidth))
			} else {
				//				if t.OuterRightBorder == false && colIdx < cols {
				if t.OuterRightBorder == true || colIdx < cols-1 {
					for i := 0; i < colWidth; i++ {
						buf.WriteByte(' ')
					}
				}
			}

			if t.AdditionalRightPadding > 0 {
				for i := 0; i < t.AdditionalRightPadding; i++ {
					buf.WriteByte(' ')
				}
			}

			if colIdx < borderedCols {
				buf.WriteString(t.ColSeparator)
			}
		}
		buf.WriteString("\n")

		// Check if we need to print the header border
		if rowIdx == 1 && t.HeaderBorder {
			for colIdx := 0; colIdx < cols; colIdx++ {
				colWidth := colWidths[colIdx]
				for i := 0; i < colWidth; i++ {
					buf.WriteByte(t.HeaderSeparator)
				}
				if t.AdditionalRightPadding > 0 {
					for i := 0; i < t.AdditionalRightPadding; i++ {
						buf.WriteByte(t.HeaderSeparator)
					}
				}
				if colIdx < borderedCols {
					buf.WriteString(t.ColSeparator)
				}

			}
			buf.WriteString("\n")
		}

	}

	return buf.String()
}
