package task

import (
	"github.com/jasonlvhit/gocron"
	"time"
)

/// 定时任务

func ScheduleTask() {
	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	gocron.Every(1).Day().At("00:00:00").Do(task) // 更新relayer账户充值记录
	gocron.Every(1).Day().At("00:05:00").Do(task) // 更新财务统计

	<-s.Start()
}
