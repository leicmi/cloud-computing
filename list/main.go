package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/leicmi/cloud-computing/util"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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

	r, err := scanDB(sess)
	if err != nil {
		return response("error while executing query", http.StatusInternalServerError), errors.Wrap(err, "unable to perform scan")
	}

	jobs := []util.Job{}
	for i := range r.Items {
		jobs = append(jobs, util.Job{ID: *r.Items[i]["id"].S, Status: *r.Items[i]["jobStatus"].S})
	}

	jobsJSON, err := json.Marshal(jobs)
	if err != nil {
		return response("error while building response", http.StatusInternalServerError), errors.Wrap(err, "unable to marshal json")
	}

	return response(string(jobsJSON), http.StatusOK), nil
}

func scanDB(sess *session.Session) (*dynamodb.ScanOutput, error) {
	svc := dynamodb.New(sess)
	query := &dynamodb.ScanInput{
		ProjectionExpression: aws.String("id, jobStatus"),
		TableName:            aws.String("jobs"),
	}

	result, err := svc.Scan(query)
	if err != nil {
		return nil, util.FormatAWSError(err)
	}

	return result, nil
}
