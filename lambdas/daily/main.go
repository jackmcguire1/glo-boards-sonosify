package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
	"time"
)

var (
	motd string
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	motd = os.Getenv("MOTD")
}

func handler(
	ctx context.Context,
	event events.CloudWatchEvent,
) (
	err error,
) {
	switch event.Time.Weekday() {
	case time.Saturday,
		time.Sunday:
	}

	return
}

func main() {
	lambda.Start(handler)
}
