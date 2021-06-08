package main

import (
	"main/config"
	"main/cron"
)

func main() {
	client := config.InitDb()
	cron.RunCron(client)
}


