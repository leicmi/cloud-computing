package util

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ListBucket(sess *session.Session, input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	svc := s3.New(sess)

	result, err := svc.ListObjectsV2(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				return nil, fmt.Errorf("%s %s", s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				return nil, aerr
			}
		}

		return nil, err
	}

	return result, nil
}
