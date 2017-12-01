package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/jaschaephraim/lrserver"
	"github.com/shreyaganguly/chirper/packed"
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
	soundFile = flag.String("s", "check.mp3", "Sound for the alarm")
)

const (
	applicationJavaScript = "application/javascript"
	contentType           = "Content-Type"
	textCSS               = "text/css"
	imageSVG              = "image/svg+xml"
	imageJPG              = "image/jpeg"
	imagePNG              = "image/png"
)

var assetMap = map[string]string{
	".js":  applicationJavaScript,
	".css": textCSS,
	".svg": imageSVG,
	".jpg": imageJPG,
	".png": imagePNG,
}

//AssetHandler handles the asset and set the header content
func assetHandler(w http.ResponseWriter, r *http.Request) {
	asset := r.URL.Path[1:]

	ext := path.Ext(r.URL.Path)
	data, err := packed.Asset(asset)
	if err != nil {
		fmt.Println("Error ", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	mt, ok := assetMap[ext]
	if ok {
		w.Header().Set(contentType, mt)
	} else {
		w.Header().Set(contentType, "text/plain")
	}

	w.Header().Set("Cache-Control", "public, max-age=31536000")
	w.Write(data)
}

func alarmHandler(w http.ResponseWriter, r *http.Request) {
	var playing bool
	var timestamp int64
	tmpl, err := template.New("setalarm").Parse(clock)
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
	data := struct {
		SoundFile string
		Alarms    []*alarm
		Playing   bool
		TimeStamp int64
	}{
		*soundFile,
		alarms,
		playing,
		timestamp,
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
	alarms = append(alarms, &alarm{r.FormValue("datetime"), timeStamp.Unix(), false, 0})
	rndr := render.New()
	rndr.HTML(w, http.StatusOK, "alarms", alarms)
}

func deleteAlarmHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	removeAlarm(int64(convertor(r.FormValue("timestamp"))))
	rndr := render.New()
	rndr.HTML(w, http.StatusOK, "alarms", alarms)
}

func snoozeAlarmHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dateTime := strings.Split(r.FormValue("time"), " ")
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
	timeStamp.Add(5 * time.Minute)
	snoozeAlarm(r.FormValue("time"), int64(convertor(r.FormValue("timestamp"))))
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
	http.HandleFunc("/snooze", snoozeAlarmHandler)
	http.HandleFunc("/delete", deleteAlarmHandler)
	http.HandleFunc(fmt.Sprintf("/%s", *soundFile), soundHandler)
	http.HandleFunc("/assets/", assetHandler)
	log.Println("Starting Server at", addr)
	go checkForAlarm()
	lr.Reload("")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Error in starting the server", err)
	}

}
