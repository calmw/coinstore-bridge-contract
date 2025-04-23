package monitor

import (
	"coinstore/bridge/monitor/task"
	"github.com/jasonlvhit/gocron"
	"log"
	"os"
	"strconv"
)

func Start() {
	task.InitTask()
	go task.NewMonitor().ProcessFailedOrder()

	intervalStr := os.Getenv("FAILED_JOB_INTERVAL")
	if len(intervalStr) == 0 {
		intervalStr = "300"
	}
	interval, err := strconv.ParseInt(intervalStr, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}
	if interval == 0 {
		interval = 30
	}
	s := gocron.NewScheduler()
	_ = s.Every(uint64(interval)).Seconds().From(gocron.NextTick()).Do(task.FailedTask)
	<-s.Start()
}
