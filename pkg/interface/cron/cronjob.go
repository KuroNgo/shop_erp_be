package cron

import (
	"fmt"
	"github.com/robfig/cron"
)

func InitCronScheduler() *cron.Cron {
	//Create a new cron instance
	c := cron.New()

	// Add a cron job that runs every 10 seconds
	err := c.AddFunc("@weekly", func() {
		fmt.Println("Cron scheduler is running...")
	})
	if err != nil {
		return nil
	}

	c.Start()

	fmt.Println("Cron scheduler initialized")
	return c
}
