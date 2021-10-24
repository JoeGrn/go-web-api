package handler

import (
	"context"
	"fmt"
)

type Event struct {
	Name string `json:"name"`
}

type Response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

func Handle(ctx context.Context, event Event) (Response, error) {
	responseBody := fmt.Sprintf("Hello %s!", event.Name)
	response := Response{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       responseBody,
	}

	return response, nil
}
