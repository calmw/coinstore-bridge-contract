package task

import (
	"github.com/jasonlvhit/gocron"
)

/// 定时任务

func ScheduleTask() {
	s := gocron.NewScheduler()
	_ = s.Every(1).Seconds().From(gocron.NextTick()).Do(GetBinancePrice)
	<-s.Start()
}
