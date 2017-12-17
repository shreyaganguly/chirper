package handlers

import (
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/shreyaganguly/chirper/views"
	"github.com/unrolled/render"
)

type alarm struct {
	DateTime  string
	TimeStamp int64
	Playing   bool
	snooze    int64
	Purpose   string
}

var (
	alarms []*alarm
)

//ChirperHandler displays the page for alarm and time
func ChirperHandler(w http.ResponseWriter, r *http.Request) {
	var playing, alarmPlaying bool
	var timestamp, alarmTimestamp int64
	tmpl, err := template.New("alarm").Parse(views.Chirper)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, v := range alarms {
		if v.Playing {
			playing = true
			timestamp = v.TimeStamp
			break
		}
	}
	for _, v := range alarmClocks {
		if v.Playing {
			alarmPlaying = true
			alarmTimestamp = v.Timestamp
			break
		}
	}
	data := struct {
		Host           string
		Alarms         []*alarm
		AlarmClocks    []*alarmClock
		Playing        bool
		AlarmPlaying   bool
		TimeStamp      int64
		AlarmTimeStamp int64
		SnoozeInterval int
	}{
		host,
		alarms,
		alarmClocks,
		playing,
		alarmPlaying,
		timestamp,
		alarmTimestamp,
		snoozeTime,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//SetAlarmHandler sets the alarm at the given date and time
func SetAlarmHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dateTime := strings.Split(r.FormValue("datetime"), " ")
	purpose := r.FormValue("purpose")
	if reminderExists(r.FormValue("datetime")) {
		http.Error(w, "Reminder Exists", http.StatusBadRequest)
		return
	}
	dateParts := strings.Split(dateTime[0], "/")
	timeParts := strings.Split(dateTime[1], ":")
	hours := convertor(timeParts[0])
	if dateTime[2] == "PM" && hours != 12 {
		hours += 12
	}
	if dateTime[2] == "AM" && hours == 12 {
		hours = 0
	}
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	timeStamp := time.Date(convertor(dateParts[2]), time.Month(convertor(dateParts[1])), convertor(dateParts[0]), hours, convertor(timeParts[1]), 0, 0, loc)
	alarms = append(alarms, &alarm{r.FormValue("datetime"), timeStamp.Unix(), false, 0, purpose})
	rndr := render.New()
	rndr.HTML(w, http.StatusOK, "alarms", alarms)
}

//DeleteReminderHandler deletes an already set reminder
func DeleteReminderHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	removeReminder(int64(convertor(r.FormValue("timestamp"))))
	rndr := render.New()
	rndr.HTML(w, http.StatusOK, "alarms", alarms)
}

//SnoozeReminderHandler snoozes a running reminder to after a configurable amount of time
func SnoozeReminderHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	snoozeReminder(r.FormValue("time"), int64(convertor(r.FormValue("timestamp"))))
	rndr := render.New()
	rndr.HTML(w, http.StatusOK, "alarms", alarms)
}
