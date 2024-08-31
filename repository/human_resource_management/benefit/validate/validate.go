package validate

import (
	"errors"
	benefitsdomain "shop_erp_mono/domain/human_resource_management/benefits"
)

func IsNilBenefit(input *benefitsdomain.Input) error {
	if input.EmployeeEmail == "" {
		return errors.New("email of employee do not nil")
	}

	if input.Amount <= 0 {
		return errors.New("amount do not invalid")
	}

	if input.BenefitType == "" {
		return errors.New("benefit type do not nil")
	}

	return nil
}
