package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response struct to enforce JSON format
type Response struct {
	Message string `json:"message"`
}

// Lambda handler function
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
	return jsonResponse("Welcome to the Home Page!")
}

func helloHandler() (events.LambdaFunctionURLResponse, error) {
	return jsonResponse("Hello, World!")
}

func byeHandler() (events.LambdaFunctionURLResponse, error) {
	return jsonResponse("Goodbye!")
}

func notFoundHandler() (events.LambdaFunctionURLResponse, error) {
	return jsonResponse("404 Not Found", http.StatusNotFound)
}

// jsonResponse helper function
func jsonResponse(message string, statusCode ...int) (events.LambdaFunctionURLResponse, error) {
	// Default to HTTP 200 if no status code is provided
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	// Convert struct to JSON string
	body, err := json.Marshal(Response{Message: message})
	if err != nil {
		return events.LambdaFunctionURLResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"message": "Internal Server Error"}`,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	return events.LambdaFunctionURLResponse{
		StatusCode: code,
		Body:       string(body),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func main() {
	lambda.Start(handler)
}
