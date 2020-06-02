package cmd

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sirupsen/logrus"
)

func start(sess *session.Session, filename string, bucket string) {
	jobName, err := upload(sess, filename, bucket)
	if err != nil {
		logrus.WithError(err).Fatal("unable to upload file")
	}

	fmt.Printf("Queued job with name: %s\n", jobName)
}

func upload(sess *session.Session, filename string, bucket string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file %q, %v", filename, err)
	}
	jobName := fmt.Sprintf("pending_%s", filename)

	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(jobName),
		Body:   f,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}

	return jobName, nil
}
