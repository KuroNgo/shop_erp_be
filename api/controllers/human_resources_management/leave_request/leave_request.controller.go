package leave_request_controller

import (
	"shop_erp_mono/bootstrap"
	leaverequestdomain "shop_erp_mono/domain/human_resource_management/leave_request"
)

type LeaveRequestController struct {
	Database            *bootstrap.Database
	LeaveRequestUseCase leaverequestdomain.ILeaveRequestUseCase
}
