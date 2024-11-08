package attendance_controller

import (
	"shop_erp_mono/internal/config"
	attendancedomain "shop_erp_mono/internal/domain/human_resource_management/attendance"
)

type AttendanceController struct {
	Database          *config.Database
	AttendanceUseCase attendancedomain.IAttendanceUseCase
}
