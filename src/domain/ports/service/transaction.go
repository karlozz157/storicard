package service

import (
	"context"

	"github.com/karlozz157/storicard/src/domain/entity"
)

type ITransactionService interface {
	GetSummary(ctx context.Context) (*entity.Summary, error)
}
