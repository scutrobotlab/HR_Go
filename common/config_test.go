package common

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	config, err := GetConfig("./config.yaml")
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", config)
}
