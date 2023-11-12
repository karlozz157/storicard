package aws

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "hello world",
	}, nil
}
