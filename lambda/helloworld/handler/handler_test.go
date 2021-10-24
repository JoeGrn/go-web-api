package handler

import (
	"context"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	testCases := []struct {
		desc     string
		request  events.APIGatewayProxyRequest
		response events.APIGatewayProxyResponse
		err      string
	}{
		{
			desc:    "Success",
			request: events.APIGatewayProxyRequest{Body: "World"},
			response: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    map[string]string{"Content-Type": "application/json"},
				Body:       "Hello World!",
			},
			err: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := Handle(context.Background(), tC.request)
			if err != nil && tC.err == "" {
				t.Fatalf("unexpected error: %v", err)
			}
			if err != nil && err.Error() != tC.err {
				t.Errorf("expected error of %s but got %s", tC.err, err.Error())
			}
			if err == nil && tC.err != "" {
				t.Errorf("expected error of %v but got none", tC.err)
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tC.response) {
				t.Errorf("Expected account of \n%+v but got \n%+v", tC.response, got)
			}
		})
	}
}
