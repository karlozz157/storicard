package service

import (
	"context"

	"github.com/karlozz157/storicard/src/domain/entity"
)

type ITransactionService interface {
	CreateTransactions(ctx context.Context, transactions []*entity.Transaction) error
	GetSummary(ctx context.Context) (*entity.Summary, error)
}
