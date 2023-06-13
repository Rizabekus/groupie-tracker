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
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		// return
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
		// fmt.Print(err)
		// w.WriteHeader(http.StatusInternalServerError)
		// w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		// return
		ErrorHandler(w, http.StatusInternalServerError)
		return
	} else {
		tmp.Execute(w, astists)
	}
}

func ArtistPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		// return
		ErrorHandler(w, http.StatusMethodNotAllowed)
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
	if url != "/info/?id="+xurl[1] {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if id > 52 || id < 1 {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	tmp, err := template.ParseFiles("templates/artist.html")

	if err != nil {
		// fmt.Print(err)
		// w.WriteHeader(http.StatusInternalServerError)
		// w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		// return
		ErrorHandler(w, http.StatusInternalServerError)
		return
	} else {
		tmp.Execute(w, artists[id-1])
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
