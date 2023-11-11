package service

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/karlozz157/storicard/src/domain/entity"
	"github.com/karlozz157/storicard/src/domain/ports/repository"
	"github.com/karlozz157/storicard/src/domain/ports/service"
	repo "github.com/karlozz157/storicard/src/infrastructure/repository"
)

type TransactionService struct {
	logger     *zap.SugaredLogger
	repository repository.ITransactionRepository
}

func InitTransactionService(db *mongo.Database, logger *zap.SugaredLogger) service.ITransactionService {
	return &TransactionService{
		repository: repo.InitTransactionRepository(db, logger),
		logger:     logger,
	}
}

func (s *TransactionService) CreateTransactions(ctx context.Context, transactions []*entity.Transaction) error {
	chanDone := make(chan bool)
	chanErr := make(chan error)

	for _, transaction := range transactions {
		go func(transaction *entity.Transaction) {
			if err := s.repository.CreateTransaction(ctx, transaction); err != nil {
				chanErr <- err
				return
			}

			chanDone <- true
		}(transaction)
	}

	for range transactions {
		select {
		case <-chanDone:
		case err := <-chanErr:
			close(chanErr)
			return err
		}
	}

	return nil
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

	s.logger.Info("summary", summary)

	return &summary, nil
}
