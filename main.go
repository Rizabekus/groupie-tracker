package main

import (
	"git/rzhampeis/groupie-tracker/pkg"
	"log"
	"net/http"
)

func main() {
	pkg.GetApi()
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", pkg.MenuHandler)
	http.HandleFunc("/artist-page", pkg.ArtistPageHandler)
	http.HandleFunc("/css/styles.css", pkg.CssHandler)

	log.Println("Server start on http://127.0.0.1:3500")
	log.Println("OK(200)")
	err := http.ListenAndServe(":3500", nil)
	log.Fatal(err)
}
