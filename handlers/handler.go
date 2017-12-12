package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func SetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("*****")
	var prefixedString string

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
	if time.Now().Hour() < hour {
		fmt.Println("TODAY")
	} else if time.Now().Hour() > hour {
		fmt.Println("NEXT DAY")
	}
	if time.Now().Hour() == hour {
		if time.Now().Minute() < minute {
			fmt.Println("TODAY")
		} else {
			fmt.Println("NEXT DAY")
		}
	}

}
