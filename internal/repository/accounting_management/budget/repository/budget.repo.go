package budget_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	budgetsdomain "shop_erp_mono/internal/domain/accounting_management/budgets"
	"time"
)

type budgetRepository struct {
	database         *mongo.Database
	collectionBudget string
}

func NewBudgetRepository(database *mongo.Database, collectionBudget string) budgetsdomain.IBudgetRepository {
	return &budgetRepository{database: database, collectionBudget: collectionBudget}
}

func (b *budgetRepository) CreateOne(ctx context.Context, budget *budgetsdomain.Budget) error {
	collectionBudget := b.database.Collection(b.collectionBudget)

	_, err := collectionBudget.InsertOne(ctx, budget)
	if err != nil {
		return err
	}
	return nil
}

func (b *budgetRepository) GetByID(ctx context.Context, id primitive.ObjectID) (budgetsdomain.Budget, error) {
	collectionBudget := b.database.Collection(b.collectionBudget)

	filter := bson.M{"_id": id}
	var budget budgetsdomain.Budget
	if err := collectionBudget.FindOne(ctx, filter).Decode(&budget); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return budgetsdomain.Budget{}, nil
		}
		return budgetsdomain.Budget{}, err
	}

	return budget, nil
}

func (b *budgetRepository) GetByName(ctx context.Context, name string) (budgetsdomain.Budget, error) {
	collectionBudget := b.database.Collection(b.collectionBudget)

	filter := bson.M{"name": name}
	var budget budgetsdomain.Budget
	if err := collectionBudget.FindOne(ctx, filter).Decode(&budget); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return budgetsdomain.Budget{}, nil
		}
		return budgetsdomain.Budget{}, err
	}

	return budget, nil
}

func (b *budgetRepository) UpdateOne(ctx context.Context, budget *budgetsdomain.Budget) error {
	collectionBudget := b.database.Collection(b.collectionBudget)

	filter := bson.M{"_id": budget.ID}
	update := bson.M{"$set": bson.M{
		"budget_name": budget.BudgetName,
		"amount":      budget.Amount,
		"start_date":  budget.StartDate,
		"end_date":    budget.EndDate,
		"category_id": budget.CategoryID,
		"updated_at":  time.Now(),
	}}

	_, err := collectionBudget.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (b *budgetRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	collectionBudget := b.database.Collection(b.collectionBudget)

	filter := bson.M{"_id": id}
	_, err := collectionBudget.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (b *budgetRepository) GetAll(ctx context.Context) ([]budgetsdomain.Budget, error) {
	collectionBudget := b.database.Collection(b.collectionBudget)

	filter := bson.M{}
	cursor, err := collectionBudget.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var budgets []budgetsdomain.Budget
	budgets = make([]budgetsdomain.Budget, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var budget budgetsdomain.Budget
		if err = cursor.Decode(&budget); err != nil {
			return nil, err
		}

		budgets = append(budgets, budget)
	}

	return budgets, nil
}

func (b *budgetRepository) GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]budgetsdomain.Budget, error) {
	collectionBudget := b.database.Collection(b.collectionBudget)

	filter := bson.M{
		"start_date": bson.M{
			"$gte": startDate,
		},
		"end_date": bson.M{
			"$lte": endDate,
		},
	}

	cursor, err := collectionBudget.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	defer func() {
		if err = cursor.Close(ctx); err != nil {
			return
		}
	}()

	var budgets []budgetsdomain.Budget
	budgets = make([]budgetsdomain.Budget, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var budget budgetsdomain.Budget
		if err = cursor.Decode(&budget); err != nil {
			return nil, err
		}
		budgets = append(budgets, budget)
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return budgets, nil
}

func (b *budgetRepository) GetTotalAmount(ctx context.Context) (float64, error) {
	collectionBudget := b.database.Collection(b.collectionBudget)

	filter := bson.M{}
	cursor, err := collectionBudget.Find(ctx, filter)
	if err != nil {
		return 0, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var budgets float64
	for cursor.Next(ctx) {
		var budget budgetsdomain.Budget
		if err = cursor.Decode(&budget); err != nil {
			return 0, err
		}

		budgets = budgets + float64(budget.Amount)
	}

	return budgets, nil
}
