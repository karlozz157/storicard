package repository

import (
	"context"
	"time"

	"github.com/karlozz157/storicard/src/domain/entity"
)

type ITransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) error
	GetAverageCreditAmount(ctx context.Context, email string) (float64, error)
	GetAverageDebitAmount(ctx context.Context, email string) (float64, error)
	GetBalance(ctx context.Context, email string) (float64, error)
	GetNumberOfTransactions(ctx context.Context, email string) (map[time.Month]int, error)
}
