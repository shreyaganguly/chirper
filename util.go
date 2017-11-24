package main

import (
	"fmt"
	"strconv"
	"time"
)

func convertor(s string) int {
	var a int
	a, _ = strconv.Atoi(s)
	return a
}

var now int64

func checkForTimeMatch() {
	for _, a := range alarms {
		if a.TimeStamp == now {
			a.Playing = true
			lr.Reload("")
		}
	}
}

func checkForAlarm() {
	timeInterval, _ := time.ParseDuration("1s")
	ticker := time.NewTicker(timeInterval)
	defer ticker.Stop()
	for range ticker.C {
		now = time.Now().Unix()
		fmt.Println("Came here", now, time.Now())
		go checkForTimeMatch()
	}
}
