package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type book struct {
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	s := &book{
		ISBN:   "978-1420931693",
		Title:  "The Republic",
		Author: "Plato",
	}

	s, err := getItem("978-0486298238")

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	bk, err := json.Marshal(s)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(bk),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
