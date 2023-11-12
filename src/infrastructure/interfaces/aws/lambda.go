package aws

import (
	"bytes"
	"context"
	"encoding/base64"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/karlozz157/storicard/src/application"
	e "github.com/karlozz157/storicard/src/domain/errors"
	"github.com/karlozz157/storicard/src/utils"
)

var logger = utils.GetLogger()

func handler(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	logger.Info("request", request)

	body, _ := base64.StdEncoding.DecodeString(request.Body)

	handler := application.NewTransactionHandler(utils.InitMongoDb())
	res, err := handler.CreateSummary(context.Background(), "karlozz157@gmail.com", bytes.NewReader(body))

	if err != nil {
		statusCode, message := e.ParseError(err)

		return events.LambdaFunctionURLResponse{
			StatusCode: statusCode,
			Body:       message,
		}, nil
	}

	return events.LambdaFunctionURLResponse{
		StatusCode: res.StatusCode,
		Body:       res.Message,
	}, nil
}

func StartLambda() {
	lambda.Start(handler)
}
