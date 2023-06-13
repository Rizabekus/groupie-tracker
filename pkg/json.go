package pkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Final struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	ConcertDates []string
}

type ErrorStruct struct {
	Status  int
	Message string
}

type Locations struct {
	Locations []string `json:"locations"`
}

type ConcertDates struct {
	Dates []string `json:"dates"`
}

func GetApi() []Artist {
	responseArtist, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer responseArtist.Body.Close()
	respArtist, err := ioutil.ReadAll(responseArtist.Body)
	if err != nil {
		log.Fatal(err)
	}
	var Artists []Artist
	json.Unmarshal(respArtist, &Artists)
	// ArtistAPI
	return Artists
}

func GetApi2(artists []Artist, num int) Final {
	Locs, err := http.Get(artists[num-1].Locations)
	if err != nil {
		log.Fatal(err)
	}
	defer Locs.Body.Close()
	respLocs, err := ioutil.ReadAll(Locs.Body)
	if err != nil {
		log.Fatal(err)
	}
	var LocationsStruct Locations
	json.Unmarshal(respLocs, &LocationsStruct)
	// Specific Locations
	Dates, err := http.Get(artists[num-1].ConcertDates)
	if err != nil {
		log.Fatal(err)
	}
	defer Dates.Body.Close()
	respDates, err := ioutil.ReadAll(Dates.Body)
	if err != nil {
		log.Fatal(err)
	}
	var DatesStruct ConcertDates
	json.Unmarshal(respDates, &DatesStruct)

	result := Final{
		Id:           artists[num-1].Id,
		Image:        artists[num-1].Image,
		Name:         artists[num-1].Name,
		Members:      artists[num-1].Members,
		CreationDate: artists[num-1].CreationDate,
		FirstAlbum:   artists[num-1].FirstAlbum,
		Locations:    LocationsStruct.Locations,
		ConcertDates: DatesStruct.Dates,
	}
	return result
	// Specific Dates
}
