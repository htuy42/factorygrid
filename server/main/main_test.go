package main

import (
	"server/config"
	"testing"
	"time"
)

func Test_main(m *testing.T) {
	config.CONFIG_REL_DIR = "../../configs/"
	go main()
	time.Sleep(30 * time.Second)
}
