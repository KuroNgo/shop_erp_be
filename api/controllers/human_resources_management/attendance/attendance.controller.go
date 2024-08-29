package attendance_controller

import (
	"shop_erp_mono/bootstrap"
	attendancedomain "shop_erp_mono/domain/human_resource_management/attendance"
)

type AttendanceController struct {
	Database          *bootstrap.Database
	AttendanceUseCase attendancedomain.IAttendanceUseCase
}
