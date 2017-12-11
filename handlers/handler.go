package handlers

import (
	"fmt"
	"net/http"
)

func SetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("*****")
	fmt.Println(r.FormValue("alarmtime"))
}
