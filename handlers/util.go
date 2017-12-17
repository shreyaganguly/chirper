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
	host       string
)

//SetOptions sets soundFile, snoozeTime and livereload server
func SetOptions(h, s string, st int, l *lrserver.Server) {
	host = h
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
			lr.Reload("#remindersection")
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
	for i, v := range alarmClocks {
		if v.Timestamp == t {
			alarmClocks = append(alarmClocks[:i], alarmClocks[i+1:]...)
			return
		}
	}
}

func removeReminder(t int64) {
	for i, v := range alarms {
		if v.TimeStamp == t {
			alarms = append(alarms[:i], alarms[i+1:]...)
			return
		}
	}
}

func snoozeAlarm(time string, timestamp int64) {
	for _, v := range alarmClocks {
		if v.Timestamp == timestamp && v.Playing {
			v.Timestamp += int64(snoozeTime) * 60
			v.AlarmTime = time
			v.Playing = false
			return
		}
	}
}

func snoozeReminder(time string, timestamp int64) {
	for _, v := range alarms {
		if v.TimeStamp == timestamp && v.Playing {
			v.TimeStamp += int64(snoozeTime) * 60
			v.DateTime = time
			v.Playing = false
			return
		}
	}
}

func reminderExists(reminderTime string) bool {
	for _, v := range alarms {
		if v.DateTime == reminderTime {
			return true
		}
	}
	return false
}
func alarmExists(alarmTime string) bool {
	for _, v := range alarmClocks {
		if v.AlarmTime == alarmTime {
			return true
		}
	}
	return false
}
