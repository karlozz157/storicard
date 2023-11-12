package application

import (
	"context"
	"io"
	"net/http"

	"github.com/karlozz157/storicard/src/domain/dto"
	"github.com/karlozz157/storicard/src/domain/service"
	"github.com/karlozz157/storicard/src/infrastructure/repository"
	"github.com/karlozz157/storicard/src/utils"
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

func (h *TransactionHandler) CreateSummary(ctx context.Context, email string, body io.Reader) (*dto.Response, error) {
	if err := utils.ValidateEmail(email); err != nil {
		return nil, err
	}

	csvTransactionReader := service.NewCsvTransactionReader(email)
	transactions, err := csvTransactionReader.GetTransactions(body)
	if err != nil {
		return nil, err
	}

	if err := h.transactionService.CreateTransactions(ctx, transactions); err != nil {
		return nil, err
	}

	if _, err := h.transactionService.GetSummary(ctx, email); err != nil {
		return nil, err
	}

	return &dto.Response{
		Message:    "ok",
		StatusCode: http.StatusOK,
	}, nil
}
