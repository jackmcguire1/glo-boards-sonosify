package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func handler(
	ctx context.Context,
	event events.APIGatewayProxyRequest,
) (
	resp *events.APIGatewayProxyResponse,
	err error,
) {
	resp = &events.APIGatewayProxyResponse{}
	resp.StatusCode = http.StatusOK

	return
}

func main() {
	lambda.Start(handler)
}
