package log_activity_controller

import (
	"shop_erp_mono/internal/config"
	activity_log_domain "shop_erp_mono/internal/domain/activity_log"
	userdomain "shop_erp_mono/internal/domain/human_resource_management/user"
)

type ActivityController struct {
	ActivityUseCase activity_log_domain.ILogUseCase
	UserUseCase     userdomain.IUserUseCase
	Database        *config.Database
}
