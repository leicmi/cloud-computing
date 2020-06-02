package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jedib0t/go-pretty/table"
	"github.com/leicmi/cloud-computing/util"
)

func pending(sess *session.Session, bucket string) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String("pending_"),
	}

	result, err := util.ListBucket(sess, input)
	if err != nil {
		fmt.Errorf("error listing bucket: %s\n", err)
	}

	util.PrintTable(result, table.Row{"jobname", "uploaded"}, func(o *s3.Object) table.Row {
		return table.Row{(*o.Key)[len("pending_"):], o.LastModified.Local()}
	})
}
