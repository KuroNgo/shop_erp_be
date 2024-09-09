package budget_usecase

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	budgetsdomain "shop_erp_mono/domain/accounting_management/budgets"
	transactioncategoriesdomain "shop_erp_mono/domain/accounting_management/transaction_categories"
	"shop_erp_mono/repository/accounting_management/budget/validate"
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

func (b *budgetUseCase) CreateBudget(ctx context.Context, input *budgetsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	if err := validate.IsNilBudget(input); err != nil {
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

	return b.budgetRepository.CreateBudget(ctx, budget)
}

func (b *budgetUseCase) GetBudget(ctx context.Context, id string) (budgetsdomain.BudgetResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	budgetID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return budgetsdomain.BudgetResponse{}, err
	}

	budgetData, err := b.budgetRepository.GetBudgetByID(ctx, budgetID)
	if err != nil {
		return budgetsdomain.BudgetResponse{}, err
	}

	budgetResponse := budgetsdomain.BudgetResponse{
		ID:         budgetData.ID,
		BudgetName: budgetData.BudgetName,
		CategoryID: budgetData.CategoryID,
		Amount:     budgetData.Amount,
		StartDate:  budgetData.StartDate,
		EndDate:    budgetData.EndDate,
		CreatedAt:  budgetData.CreatedAt,
		UpdatedAt:  budgetData.UpdatedAt,
	}

	return budgetResponse, nil
}

func (b *budgetUseCase) GetBudgetByName(ctx context.Context, name string) (budgetsdomain.BudgetResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	if name == "" {
		return budgetsdomain.BudgetResponse{}, nil
	}

	budgetData, err := b.budgetRepository.GetBudgetByName(ctx, name)
	if err != nil {
		return budgetsdomain.BudgetResponse{}, err
	}

	budgetResponse := budgetsdomain.BudgetResponse{
		ID:         budgetData.ID,
		BudgetName: budgetData.BudgetName,
		CategoryID: budgetData.CategoryID,
		Amount:     budgetData.Amount,
		StartDate:  budgetData.StartDate,
		EndDate:    budgetData.EndDate,
		CreatedAt:  budgetData.CreatedAt,
		UpdatedAt:  budgetData.UpdatedAt,
	}

	return budgetResponse, nil
}

func (b *budgetUseCase) UpdateBudget(ctx context.Context, id string, input *budgetsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	budgetID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if err = validate.IsNilBudget(input); err != nil {
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

	return b.budgetRepository.UpdateBudget(ctx, budget)
}

func (b *budgetUseCase) DeleteBudget(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	budgetID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return b.budgetRepository.DeleteBudget(ctx, budgetID)
}

func (b *budgetUseCase) ListBudgets(ctx context.Context) ([]budgetsdomain.BudgetResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	budgetData, err := b.budgetRepository.ListBudgets(ctx)
	if err != nil {
		return nil, err
	}

	var budgetResponses []budgetsdomain.BudgetResponse
	budgetResponses = make([]budgetsdomain.BudgetResponse, 0, len(budgetData))
	for _, budget := range budgetData {
		budgetResponse := budgetsdomain.BudgetResponse{
			ID:         budget.ID,
			BudgetName: budget.BudgetName,
			CategoryID: budget.CategoryID,
			Amount:     budget.Amount,
			StartDate:  budget.StartDate,
			EndDate:    budget.EndDate,
			CreatedAt:  budget.CreatedAt,
			UpdatedAt:  budget.UpdatedAt,
		}

		budgetResponses = append(budgetResponses, budgetResponse)
	}

	return budgetResponses, nil
}

func (b *budgetUseCase) GetBudgetsByDateRange(ctx context.Context, startDate, endDate string) ([]budgetsdomain.BudgetResponse, error) {
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

	budgetData, err := b.budgetRepository.GetBudgetsByDateRange(ctx, startDated, endDated)
	if err != nil {
		return nil, err
	}

	var budgetResponses []budgetsdomain.BudgetResponse
	budgetResponses = make([]budgetsdomain.BudgetResponse, 0, len(budgetData))
	for _, budget := range budgetData {
		budgetResponse := budgetsdomain.BudgetResponse{
			ID:         budget.ID,
			BudgetName: budget.BudgetName,
			CategoryID: budget.CategoryID,
			Amount:     budget.Amount,
			StartDate:  budget.StartDate,
			EndDate:    budget.EndDate,
			CreatedAt:  budget.CreatedAt,
			UpdatedAt:  budget.UpdatedAt,
		}

		budgetResponses = append(budgetResponses, budgetResponse)
	}

	return budgetResponses, nil
}

func (b *budgetUseCase) GetTotalBudgetAmount(ctx context.Context) (float64, error) {
	//TODO implement me
	panic("implement me")
}
