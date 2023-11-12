package application

import (
	"context"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/karlozz157/storicard/src/domain/dto"
	ps "github.com/karlozz157/storicard/src/domain/ports/service"
	ds "github.com/karlozz157/storicard/src/domain/service"
	"github.com/karlozz157/storicard/src/infrastructure/repository"
	"github.com/karlozz157/storicard/src/infrastructure/service"
	"github.com/karlozz157/storicard/src/utils"
)

type TransactionHandler struct {
	transactionService ds.TransactionService
	noticationService  ps.INotificationService
}

func NewTransactionHandler(db *mongo.Database) *TransactionHandler {
	repo := repository.NewTransactionRepository(db)

	return &TransactionHandler{
		transactionService: *ds.NewTransactionService(repo),
		noticationService:  service.NewMailerNotification(),
	}
}

func (h *TransactionHandler) CreateSummary(ctx context.Context, email string, body io.Reader) (*dto.Response, error) {
	if err := utils.ValidateEmail(email); err != nil {
		return nil, err
	}

	csvTransactionReader := ds.NewCsvTransactionReader(email)
	transactions, err := csvTransactionReader.GetTransactions(body)
	if err != nil {
		return nil, err
	}

	if err := h.transactionService.CreateTransactions(ctx, transactions); err != nil {
		return nil, err
	}

	summary, err := h.transactionService.GetSummary(ctx, email)
	if err != nil {
		return nil, err
	}

	if err := h.noticationService.Notify(summary); err != nil {
		return nil, err
	}

	return &dto.Response{
		Message:    "ok",
		StatusCode: http.StatusOK,
	}, nil
}
