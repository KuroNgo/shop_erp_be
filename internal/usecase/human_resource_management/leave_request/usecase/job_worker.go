package leave_request_usecase

import cronjob "shop_erp_mono/pkg/interface/cron"

func (l *leaveRequestUseCase) StartSchedulerUpdateRemainingLeaveDays(cs *cronjob.CronScheduler) {
	cs.AddCronJob("updateRemainingLeaveDays", "0 0 1 1 *", l.UpdateRemainingLeaveDays)
}

func (l *leaveRequestUseCase) StopSchedulerUpdateRemainingLeaveDays(cs *cronjob.CronScheduler) {
	err := cs.RemoveJob("updateRemainingLeaveDays")
	if err != nil {
		return
	}
}
