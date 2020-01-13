package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	handler "github.com/RyuseiNomi/shinkansen_bot/handler"
)

func main() {
	lambda.Start(handler.Tweet)
}
