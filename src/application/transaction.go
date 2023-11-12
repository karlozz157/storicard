package application

import (
	"context"
	"io"
	"net/http"

	"github.com/karlozz157/storicard/src/domain/dto"
	"github.com/karlozz157/storicard/src/domain/service"
	"github.com/karlozz157/storicard/src/infrastructure/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(db *mongo.Database) *TransactionHandler {
	repo := repository.NewTransactionRepository(db)

	return &TransactionHandler{
		transactionService: *service.NewTransactionService(repo),
	}
}

func (h *TransactionHandler) CreateSummary(ctx context.Context, body io.Reader) (*dto.Response, error) {
	csvTransactionReader := service.NewCsvTransactionReader()

	transactions, err := csvTransactionReader.GetTransactions(body)
	if err != nil {
		return nil, err
	}

	h.transactionService.CreateTransactions(ctx, transactions)
	h.transactionService.GetSummary(ctx)

	return &dto.Response{
		Message:    "ok",
		StatusCode: http.StatusOK,
	}, nil
}
