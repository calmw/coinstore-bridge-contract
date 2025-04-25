package task

import (
	"github.com/jasonlvhit/gocron"
	"time"
)

/// 定时任务

func ScheduleTask() {
	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	gocron.Every(1).Seconds().Do(GetBinancePrice)

	<-s.Start()
}
