package common

import (
	"fmt"
	"time"
)

type TimeLogger struct {
	StartTime time.Time
	TimeMap   map[int]int64
	Enabled   bool
}

func NewTimeLogger(enabled bool) *TimeLogger {
	return &TimeLogger{
		StartTime: time.Now(),
		TimeMap:   make(map[int]int64),
		Enabled:   enabled,
	}
}

func (t *TimeLogger) Log(id int) {
	if !t.Enabled {
		return
	}
	since := time.Since(t.StartTime).Milliseconds()
	t.TimeMap[id] += since
	fmt.Printf("id: %d, time: %dms\t", id, since)
	for i := int64(0); i < since; i++ {
		fmt.Print(".")
	}
	fmt.Println()
	t.StartTime = time.Now()
}

func (t *TimeLogger) PrintResult() {
	if !t.Enabled {
		return
	}
	fmt.Println("===================== Total time =====================")
	for id, item := range t.TimeMap {
		fmt.Printf("id: %d, time: %dms\n", id, item)
	}
}
