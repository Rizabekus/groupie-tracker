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

		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	astists := GetApi()
	tmp, err := template.ParseFiles("templates/form.html")
	if err != nil {

		ErrorHandler(w, http.StatusInternalServerError)
		return
	} else {
		tmp.Execute(w, astists)
	}
}

func ArtistPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {

		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	artists := GetApi()

	url := r.URL.String()
	xurl := strings.Split(url, "id=")
	id, _ := strconv.Atoi(xurl[1])
	GetApi2(artists, id)
	Final := GetApi3(artists, id)

	if url != "/artist-page?id="+strconv.Itoa(id) {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if id > 52 || id < 1 {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	tmp, err := template.ParseFiles("templates/artist.html")

	if err != nil {

		ErrorHandler(w, http.StatusInternalServerError)
		return
	} else {
		tmp.Execute(w, Final)
	}
}

func ErrorHandler(w http.ResponseWriter, status int) {
	tmp, err := template.ParseFiles("./templates/error.html")
	if err != nil || tmp == nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	var Err ErrorStruct
	Err.Message = http.StatusText(status)
	Err.Status = status
	err = tmp.Execute(w, Err)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
}

func CssHandler(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, http.StatusNotFound)
	return
}
