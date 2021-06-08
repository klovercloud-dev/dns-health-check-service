package cron

import (
	"github.com/go-redis/redis"
	"github.com/robfig/cron"
	"log"
	"main/health"
	"time"
)

func RunCron(client *redis.Client) {
	// running cron job for health check
	log.Println("Create new cron")
	c := cron.New()
	c.AddFunc("@every 5s", func() {
		health.CheckHealth(client)
	})

	// Start cron with one scheduled job
	log.Println("Start cron")
	c.Start()
	time.Sleep(2 * time.Minute)
}
