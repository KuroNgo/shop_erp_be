package validate

import (
	"errors"
	accountdomain "shop_erp_mono/domain/accounting_management/account"
)

func IsNilAccount(input *accountdomain.Input) error {
	if input.AccountName == "" {
		return errors.New("the account's information do not nil")
	}

	if input.AccountType == "" {
		return errors.New("the account's information do not nil")
	}

	if input.Balance < 0 {
		return errors.New("the account's information is invalid")
	}

	if input.AccountNumber == "" {
		return errors.New("the account's information do not nil")
	}

	return nil
}