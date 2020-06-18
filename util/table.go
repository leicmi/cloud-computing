package util

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"
)

func PrintTable(result []Job, header table.Row, formatFn func(Job) table.Row) {
	tw := table.NewWriter()
	tw.AppendHeader(header)
	for i := range result {
		tw.AppendRow(formatFn(result[i]))
	}

	fmt.Println(tw.Render())
}
