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
)

type alarm struct {
	time    int64
	Playing bool
	snooze  int64
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
	tmpl, err := template.New("alarm").Parse(clock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Came herererererere")
	for _, v := range alarms {
		fmt.Println(v.Playing, v.time)
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
	w.Write([]byte("Alarm SuccessfullySet"))
	r.ParseForm()
	d := strings.Split(r.FormValue("date"), "/")
	t := strings.Split(r.FormValue("time"), ":")
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	timeStamp := time.Date(convertor(d[2]), time.Month(convertor(d[1])), convertor(d[0]), convertor(t[0]), convertor(t[1]), 0, 0, loc)
	fmt.Println(timeStamp)
	alarms = append(alarms, &alarm{timeStamp.Unix(), false, 0})
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
	lr = lrserver.New("Terminator", lrserver.DefaultPort)
	lr.SetErrorLog(nil)
	lr.SetStatusLog(nil)
	go lr.ListenAndServe()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	http.HandleFunc("/", alarmHandler)
	http.HandleFunc("/set", setAlarmHandler)
	http.HandleFunc(fmt.Sprintf("/%s", *soundFile), soundHandler)
	log.Println("Starting Server at", addr)
	go checkForAlarm()
	lr.Reload("")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Error in starting the server", err)
	}

}
