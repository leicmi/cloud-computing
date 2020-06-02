package cmd

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jedib0t/go-pretty/table"
	"github.com/leicmi/cloud-computing/util"
)

func list(sess *session.Session, bucket string) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	}

	result, err := util.ListBucket(sess, input)
	if err != nil {
		fmt.Errorf("error listing bucket: %s\n", err)
	}

	util.PrintTable(result, table.Row{"jobname", "status", "uploaded"}, func(o *s3.Object) table.Row {
		key := *o.Key
		split := strings.Index(key, "_")

		return table.Row{key[split+1:], key[:split], o.LastModified.Local()}
	})
}
