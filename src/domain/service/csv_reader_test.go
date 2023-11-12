package service

import (
	"bytes"
	"encoding/csv"
	"io"
	"testing"
	"time"

	"github.com/karlozz157/storicard/src/domain/entity"
)

func TestNewCsvTransactionReader(t *testing.T) {

	tests := []struct {
		Rows     [][]string
		Expected []*entity.Transaction
	}{
		{
			Rows: [][]string{
				[]string{
					"Date", "Amount",
				},
				[]string{
					"7/15", "+60.5",
				},
				[]string{
					"7/28", "-10.3",
				},
				[]string{
					"8/2", "-20.46",
				},
				[]string{
					"8/13", "+10",
				},
			},
			Expected: []*entity.Transaction{
				{
					Date:   time.Date(0000, time.July, 15, 0, 0, 0, 0, time.UTC),
					Amount: 60.5,
				},
				{
					Date:   time.Date(0000, time.July, 28, 0, 0, 0, 0, time.UTC),
					Amount: -10.3,
				},
				{
					Date:   time.Date(0000, time.August, 2, 0, 0, 0, 0, time.UTC),
					Amount: -20.46,
				},
				{
					Date:   time.Date(0000, time.August, 13, 0, 0, 0, 0, time.UTC),
					Amount: 10,
				},
			},
		},
		{
			Rows: [][]string{
				[]string{
					"Date", "Amount",
				},
				[]string{
					"7/15", "+60.5",
				},
				[]string{
					"7/28", "dsadsada",
				},
				[]string{
					"8/2", "-20.46",
				},
			},
			Expected: []*entity.Transaction{
				{
					Date:   time.Date(0000, time.July, 15, 0, 0, 0, 0, time.UTC),
					Amount: 60.5,
				},
				{
					Date:   time.Date(0000, time.August, 2, 0, 0, 0, 0, time.UTC),
					Amount: -20.46,
				},
			},
		},
	}

	csvTransactionReader := NewCsvTransactionReader("karlozz157@gmail.com")

	for _, test := range tests {

		reader := getReaderFromRows(test.Rows)
		transactions, _ := csvTransactionReader.GetTransactions(reader)

		for i, transactionExpected := range test.Expected {

			if !(transactionExpected.Amount == transactions[i].Amount && transactionExpected.Date == transactions[i].Date) {
				t.Errorf("expected %v have %v", transactionExpected, transactions[i])
			}
		}
	}
}

func getReaderFromRows(rows [][]string) io.Reader {
	var csvData bytes.Buffer
	csvWriter := csv.NewWriter(&csvData)
	csvWriter.WriteAll(rows)
	csvWriter.Flush()

	csvWriter.Error()

	return bytes.NewReader(csvData.Bytes())
}
