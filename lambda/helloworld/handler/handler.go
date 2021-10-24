package handler

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func Handle(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Processing Lambda request %s\n", event.RequestContext.RequestID)

	if len(event.Body) < 1 {
		return events.APIGatewayProxyResponse{}, errors.New("No name provided in the request body")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       "Hello " + event.Body + "!",
	}, nil
}
