package task

import (
	"github.com/robfig/cron"
	"log"
	"sync"
	"time"
)

var (
	once sync.Once
)

func InitTask() {
	once.Do(func() {
		c := cron.New()
		err := c.AddFunc("0 0 0/2 * * ? *", taskTemplate)
		if err != nil {
			log.Printf("[system][task][InitTask] add task error, task name is:taskTemplate")
		}
		c.Start()

		log.Println("[system][task][InitTask] init task success!")
	})
}

// taskTemplate 定时任务模板
func taskTemplate() {
	time1 := time.Now()
	log.Printf("[system][task][taskTemplate] task start, current time:%s", time1)

	// do something

	time2 := time.Now()
	usedSecond := time2.Sub(time1)
	log.Printf("[system][task][taskTemplate] task end, used cesond:%s", usedSecond)
}
