package cronjob

import (
	"fmt"
	"time"
)

type TestJob struct {
	BaseJob
}

func (t *TestJob) Run() {
	fmt.Println(time.Now().Format(time.DateTime) + ": --- test cronjob ---")
}
