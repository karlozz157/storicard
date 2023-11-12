package service

import (
	"context"
	"testing"
	"time"

	"github.com/karlozz157/storicard/src/domain/entity"
	e "github.com/karlozz157/storicard/src/domain/errors"
	"github.com/karlozz157/storicard/src/utils"
)

func TestCreateTransactions(t *testing.T) {

	tests := []struct {
		Transactions []*entity.Transaction
		Expected     error
	}{
		{
			Transactions: []*entity.Transaction{
				{
					Date:   time.Now(),
					Amount: +10,
				},
				{
					Date:   time.Now(),
					Amount: -5,
				},
			},
			Expected: nil,
		},
		{
			Transactions: []*entity.Transaction{
				{
					Date:   time.Now(),
					Amount: 1,
				},
				{
					Date:   time.Now(),
					Amount: 157,
				},
			},
			Expected: e.ErrInternal,
		},
	}

	service := TransactionService{
		repository: &TransactionRepositoryMock{},
	}

	for _, test := range tests {
		err := service.CreateTransactions(context.Background(), test.Transactions)

		if err != nil {
			if test.Expected.Error() != err.Error() {
				t.Errorf("expected %v have %v", test.Expected, err)
			}
		}
	}
}

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
		logger:     utils.GetLogger(),
	}

	summary, _ := service.GetSummary(context.Background(), "karlozz157@gmail.com")

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

func (r *TransactionRepositoryMock) GetAverageCreditAmount(ctx context.Context, email string) (float64, error) {
	return 35.25, nil
}

func (r *TransactionRepositoryMock) GetAverageDebitAmount(ctx context.Context, email string) (float64, error) {
	return -15.38, nil
}

func (r *TransactionRepositoryMock) GetBalance(ctx context.Context, email string) (float64, error) {
	return 39.74, nil
}

func (r *TransactionRepositoryMock) GetNumberOfTransactions(ctx context.Context, email string) (map[time.Month]int, error) {
	numberOfTransactions := make(map[time.Month]int)
	numberOfTransactions[time.July] = 2
	numberOfTransactions[time.August] = 2

	return numberOfTransactions, nil
}

func (r *TransactionRepositoryMock) CreateTransaction(ctx context.Context, transaction *entity.Transaction) error {
	if transaction.Amount == 157 {
		return e.ErrInternal
	}

	return nil
}
