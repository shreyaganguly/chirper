package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/unrolled/render"
)

type alarmClock struct {
	AlarmTime string
	Timestamp int64
	Playing   bool
}

var (
	alarmClocks []*alarmClock
)

//SetHandler sets the alarm clock
func SetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var prefixedString string
	timeNow := time.Now()

	alarmTime := r.FormValue("alarmtime")
	if strings.HasSuffix(alarmTime, "AM") {
		prefixedString = "AM"
	}
	if strings.HasSuffix(alarmTime, "PM") {
		prefixedString = "PM"
	}
	timeStamp := strings.TrimSuffix(alarmTime, prefixedString)
	hour := convertor(strings.Split(timeStamp, ":")[0])
	minute := convertor(strings.Split(timeStamp, ":")[1])
	if hour == 12 && prefixedString == "AM" {
		hour = 0
	}
	if hour == 12 && prefixedString == "PM" {
		hour = 12
	}
	if prefixedString == "PM" {
		hour += 12
	}
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	alarmTimestamp := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), hour, minute, 0, 0, loc).Unix()
	if time.Now().Hour() > hour || (time.Now().Hour() == hour && time.Now().Minute() >= minute) {
		alarmTimestamp += 24 * 60 * 60
	}
	alarmClocks = append(alarmClocks, &alarmClock{alarmTime, alarmTimestamp, false})
	rndr := render.New()
	err = rndr.HTML(w, http.StatusOK, "alarmClocks", alarmClocks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//DeleteAlarmHandler deletes an already set alarm
func DeleteAlarmHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	removeAlarm(int64(convertor(r.FormValue("timestamp"))))
	rndr := render.New()
	rndr.HTML(w, http.StatusOK, "alarmClocks", alarmClocks)
}

//SnoozeAlarmHandler snoozes a running alarm to after a configurable amount of time
func SnoozeAlarmHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	snoozeAlarm(r.FormValue("time"), int64(convertor(r.FormValue("timestamp"))))
	rndr := render.New()
	rndr.HTML(w, http.StatusOK, "alarmClocks", alarmClocks)
}
