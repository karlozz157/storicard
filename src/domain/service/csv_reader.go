package service

import (
	"encoding/csv"
	"io"
	"strconv"
	"time"

	"github.com/karlozz157/storicard/src/domain/entity"
	e "github.com/karlozz157/storicard/src/domain/errors"
	"github.com/karlozz157/storicard/src/utils"
	"go.uber.org/zap"
)

const (
	dateIndex   = 0
	amountIndex = 1
)

type CsvTransactionReader struct {
	logger *zap.SugaredLogger
}

func NewCsvTransactionReader() *CsvTransactionReader {
	return &CsvTransactionReader{
		logger: utils.GetLogger(),
	}
}

func (r *CsvTransactionReader) GetTransactions(data io.Reader) ([]*entity.Transaction, error) {
	csvReader := csv.NewReader(data)

	rows, err := csvReader.ReadAll()

	if err != nil {
		r.logger.Errorw("reading csv", "err", err)
		return nil, e.ErrInternal
	}

	var transactions []*entity.Transaction

	for index, row := range rows {
		if index == 0 {
			continue
		}

		transaction, err := r.transform(row)
		if err != nil {
			r.logger.Errorw("transform row in transaction", "err", err)
			continue
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r CsvTransactionReader) transform(row []string) (*entity.Transaction, error) {
	date, err := time.Parse("1/2", row[dateIndex])
	if err != nil {
		return nil, err
	}

	amount, err := strconv.ParseFloat(row[amountIndex], 64)
	if err != nil {
		return nil, err
	}

	transaction := &entity.Transaction{
		Date:   date,
		Amount: amount,
	}

	return transaction, nil
}
