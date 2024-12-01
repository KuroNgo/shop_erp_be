package leave_request_controller

import (
	"shop_erp_mono/internal/config"
	leaverequestdomain "shop_erp_mono/internal/domain/human_resource_management/leave_request"
	"shop_erp_mono/pkg/shared/cron"
)

type LeaveRequestController struct {
	Database            *config.Database
	CronJob             *cronjob.CronScheduler
	LeaveRequestUseCase leaverequestdomain.ILeaveRequestUseCase
}
