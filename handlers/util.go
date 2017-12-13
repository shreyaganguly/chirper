package handlers

import (
	"strconv"
	"time"

	"github.com/jaschaephraim/lrserver"
)

var (
	lr         *lrserver.Server
	sound      string
	snoozeTime int
)

//SetOptions sets soundFile, snoozeTime and livereload server
func SetOptions(s string, st int, l *lrserver.Server) {
	sound = s
	snoozeTime = st
	lr = l
}
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

func checkForAlarmMatch() {
	for _, a := range alarmClocks {
		if a.Timestamp == now {
			a.Playing = true
			lr.Reload("")
		}
	}
}

//CheckForAlarm checks whether an alarm is configured at that time
func CheckForAlarm() {
	timeInterval, _ := time.ParseDuration("1s")
	ticker := time.NewTicker(timeInterval)
	defer ticker.Stop()
	for range ticker.C {
		now = time.Now().Unix()
		go checkForTimeMatch()
		go checkForAlarmMatch()
	}
}

func removeAlarm(t int64) {
	for i, v := range alarms {
		if v.TimeStamp == t {
			alarms = append(alarms[:i], alarms[i+1:]...)
			return
		}
	}
}

func snoozeAlarm(time string, timestamp int64) {
	for _, v := range alarms {
		if v.TimeStamp == timestamp && v.Playing {
			v.TimeStamp += int64(snoozeTime) * 60
			v.DateTime = time
			v.Playing = false
			return
		}
	}
}
