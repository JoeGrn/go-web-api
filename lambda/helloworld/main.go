package main

import (
	"go-web-api/helloworld/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.Handle)
}
