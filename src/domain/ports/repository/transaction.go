package repository

import (
	"context"
	"time"

	"github.com/karlozz157/storicard/src/domain/entity"
)

type ITransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) error
	GetAverageCreditAmount(ctx context.Context) (float64, error)
	GetAverageDebitAmount(ctx context.Context) (float64, error)
	GetBalance(ctx context.Context) (float64, error)
	GetNumberOfTransactions(ctx context.Context) (map[time.Month]int, error)
}
