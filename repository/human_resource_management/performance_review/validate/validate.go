package validate

import (
	"errors"
	performance_review_domain "shop_erp_mono/domain/human_resource_management/performance_review"
)

func IsNilPerformanceReview(input *performance_review_domain.Input) error {
	if input.Reviewer == "" {
		return errors.New("the performance review's information do not nil ")
	}

	if input.Comments == "" {
		return errors.New("the performance review's information do not nil ")
	}

	if input.Employee == "" {
		return errors.New("the performance review's information do not nil ")
	}

	if input.PerformanceScore == 0 {
		return errors.New("the performance review's information do not nil ")
	}

	if input.ReviewDate.IsZero() {
		return errors.New("the performance review's information do not nil ")
	}

	return nil
}
