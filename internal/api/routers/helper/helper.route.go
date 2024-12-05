package helper_route

import (
	"github.com/gin-gonic/gin"
	helpercontroller "shop_erp_mono/internal/api/controllers/helper"
	"shop_erp_mono/internal/config"
	cronjob "shop_erp_mono/pkg/shared/cron"
)

func HelperRouter(env *config.Database, cron *cronjob.CronScheduler, group *gin.RouterGroup) {
	helper := helpercontroller.HelperController{
		Cr:       cron,
		Database: env,
	}

	router := group.Group("/helper")
	router.GET("/cron/count", helper.GetJobCount)
}
