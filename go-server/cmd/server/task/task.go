package task

import (
	"coinstore/cmd/server/token"
	"github.com/jasonlvhit/gocron"
)

/// 定时任务

func ScheduleTask() {
	s := gocron.NewScheduler()
	_ = s.Every(2).Seconds().From(gocron.NextTick()).Do(token.GetBinancePrice)
	_ = s.Every(2).Seconds().From(gocron.NextTick()).Do(token.GetBybitPrice)
	<-s.Start()
}
