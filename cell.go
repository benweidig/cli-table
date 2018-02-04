package clitable

import (
	"bytes"
	"fmt"

	"github.com/mattn/go-runewidth"
)

// Cell represents a single column in a row
type cell struct {
	// Content of the cell
	content string

	// Real width of the content (not byte-size)
	width int
}

func newCell(content interface{}) *cell {
	if content == nil {
		return &cell{
			content: "",
			width:   0,
		}
	}

	contentStr := fmt.Sprintf("%v", content)
	return &cell{
		content: contentStr,
		width:   runewidth.StringWidth(contentStr),
	}
}

func (c *cell) paddedContent(colWidth int) string {
	var buf bytes.Buffer
	buf.WriteString(c.content)
	for i := 0; i < colWidth-c.width; i++ {
		buf.WriteByte(' ')
	}
	return buf.String()
}
