package aws

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/karlozz157/storicard/src/utils"
)

var logger = utils.GetLogger()

func handler(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	logger.Info("request", request)

	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Body:       "hello world",
	}, nil
}

func StartLambda() {
	lambda.Start(handler)
}
