package handlers

import (
	"io"
	"net/http"
	"os"
)

//SoundHandler servers the mp3 sound file which will serve as alarm and timer music
func SoundHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(sound)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	io.Copy(w, file)
}
