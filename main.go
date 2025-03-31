package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Lambda handler
func handler(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	fmt.Println("Received request:", request.RequestContext.HTTP.Method, request.RawPath)

	switch request.RawPath {
	case "/":
		return homeHandler()
	case "/hello":
		return helloHandler()
	case "/bye":
		return byeHandler()
	default:
		return notFoundHandler()
	}
}

// Handlers for different routes
func homeHandler() (events.LambdaFunctionURLResponse, error) {
	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Body:       "Welcome to the Home Page!",
	}, nil
}

func helloHandler() (events.LambdaFunctionURLResponse, error) {
	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Body:       "Hello, World!",
	}, nil
}

func byeHandler() (events.LambdaFunctionURLResponse, error) {
	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Body:       "Goodbye!",
	}, nil
}

func notFoundHandler() (events.LambdaFunctionURLResponse, error) {
	return events.LambdaFunctionURLResponse{
		StatusCode: http.StatusNotFound,
		Body:       "404 Not Found",
	}, nil
}

func main() {
	lambda.Start(handler)
}
