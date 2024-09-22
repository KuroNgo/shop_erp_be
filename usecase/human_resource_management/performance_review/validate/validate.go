package validate

import (
	"errors"
	performancereviewdomain "shop_erp_mono/domain/human_resource_management/performance_review"
	"time"
)

func PerformanceReviewV1(input *performancereviewdomain.Input1) error {
	if input.ReviewerEmail == "" {
		return errors.New("the performance review's information do not nil ")
	}

	if input.EmployeeEmail == "" {
		return errors.New("the performance review's information do not nil ")
	}

	if input.PerformanceScore < 0 {
		return errors.New("the performance review's information do not nil ")
	}

	if input.ReviewDate.IsZero() {
		return errors.New("the performance review's information do not nil ")
	}

	if input.ReviewDate.After(time.Now()) {
		return errors.New("review date cannot be in the future")
	}

	return nil
}

func ValidatePerformanceReviewV2(input *performancereviewdomain.Input2) error {
	if input.ReviewerID == "" {
		return errors.New("the performance review's information do not nil ")
	}

	if input.EmployeeID == "" {
		return errors.New("the performance review's information do not nil ")
	}

	if input.PerformanceScore < 0 {
		return errors.New("the performance review's information do not nil ")
	}

	if input.ReviewDate.IsZero() {
		return errors.New("the performance review's information do not nil ")
	}

	if input.ReviewDate.After(time.Now()) {
		return errors.New("review date cannot be in the future")
	}

	return nil
}
