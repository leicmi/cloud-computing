package util

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// FormatAWSError formats an error from calling an AWS service and formats the error with the error code
func FormatAWSError(err error) error {
	if aerr, ok := err.(awserr.Error); ok {
		return fmt.Errorf("(%s), %s", aerr.Code(), aerr.Error())
	}

	return err
}
