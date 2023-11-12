package aws

import (
	"bytes"
	"context"
	"encoding/base64"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/karlozz157/storicard/src/application"
	e "github.com/karlozz157/storicard/src/domain/errors"
	"github.com/karlozz157/storicard/src/utils"
)

var logger = utils.GetLogger()

func handler(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	logger.Info("request", request)

	var email string
	queryStringSplited := strings.Split(request.RawPath, "/")
	if len(queryStringSplited) == 2 {
		email = queryStringSplited[1]
	}

	body, _ := base64.StdEncoding.DecodeString(request.Body)

	handler := application.NewTransactionHandler(utils.InitMongoDb())
	res, err := handler.CreateSummary(context.Background(), email, bytes.NewReader(body))

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
