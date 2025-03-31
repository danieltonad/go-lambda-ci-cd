# AWS Lambda API in Go

## Overview

This project is a **serverless API** built with **Go** and deployed to **AWS Lambda**. It features:

- **Multiple Routes** with individual handlers.
- **Automatic Swagger Documentation** (`/docs`).
- **JSON Response Format** for all endpoints.
- **Optimized for AWS Lambda** using `GOOS=linux` and `GOARCH=arm64/x86_64` builds.

## Features

- 📡 **Lambda Function URL Support**
- 📜 **Auto-generated OpenAPI 3.0 Docs**
- ⚡ **Fast Execution with Go**
- 🛠 **Simple and Modular Codebase**

## Setup & Deployment

### Prerequisites

- Go 1.21+
- AWS CLI
- AWS Lambda configured

### Installation

1. Clone the repo:
   ```sh
   git clone https://github.com/danieltonad/go-lambda-ci-cd.git
   cd lambda-api-go
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```

### Local Testing

To test locally:

```sh
go run main.go
```

### Build for AWS Lambda

Build the Go binary for Lambda:

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bootstrap main.go
```

### Deploy to AWS Lambda

1. Zip the binary:
   ```sh
   zip deployment.zip bootstrap
   ```
2. Upload to AWS Lambda:
   ```sh
   aws lambda update-function-code --function-name YourLambdaFunction --zip-file fileb://deployment.zip
   ```

### Swagger API Docs

Once deployed, access the **Swagger UI**:

```
https://your-lambda-url/docs
```

If the `/docs` page is empty, check:

- `/docs/swagger.json` for missing or incorrect routes.
- Ensure HTTP methods are **lowercase** (`get`, `post`, etc.).

## Routes

| Method | Path     | Description     |
| ------ | -------- | --------------- |
| GET    | `/`      | Welcome Page    |
| GET    | `/hello` | Hello World API |
| GET    | `/bye`   | Goodbye Message |

## Contributing

Pull requests are welcome! Feel free to open an issue for bug fixes or improvements.

## License

MIT License
