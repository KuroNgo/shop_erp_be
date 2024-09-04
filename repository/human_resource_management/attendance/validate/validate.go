package validate

import (
	"errors"
	attendancedomain "shop_erp_mono/domain/human_resource_management/attendance"
)

func IsNilAttendance(input *attendancedomain.Input) error {
	if input == nil {
		return errors.New("the attendance's information do not null")
	}
	return nil
}
