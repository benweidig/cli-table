package clitable

// Represents a row in a table
type row struct {
	// Group of cells in the row
	cells []*cell
}

// Return new row with contents as Cells
func newRow(contents ...interface{}) *row {
	r := &row{
		cells: make([]*cell, len(contents)),
	}

	for idx, content := range contents {
		cell := newCell(content)
		r.cells[idx] = cell
	}

	return r
}
