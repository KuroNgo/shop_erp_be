package validate

import (
	"errors"
	benefitsdomain "shop_erp_mono/domain/human_resource_management/benefits"
)

func Benefit(input *benefitsdomain.Input) error {
	if input.EmployeeEmail == "" {
		return errors.New("email of employee do not nil")
	}

	if input.Amount < 0 {
		return errors.New("amount value is invalid")
	}

	if input.BenefitType == "" {
		return errors.New("benefit type do not nil")
	}

	if input.StartDate.Before(input.EndDate) {
		return errors.New("endDate do not before startDate")
	}

	return nil
}
