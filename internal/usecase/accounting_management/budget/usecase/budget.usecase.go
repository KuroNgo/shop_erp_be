package budget_usecase

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	budgetsdomain "shop_erp_mono/internal/domain/accounting_management/budgets"
	transactioncategoriesdomain "shop_erp_mono/internal/domain/accounting_management/transaction_categories"
	"shop_erp_mono/internal/usecase/accounting_management/budget/validate"
	"time"
)

type budgetUseCase struct {
	contextTimeout                time.Duration
	budgetRepository              budgetsdomain.IBudgetRepository
	transactionCategoryRepository transactioncategoriesdomain.ITransactionCategoriesRepository
}

func NewBudgetUseCase(contextTimeout time.Duration, budgetRepository budgetsdomain.IBudgetRepository,
	transactionCategoryRepository transactioncategoriesdomain.ITransactionCategoriesRepository) budgetsdomain.IBudgetUseCase {
	return &budgetUseCase{contextTimeout: contextTimeout, budgetRepository: budgetRepository, transactionCategoryRepository: transactionCategoryRepository}
}

func (b *budgetUseCase) CreateOne(ctx context.Context, input *budgetsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	if err := validate.Budget(input); err != nil {
		return err
	}

	categoryData, err := b.transactionCategoryRepository.GetByName(ctx, input.BudgetName)
	if err != nil {
		return err
	}

	budget := &budgetsdomain.Budget{
		ID:         primitive.NewObjectID(),
		BudgetName: input.BudgetName,
		CategoryID: categoryData.ID,
		Amount:     input.Amount,
		StartDate:  input.StartDate,
		EndDate:    input.EndDate,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return b.budgetRepository.CreateOne(ctx, budget)
}

func (b *budgetUseCase) GetByID(ctx context.Context, id string) (budgetsdomain.BudgetResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	budgetID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return budgetsdomain.BudgetResponse{}, err
	}

	budgetData, err := b.budgetRepository.GetByID(ctx, budgetID)
	if err != nil {
		return budgetsdomain.BudgetResponse{}, err
	}

	budgetResponse := budgetsdomain.BudgetResponse{
		Budget: budgetData,
	}

	return budgetResponse, nil
}

func (b *budgetUseCase) GetByName(ctx context.Context, name string) (budgetsdomain.BudgetResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	if name == "" {
		return budgetsdomain.BudgetResponse{}, nil
	}

	budgetData, err := b.budgetRepository.GetByName(ctx, name)
	if err != nil {
		return budgetsdomain.BudgetResponse{}, err
	}

	budgetResponse := budgetsdomain.BudgetResponse{
		Budget: budgetData,
	}

	return budgetResponse, nil
}

func (b *budgetUseCase) UpdateOne(ctx context.Context, id string, input *budgetsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	budgetID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if err = validate.Budget(input); err != nil {
		return err
	}

	categoryData, err := b.transactionCategoryRepository.GetByName(ctx, input.BudgetName)
	if err != nil {
		return err
	}

	budget := &budgetsdomain.Budget{
		ID:         budgetID,
		BudgetName: input.BudgetName,
		CategoryID: categoryData.ID,
		Amount:     input.Amount,
		StartDate:  input.StartDate,
		EndDate:    input.EndDate,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return b.budgetRepository.UpdateOne(ctx, budget)
}

func (b *budgetUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	budgetID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return b.budgetRepository.DeleteOne(ctx, budgetID)
}

func (b *budgetUseCase) GetAll(ctx context.Context) ([]budgetsdomain.BudgetResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	budgetData, err := b.budgetRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var budgetResponses []budgetsdomain.BudgetResponse
	budgetResponses = make([]budgetsdomain.BudgetResponse, 0, len(budgetData))
	for _, budget := range budgetData {
		budgetResponse := budgetsdomain.BudgetResponse{
			Budget: budget,
		}

		budgetResponses = append(budgetResponses, budgetResponse)
	}

	return budgetResponses, nil
}

func (b *budgetUseCase) GetByDateRange(ctx context.Context, startDate, endDate string) ([]budgetsdomain.BudgetResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	if startDate == "" && endDate == "" {
		return nil, errors.New("")
	}

	layout := "2006-01-02 15:04:05"
	startDated, err := time.Parse(layout, startDate)
	if err != nil {
		fmt.Println("Lỗi khi parse:", err)
	}

	endDated, err := time.Parse(layout, endDate)
	if err != nil {
		fmt.Println("Lỗi khi parse:", err)
	}

	budgetData, err := b.budgetRepository.GetByDateRange(ctx, startDated, endDated)
	if err != nil {
		return nil, err
	}

	var budgetResponses []budgetsdomain.BudgetResponse
	budgetResponses = make([]budgetsdomain.BudgetResponse, 0, len(budgetData))
	for _, budget := range budgetData {
		budgetResponse := budgetsdomain.BudgetResponse{
			Budget: budget,
		}

		budgetResponses = append(budgetResponses, budgetResponse)
	}

	return budgetResponses, nil
}

func (b *budgetUseCase) GetTotalAmount(ctx context.Context) (float64, error) {
	//TODO implement me
	panic("implement me")
}
