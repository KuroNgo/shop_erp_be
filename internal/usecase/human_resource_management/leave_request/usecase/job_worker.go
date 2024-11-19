package leave_request_usecase

import cronjob "shop_erp_mono/pkg/interface/cron"

func (l *leaveRequestUseCase) StartSchedulerUpdateRemainingLeaveDays(cs *cronjob.CronScheduler) {
	cs.AddCronJob("updateRemainingLeaveDays", "0 0 1 1 *", l.UpdateRemainingLeaveDays)
}
