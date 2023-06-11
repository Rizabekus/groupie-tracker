package pkg

import (
	"fmt"
	"html/template"
	"net/http"
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
	tmp, err := template.ParseFiles("templates/form.html")
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	} else {
		tmp.Execute(w, nil)
	}
}

func ArtistPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	if r.URL.Path != "/artist-page" {
		http.NotFound(w, r)
		return
	}
	tmp, err := template.ParseFiles("templates/form.html")
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	} else {
		tmp.Execute(w, nil)
	}
}
