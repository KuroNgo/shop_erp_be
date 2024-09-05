package account_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	accountdomain "shop_erp_mono/domain/accounting_management/account"
	"shop_erp_mono/repository/accounting_management/account/validate"
	"time"
)

type accountUseCase struct {
	contextTimeout    time.Duration
	accountRepository accountdomain.IAccountRepository
}

func NewAccountUseCase(contextTimeout time.Duration, accountRepository accountdomain.IAccountRepository) accountdomain.IAccountUseCase {
	return &accountUseCase{contextTimeout: contextTimeout, accountRepository: accountRepository}
}

func (a accountUseCase) CreateAccount(ctx context.Context, input *accountdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	if err := validate.IsNilAccount(input); err != nil {
		return err
	}

	account := &accountdomain.Accounts{
		AccountID:     primitive.NewObjectID(),
		AccountName:   input.AccountName,
		AccountNumber: input.AccountNumber,
		Balance:       input.Balance,
		AccountType:   input.AccountType,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return a.accountRepository.CreateAccount(ctx, account)
}

func (a accountUseCase) GetAccountByID(ctx context.Context, id string) (accountdomain.AccountResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	accountID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return accountdomain.AccountResponse{}, err
	}

	accountData, err := a.accountRepository.GetAccountByID(ctx, accountID)
	if err != nil {
		return accountdomain.AccountResponse{}, err
	}

	response := accountdomain.AccountResponse{
		Accounts: accountData,
	}

	return response, nil
}

func (a accountUseCase) GetAccountByName(ctx context.Context, name string) (accountdomain.AccountResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	accountData, err := a.accountRepository.GetAccountByName(ctx, name)
	if err != nil {
		return accountdomain.AccountResponse{}, err
	}

	response := accountdomain.AccountResponse{
		Accounts: accountData,
	}

	return response, nil
}

func (a accountUseCase) UpdateAccount(ctx context.Context, id string, input *accountdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	accountID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	account := &accountdomain.Accounts{
		AccountID:     accountID,
		AccountName:   input.AccountName,
		AccountNumber: input.AccountNumber,
		Balance:       input.Balance,
		AccountType:   input.AccountType,
		UpdatedAt:     time.Now(),
	}

	return a.accountRepository.UpdateAccount(ctx, account)
}

func (a accountUseCase) DeleteAccount(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	accountID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return a.accountRepository.DeleteAccount(ctx, accountID)
}

func (a accountUseCase) ListAccounts(ctx context.Context) ([]accountdomain.AccountResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	accountData, err := a.accountRepository.ListAccounts(ctx)
	if err != nil {
		return nil, err
	}

	var responses []accountdomain.AccountResponse
	responses = make([]accountdomain.AccountResponse, 0, len(accountData))
	for _, account := range accountData {
		response := accountdomain.AccountResponse{
			Accounts: account,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (a accountUseCase) GetAccountsByDateRange(ctx context.Context, startDate, endDate string) ([]accountdomain.AccountResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountUseCase) GetTotalAccountBalance(ctx context.Context) (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountUseCase) DeactivateAccount(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (a accountUseCase) ReactivateAccount(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
