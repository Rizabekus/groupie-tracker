package pkg

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func MenuHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	astists := GetApi()
	tmp, err := template.ParseFiles("templates/form.html")
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	} else {
		tmp.Execute(w, astists)
	}
}

func ArtistPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	artists := GetApi()
	if r.URL.Path != "/artist-page" {
		http.NotFound(w, r)
		return
	}
	url := r.URL.String()
	xurl := strings.Split(url, "id=")
	id, _ := strconv.Atoi(xurl[1])
	tmp, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	} else {
		tmp.Execute(w, artists[id-1])
	}
}
