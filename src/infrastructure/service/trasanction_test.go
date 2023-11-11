package service

import (
	"context"
	"testing"
	"time"

	"github.com/karlozz157/storicard/src/domain/entity"
)

func TestGetSummary(t *testing.T) {

	test := struct {
		summary *entity.Summary
	}{
		summary: &entity.Summary{
			Balance:              39.74,
			AverageCredit:        35.25,
			AverageDebit:         -15.38,
			NumberOfTransactions: map[time.Month]int{time.July: 2, time.August: 2},
		},
	}

	service := TransactionService{
		repository: &TransactionRepositoryMock{},
	}

	summary, _ := service.GetSummary(context.Background())

	if test.summary.Balance != summary.Balance {
		t.Errorf("balance expected %.2f have %.2f", test.summary.Balance, summary.Balance)
	}

	if test.summary.AverageCredit != summary.AverageCredit {
		t.Errorf("average credit expected %.2f have %.2f", test.summary.AverageCredit, summary.AverageCredit)
	}

	if test.summary.AverageDebit != summary.AverageDebit {
		t.Errorf("average debit expected %.2f have %.2f", test.summary.AverageDebit, summary.AverageDebit)
	}

	for month, numberOfTransactions := range summary.NumberOfTransactions {
		expectedNumberOfTransactions, ok := test.summary.NumberOfTransactions[month]

		if !ok {
			t.Errorf("month %v is not present", month)
		}

		if expectedNumberOfTransactions != numberOfTransactions {
			t.Errorf("number of transaction expected %d have %d", expectedNumberOfTransactions, numberOfTransactions)
		}
	}
}

type TransactionRepositoryMock struct {
}

func (r *TransactionRepositoryMock) GetAverageCreditAmount(ctx context.Context) (float64, error) {
	return 35.25, nil
}

func (r *TransactionRepositoryMock) GetAverageDebitAmount(ctx context.Context) (float64, error) {
	return -15.38, nil
}

func (r *TransactionRepositoryMock) GetBalance(ctx context.Context) (float64, error) {
	return 39.74, nil
}

func (r *TransactionRepositoryMock) GetNumberOfTransactions(ctx context.Context) (map[time.Month]int, error) {
	numberOfTransactions := make(map[time.Month]int)
	numberOfTransactions[time.July] = 2
	numberOfTransactions[time.August] = 2

	return numberOfTransactions, nil
}

func (r *TransactionRepositoryMock) CreateTransaction(ctx context.Context, account *entity.Transaction) error {
	return nil
}
