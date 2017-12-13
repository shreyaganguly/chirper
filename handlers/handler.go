package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type alarmClock struct {
	AlarmTime string
	Timestamp int64
}

var (
	alarmClocks []*alarmClock
)

func SetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("*****")
	var prefixedString string
	var alarmTimestamp int64
	timeNow := time.Now()
	alarmTime := r.FormValue("alarmtime")
	fmt.Println(r.FormValue("alarmtime"))
	if strings.HasSuffix(alarmTime, "AM") {
		prefixedString = "AM"
	}
	if strings.HasSuffix(alarmTime, "PM") {
		prefixedString = "PM"
	}
	timeStamp := strings.TrimSuffix(alarmTime, prefixedString)
	fmt.Println(strings.Split(timeStamp, ":"))
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
	if time.Now().Hour() < hour {
		fmt.Println("TODAY")
		alarmTimestamp = time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), hour, minute, 0, 0, loc).Unix()
	} else if time.Now().Hour() > hour {
		fmt.Println("NEXT DAY")
		alarmTimestamp = time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), hour, minute, 0, 0, loc).Unix() + 24*60*60
	}
	if time.Now().Hour() == hour {
		if time.Now().Minute() < minute {
			fmt.Println("TODAY")
			alarmTimestamp = time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), hour, minute, 0, 0, loc).Unix()
		} else {
			fmt.Println("NEXT DAY")
			alarmTimestamp = time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), hour, minute, 0, 0, loc).Unix() + 24*60*60
		}
	}
	fmt.Println("TIME ", time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), hour, minute, 0, 0, loc).Unix(), alarmTimestamp)
	alarmClocks = append(alarmClocks, &alarmClock{alarmTime, alarmTimestamp})
	for _, v := range alarmClocks {
		fmt.Printf("%#v", v)
	}
}
