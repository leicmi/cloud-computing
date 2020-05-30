package util

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jedib0t/go-pretty/table"
)

func PrintTable(result *s3.ListObjectsV2Output, header table.Row, formatFn func(*s3.Object) table.Row) {
	tw := table.NewWriter()
	tw.AppendHeader(header)
	for i := range result.Contents {
		tw.AppendRow(formatFn(result.Contents[i]))
	}

	fmt.Println(tw.Render())
}
