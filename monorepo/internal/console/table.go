package console

import (
	"fmt"
	"strings"
)

type Table struct {
	columnHeaders []any
	columnSizes   []int
	rows          [][]any
}

func NewTable(headers ...string) *Table {
	columnHeaders := make([]any, len(headers))
	columnSizes := make([]int, len(headers))
	for index, header := range headers {
		columnHeaders[index] = header
		columnSizes[index] = len(header)
	}
	rows := make([][]any, 0, 10)
	return &Table{
		columnHeaders,
		columnSizes,
		rows,
	}
}

func (t *Table) AddRow(values ...any) {
	if len(values) > len(t.columnHeaders) {
		panic("adding row with more than the specfied number of columns.")
	}

	row := make([]any, len(t.columnHeaders))
	for index, value := range values {
		columnValue := fmt.Sprintf("%v", value)
		if len(columnValue) > t.columnSizes[index] {
			t.columnSizes[index] = len(columnValue)
		}
		row[index] = columnValue
	}
	t.rows = append(t.rows, row)
}

func (t *Table) Print() {
	var rowFormat string
	for _, columnSize := range t.columnSizes {
		rowFormat = rowFormat + "%-" + fmt.Sprintf("%ds ", columnSize)
	}
	rowFormat = rowFormat + "\n"
	fmt.Printf(rowFormat, t.columnHeaders...)

	for _, columnSize := range t.columnSizes {
		fmt.Print(strings.Repeat("-", columnSize) + " ")
	}
	fmt.Println()

	for _, row := range t.rows {
		fmt.Printf(rowFormat, row...)
	}
}
