package validate

import (
	"errors"
	leave_request_domain "shop_erp_mono/internal/domain/human_resource_management/leave_request"
)

func LeaveRequest(input *leave_request_domain.Input) error {
	if input.EmployeeEmail == "" {
		return errors.New("email employee do not nil")
	}

	if input.ApprovesEmail == "" {
		return errors.New("email approves do not nil")
	}

	if input.LeaveType == "" {
		return errors.New("LeaveType do not nil")
	}

	if input.Status == "" {
		return errors.New("status do not nil")
	}

	if input.StartDate.IsZero() {
		return errors.New("StartDate do not nil")
	}

	if input.EndDate.IsZero() {
		return errors.New("EndDate do not nil")
	}

	if input.StartDate.Before(input.EndDate) {
		return errors.New("EndDate can not be before startDate")
	}
	return nil
}
