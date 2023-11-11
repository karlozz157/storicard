package service

import (
	"context"
	"sync"

	"github.com/karlozz157/storicard/src/domain/entity"
	"github.com/karlozz157/storicard/src/domain/ports/repository"
	"github.com/karlozz157/storicard/src/domain/ports/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type TransactionService struct {
	repository repository.ITransactionRepository
}

func InitTransactionService(db *mongo.Database, logger *zap.SugaredLogger) service.ITransactionService {
	return &TransactionService{}
}

func (s *TransactionService) GetSummary(ctx context.Context) (*entity.Summary, error) {

	summary := entity.Summary{}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		b, err := s.repository.GetBalance(ctx)
		if err == nil {
			summary.Balance = b
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		a, err := s.repository.GetAverageCreditAmount(ctx)
		if err == nil {
			summary.AverageCredit = a
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		a, err := s.repository.GetAverageDebitAmount(ctx)
		if err == nil {
			summary.AverageDebit = a
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		t, err := s.repository.GetNumberOfTransactions(ctx)
		if err == nil {
			summary.NumberOfTransactions = t
		}
	}()

	wg.Wait()

	return &summary, nil
}
