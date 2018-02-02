package clitable

import (
	"bytes"
	"sync"
)

type Table struct {
	rows []*row

	ColSeparator           string
	AdditionalRightPadding int
	OuterRightBorder       bool
	HeaderBorder           bool
	HeaderSeparator        byte

	mtx *sync.RWMutex
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
func (table *Table) AddRow(contents ...interface{}) {
	table.mtx.Lock()
	defer table.mtx.Unlock()

	row := newRow(contents...)
	table.rows = append(table.rows, row)
}

// Returns string representation of the table
func (table *Table) String() string {
	table.mtx.RLock()
	defer table.mtx.RUnlock()

	// Empty table == empty string
	if len(table.rows) == 0 {
		return ""
	}

	// Determinate each column width
	var colWidths []int
	for _, row := range table.rows {

		// Change size of widths if necessary

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
	if table.OuterRightBorder == false {
		borderedCols -= 1
	}

	// Our string representation
	var buf bytes.Buffer

	// Build table data
	for rowIdx, row := range table.rows {

		for colIdx := 0; colIdx < cols; colIdx++ {
			colWidth := colWidths[colIdx]

			if colIdx < len(row.cells) {
				cell := row.cells[colIdx]
				buf.WriteString(cell.paddedContent(colWidth))
			} else {
				for i := 0; i < colWidth; i++ {
					buf.WriteByte(' ')
				}
			}

			if table.AdditionalRightPadding > 0 {
				for i := 0; i < table.AdditionalRightPadding; i++ {
					buf.WriteByte(' ')
				}
			}

			if colIdx < borderedCols {
				buf.WriteString(table.ColSeparator)
			}

		}
		buf.WriteString("\n")

		// Check if we need to print the header border
		if rowIdx == 1 && table.HeaderBorder {
			for colIdx := 0; colIdx < cols; colIdx++ {
				colWidth := colWidths[colIdx]
				for i := 0; i < colWidth; i++ {
					buf.WriteByte(table.HeaderSeparator)
				}
				if table.AdditionalRightPadding > 0 {
					for i := 0; i < table.AdditionalRightPadding; i++ {
						buf.WriteByte(table.HeaderSeparator)
					}
				}
				if colIdx < borderedCols {
					buf.WriteString(table.ColSeparator)
				}

			}
			buf.WriteString("\n")
		}

	}

	return buf.String()
}
