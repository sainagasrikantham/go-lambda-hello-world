package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest - Basic Lambda request handler.
func HandleRequest(ctx context.Context) error {
	fmt.Println("Hello Go from Lambda!")
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
