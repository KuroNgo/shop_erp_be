package inventory_usecase

import cronjob "shop_erp_mono/pkg/interface/cron"

func (i *inventoryUseCase) StartSchedule(cs *cronjob.CronScheduler) {
	cs.AddCronJob("inventoryCheck", "0 8 * * *", i.CheckAndNotifyWarning)
}

func (i *inventoryUseCase) RemoveSchedules(cs *cronjob.CronScheduler) {
	err := cs.RemoveJob("inventoryCheck")
	if err != nil {
		return
	}
}