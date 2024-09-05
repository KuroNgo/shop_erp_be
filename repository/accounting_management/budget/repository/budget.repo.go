package budget_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	budgetsdomain "shop_erp_mono/domain/accounting_management/budgets"
	"time"
)

type budgetRepository struct {
	database         *mongo.Database
	collectionBudget string
}

func NewBudgetRepository(database *mongo.Database, collectionBudget string) budgetsdomain.IBudgetRepository {
	return &budgetRepository{database: database, collectionBudget: collectionBudget}
}

func (b budgetRepository) CreateBudget(ctx context.Context, budget *budgetsdomain.Budget) error {
	collectionBudget := b.database.Collection(b.collectionBudget)

	_, err := collectionBudget.InsertOne(ctx, budget)
	if err != nil {
		return err
	}
	return nil
}

func (b budgetRepository) GetBudgetByID(ctx context.Context, id primitive.ObjectID) (budgetsdomain.Budget, error) {
	collectionBudget := b.database.Collection(b.collectionBudget)

	filter := bson.M{"_id": id}
	var budget budgetsdomain.Budget
	if err := collectionBudget.FindOne(ctx, filter).Decode(&budget); err != nil {
		return budgetsdomain.Budget{}, err
	}

	return budget, nil
}

func (b budgetRepository) GetBudgetByName(ctx context.Context, name string) (budgetsdomain.Budget, error) {
	collectionBudget := b.database.Collection(b.collectionBudget)

	filter := bson.M{"name": name}
	var budget budgetsdomain.Budget
	if err := collectionBudget.FindOne(ctx, filter).Decode(&budget); err != nil {
		return budgetsdomain.Budget{}, err
	}

	return budget, nil
}

func (b budgetRepository) UpdateBudget(ctx context.Context, budget *budgetsdomain.Budget) error {
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

func (b budgetRepository) DeleteBudget(ctx context.Context, id primitive.ObjectID) error {
	collectionBudget := b.database.Collection(b.collectionBudget)

	filter := bson.M{"_id": id}
	_, err := collectionBudget.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (b budgetRepository) ListBudgets(ctx context.Context) ([]budgetsdomain.Budget, error) {
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

func (b budgetRepository) GetBudgetsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]budgetsdomain.Budget, error) {
	//TODO implement me
	panic("implement me")
}

func (b budgetRepository) GetTotalBudgetAmount(ctx context.Context) (float64, error) {
	//TODO implement me
	panic("implement me")
}
