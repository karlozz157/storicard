package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/karlozz157/storicard/src/domain/entity"
	e "github.com/karlozz157/storicard/src/domain/errors"
	"github.com/karlozz157/storicard/src/domain/ports/repository"
)

const transactionCollection = "transactions"

type TransactionRepository struct {
	db     *mongo.Database
	logger *zap.SugaredLogger
}

func InitTransactionRepository(db *mongo.Database, logger *zap.SugaredLogger) repository.ITransactionRepository {
	return &TransactionRepository{
		db:     db,
		logger: logger,
	}
}

func (r *TransactionRepository) GetAverageCreditAmount(ctx context.Context) (float64, error) {
	return r.getAverage(ctx, "$gt")
}

func (r *TransactionRepository) GetAverageDebitAmount(ctx context.Context) (float64, error) {
	return r.getAverage(ctx, "$lt")
}

func (r *TransactionRepository) GetBalance(ctx context.Context) (float64, error) {

	pipeline := bson.A{
		bson.D{
			{"$group",
				bson.D{
					{"_id", primitive.Null{}},
					{"total", bson.D{{"$sum", "$amount"}}},
				},
			},
		},
	}

	return r.getTotal(ctx, pipeline)
}

func (r *TransactionRepository) GetNumberOfTransactions(ctx context.Context) (map[time.Month]int, error) {
	cursor, err := r.db.Collection(transactionCollection).Find(ctx, bson.M{})

	if err != nil {
		r.logger.Errorw("getting transctions", "error", err)
		return nil, e.ErrInternal
	}

	numberOfTransactions := make(map[time.Month]int)

	for cursor.Next(ctx) {
		var t entity.Transaction
		if err := cursor.Decode(&t); err != nil {
			r.logger.Errorw("decoding result", "error", err)
			return nil, e.ErrInternal
		}

		m := t.Date.Month()

		if _, ok := numberOfTransactions[m]; !ok {
			numberOfTransactions[m] = 1
		} else {
			numberOfTransactions[m]++
		}
	}

	return numberOfTransactions, nil
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, account *entity.Transaction) error {
	result, err := r.db.Collection(transactionCollection).InsertOne(ctx, account)

	if err != nil {
		r.logger.Errorw("creating account", "error", err)
		return e.ErrInternal
	}

	r.logger.Infow("created", "result", result)

	return nil
}

func (r *TransactionRepository) getAverage(ctx context.Context, operator string) (float64, error) {
	pipeline := bson.A{
		bson.D{{"$match", bson.D{{"amount", bson.D{{operator, 0}}}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", primitive.Null{}},
					{"total", bson.D{{"$avg", "$amount"}}},
				},
			},
		},
	}

	return r.getTotal(ctx, pipeline)
}

func (r *TransactionRepository) getTotal(ctx context.Context, pipeline bson.A) (float64, error) {
	r.logger.Info("pipeline", pipeline)

	cursor, err := r.db.Collection(transactionCollection).Aggregate(ctx, pipeline)

	if err != nil {
		r.logger.Errorw("doing aggregate", "error", err)
		return 0, e.ErrInternal
	}

	type Result struct {
		Total float64 `bson:"total"`
	}

	var result Result

	for cursor.Next(ctx) {
		err := cursor.Decode(&result)

		if err != nil {
			r.logger.Errorw("decoding result", "error", err)
			return 0, e.ErrInternal
		}
	}

	return result.Total, nil
}
