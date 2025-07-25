package scheduler

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

var c *cron.Cron

func StartScheduler() {
	c = cron.New(cron.WithSeconds())

	_, err := c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("==> Tác vụ đang chạy:", time.Now())
	})

	if err != nil {
		log.Println("Cron lỗi:", err)
		return
	}

	log.Println("Scheduler đã khởi động!") // ← bạn sẽ thấy dòng này nếu cron khởi động
	c.Start()
}
