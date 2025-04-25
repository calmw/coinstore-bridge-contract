package task

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

/// 定时任务

func ScheduleTask() {
	s := gocron.NewScheduler()
	err := gocron.Every(1).Second().Do(GetBinancePrice)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~", err)
	<-s.Start()
}
