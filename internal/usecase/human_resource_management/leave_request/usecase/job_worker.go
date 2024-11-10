package leave_request_usecase

import (
	"context"
	"github.com/robfig/cron"
)

func (l *leaveRequestUseCase) StartSchedulerUpdateRemainingLeaveDays() error {
	c := cron.New()
	errCh := make(chan error, 1)

	err := c.AddFunc("0 0 1 1 *", func() {
		err := l.UpdateRemainingLeaveDays(context.TODO())
		if err != nil {
			errCh <- err
			return
		}
	})

	if err != nil {
		return err
	}

	c.Start()

	select {
	case err = <-errCh:
		if err != nil {
			return err
		}
	}

	return nil
}
