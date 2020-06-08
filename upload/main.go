package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/leicmi/cloud-computing/util"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
)

func main() {
	lambda.Start(HandleRequest)
}

func response(body string, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: body, StatusCode: statusCode}
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sess, err := session.NewSession()
	if err != nil {
		return response("unable to create new session", http.StatusInternalServerError), errors.Wrap(err, "unable to create new session")
	}

	job := &util.Job{}
	err = json.Unmarshal([]byte(req.Body), job)
	if err != nil {
		return response("unable to unmarshal body", http.StatusBadRequest), errors.Wrap(err, "unable to unmarshal body")
	}

	// Save file in S3
	err = upload(sess, job)
	if err != nil {
		return response("unable to upload to S3", http.StatusInternalServerError), errors.Wrap(err, "unable to upload job")
	}

	// Add to dynamodb
	err = addToDB(sess, job.Name)
	if err != nil {
		return response("unable to add job to db", http.StatusInternalServerError), errors.Wrap(err, "unable to queue job in database")
	}

	return response(job.Name, http.StatusOK), nil
}

func upload(sess *session.Session, job *util.Job) error {
	jobName := fmt.Sprintf("pending/%s", job.Name)

	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("lamq"),
		Key:    aws.String(jobName),
		Body:   bytes.NewReader(job.Data),
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}

	return nil
}

func addToDB(sess *session.Session, jobName string) error {
	svc := dynamodb.New(sess)
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(jobName),
			},
			"status": {
				S: aws.String(util.JOB_STATUS_PENDING),
			},
		},
		TableName: aws.String("jobs"),
	}

	result, err := svc.PutItem(input)
	if err != nil {
		return util.FormatAWSError(err)
	}

	fmt.Println(result)

	return nil
}
