package logic

import (
	"fmt"
	"testing"
)

func TestGetStatisticsDailyNewLogic_GetStatisticsDailyNew(t *testing.T) {
	timeCountMap := make(map[string]int)
	timeCountMap["2023-8-8"]++
	fmt.Printf("%+v", timeCountMap)
}
