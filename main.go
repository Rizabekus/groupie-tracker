package main

import (
	"git/rzhampeis/groupie-tracker/pkg"
	"log"
	"net/http"
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", pkg.MenuHandler)
	http.HandleFunc("/artist-page", pkg.ArtistPageHandler)
	log.Println("Server start on http://127.0.0.1:8000")
	log.Println("OK(200)")
	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)
}
