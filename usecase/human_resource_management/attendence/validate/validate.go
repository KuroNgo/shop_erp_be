package validate

import (
	"errors"
	attendancedomain "shop_erp_mono/domain/human_resource_management/attendance"
)

func Attendance(input *attendancedomain.Input) error {
	if input == nil {
		return errors.New("the attendance's information do not null")
	}

	if input.CheckInTime.Before(input.CheckOutTime) {
		return errors.New("check-out time cannot be before check-in time")
	}

	return nil
}
