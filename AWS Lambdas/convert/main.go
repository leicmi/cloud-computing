package convert

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/leicmi/cloud-computing/util"
)

func handler(ctx context.Context, s3Event events.S3Event) {
	sess, err := session.NewSession()
	if err != nil {
		// return err // TODO
	}

	for _, record := range s3Event.Records {
		s3 := record.S3
		fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)

		addStartToDB(sess, s3.Object.Key, "STARTED", record.EventTime, time.Now()) // util.JOB_STATUS_PENDING
		// download
		time.Sleep(2 * time.Second) // do the work
		// upload result
		// urlStr, err := req.Presign(15 * time.Minute)
		addFinishToDB(sess, s3.Object.Key, "FINISHED", time.Now(), "todo: download link")

	}
}

func addStartToDB(sess *session.Session, jobName string, jobStatus string, uploadTime time.Time, startTime time.Time) error {
	svc := dynamodb.New(sess)
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(jobName),
			},
			"jobStatus": {
				S: aws.String(jobStatus),
			},
			"uploadTime": {
				S: aws.String(uploadTime.String()),
			},
			"startTime": {
				S: aws.String(startTime.String()),
			},
		},
		TableName: aws.String(os.Getenv("DynamoDBTable")),
	}

	_, err := svc.PutItem(input)
	if err != nil {
		return util.FormatAWSError(err)
	}

	return nil
}

func addFinishToDB(sess *session.Session, jobName string, jobStatus string, finishTime time.Time, downloadURI string) error {
	svc := dynamodb.New(sess)
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(jobName),
			},
			"jobStatus": {
				S: aws.String(jobStatus),
			},
			"finishTime": {
				S: aws.String(finishTime.String()),
			},
			"downloadURI": {
				S: aws.String(downloadURI),
			},
		},
		TableName: aws.String(os.Getenv("DynamoDBTable")),
	}

	_, err := svc.PutItem(input)
	if err != nil {
		return util.FormatAWSError(err)
	}

	return nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
