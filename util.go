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
		go checkForTimeMatch()
	}
}

func removeAlarm(t int64) {
	for i, v := range alarms {
		fmt.Printf("%#v", v)
		fmt.Println(t)
		if v.TimeStamp == t {
			alarms = append(alarms[:i], alarms[i+1:]...)
			return
		}
	}
}

func snoozeAlarm(time string, timestamp int64) {
	for _, v := range alarms {
		fmt.Printf("%#v", v)
		if v.TimeStamp == timestamp && v.Playing {
			v.TimeStamp += 5 * 60
			v.DateTime = time
			v.Playing = false
			return
		}
	}
}
