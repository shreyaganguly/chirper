package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jaschaephraim/lrserver"
	"github.com/unrolled/render"
)

type alarm struct {
	DateTime  string
	TimeStamp int64
	Playing   bool
	snooze    int64
}

var (
	alarms []*alarm
	lr     *lrserver.Server
)

var (
	host      = flag.String("b", "0.0.0.0", "Hostname or IP address to start")
	port      = flag.Int("p", 8080, "Bind port for public web")
	soundFile = flag.String("s", "check.mp3", "Sound fot the alarm")
)

func alarmHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("setalarm").Parse(clock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Came here in index handler")
	for _, v := range alarms {
		fmt.Println(v)
	}
	data := struct {
		SoundFile string
		Alarms    []*alarm
	}{
		*soundFile,
		alarms,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func setAlarmHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dateTime := strings.Split(r.FormValue("datetime"), " ")
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
	// fmt.Println(timeStamp)
	alarms = append(alarms, &alarm{r.FormValue("datetime"), timeStamp.Unix(), false, 0})
	rndr := render.New()
	rndr.HTML(w, http.StatusOK, "alarms", alarms)
	// err = tmpl.Execute(w, alarms)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteAlarmHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	removeAlarm(int64(convertor(r.FormValue("timestamp"))))
	rndr := render.New()
	rndr.HTML(w, http.StatusOK, "alarms", alarms)
}

func soundHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(*soundFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	io.Copy(w, file)
}

func main() {
	flag.Parse()
	lr = lrserver.New("chirper", lrserver.DefaultPort)
	lr.SetErrorLog(nil)
	lr.SetStatusLog(nil)
	go lr.ListenAndServe()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	http.HandleFunc("/", alarmHandler)
	http.HandleFunc("/set", setAlarmHandler)
	http.HandleFunc("/delete", deleteAlarmHandler)
	http.HandleFunc(fmt.Sprintf("/%s", *soundFile), soundHandler)
	log.Println("Starting Server at", addr)
	go checkForAlarm()
	lr.Reload("")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Error in starting the server", err)
	}

}
