package helper_controller

import (
	"shop_erp_mono/internal/config"
	cronjob "shop_erp_mono/pkg/shared/cron"
)

type HelperController struct {
	Cr       *cronjob.CronScheduler
	Database *config.Database
}
