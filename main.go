package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/shreyaganguly/chirper/handlers"

	"github.com/jaschaephraim/lrserver"
)

var (
	host       = flag.String("b", "0.0.0.0", "Hostname or IP address to start")
	port       = flag.Int("p", 8080, "Bind port for public web")
	soundFile  = flag.String("s", "alarm_timer.mp3", "Sound for the alarm")
	snoozeTime = flag.Int("snooze-interval", 5, "Snooze Interval of alarms in minutes")
)

func main() {
	flag.Parse()
	lr := lrserver.New("chirper", lrserver.DefaultPort)
	lr.SetErrorLog(nil)
	lr.SetStatusLog(nil)
	go lr.ListenAndServe()
	handlers.SetOptions(*soundFile, *snoozeTime, lr)
	addr := fmt.Sprintf("%s:%d", *host, *port)
	http.HandleFunc("/", handlers.ChirperHandler)
	http.HandleFunc("/set", handlers.SetAlarmHandler)
	http.HandleFunc("/snooze", handlers.SnoozeAlarmHandler)
	http.HandleFunc("/delete", handlers.DeleteAlarmHandler)
	http.HandleFunc(fmt.Sprintf("/%s", *soundFile), handlers.SoundHandler)
	http.HandleFunc("/assets/", handlers.AssetHandler)
	log.Println("Starting Server at", addr)
	go handlers.CheckForAlarm()
	lr.Reload("")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Error in starting the server", err)
	}

}
