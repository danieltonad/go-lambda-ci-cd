package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Route struct
type Route struct {
	Path        string
	Method      string
	Summary     string
	Description string
	Handler     func() (events.LambdaFunctionURLResponse, error)
}

// Define API routes
var routes = []Route{
	{Path: "/", Method: "GET", Summary: "Home", Description: "Welcome page", Handler: homeHandler},
	{Path: "/hello", Method: "GET", Summary: "Hello", Description: "Hello world message", Handler: helloHandler},
	{Path: "/bye", Method: "GET", Summary: "Goodbye", Description: "Say goodbye", Handler: byeHandler},
}

// Auto-generate Swagger JSON
func generateSwaggerJSON() string {
	swagger := map[string]interface{}{
		"openapi": "3.0.0",
		"info": map[string]string{
			"title":       "Lambda API",
			"description": "Auto-generated API docs",
			"version":     "1.0.0",
		},
		"paths": map[string]interface{}{},
	}

	paths := swagger["paths"].(map[string]interface{})
	for _, route := range routes {
		paths[route.Path] = map[string]interface{}{
			route.Method: map[string]interface{}{
				"summary":     route.Summary,
				"description": route.Description,
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Success",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type":       "object",
									"properties": map[string]interface{}{"message": map[string]string{"type": "string"}},
								},
							},
						},
					},
				},
			},
		}
	}

	// Convert to JSON
	swaggerJSON, _ := json.MarshalIndent(swagger, "", "  ")
	return string(swaggerJSON)
}

// Route Handlers
func homeHandler() (events.LambdaFunctionURLResponse, error) {
	return jsonResponse("Welcome to Home!")
}
func helloHandler() (events.LambdaFunctionURLResponse, error) {
	return jsonResponse("Hello, World!")
}
func byeHandler() (events.LambdaFunctionURLResponse, error) {
	return jsonResponse("Goodbye!")
}

// Serve Swagger JSON
func swaggerHandler() (events.LambdaFunctionURLResponse, error) {
	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Body:       generateSwaggerJSON(),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

// Serve Swagger UI (Embedded)
func swaggerUIHandler() (events.LambdaFunctionURLResponse, error) {
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Swagger UI</title>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/5.1.0/swagger-ui.min.css" />
	</head>
	<body>
		<div id="swagger-ui"></div>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/5.1.0/swagger-ui-bundle.min.js"></script>
		<script>
			const ui = SwaggerUIBundle({
				url: "/docs/swagger.json",
				dom_id: "#swagger-ui",
			});
		</script>
	</body>
	</html>
	`
	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Body:       html,
		Headers:    map[string]string{"Content-Type": "text/html"},
	}, nil
}

// JSON Response Helper
func jsonResponse(message string) (events.LambdaFunctionURLResponse, error) {
	body, _ := json.Marshal(map[string]string{"message": message})
	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Body:       string(body),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

// Lambda Handler
func handler(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	fmt.Println("Received:", request.RequestContext.HTTP.Method, request.RawPath)

	// Match API Routes
	for _, route := range routes {
		if request.RawPath == route.Path {
			return route.Handler()
		}
	}

	// Serve Swagger JSON
	if request.RawPath == "/docs/swagger.json" {
		return swaggerHandler()
	}

	// Serve Swagger UI
	if request.RawPath == "/docs" {
		return swaggerUIHandler()
	}

	return jsonResponse("404 Not Found")
}

func main() {
	lambda.Start(handler)
}
