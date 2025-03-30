package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	fmt.Print("Func Exec")
	response := events.LambdaFunctionURLResponse{
		Body:       "Hello, World!",
		StatusCode: 200,
	}

	return response, nil
}
