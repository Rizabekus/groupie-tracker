package pkg

type Artists struct {
	id           int      `json:"id"`
	image        string   `json:"image"`
	name         string   `json:"name"`
	members      []string `json:"members"`
	creationDate int      `json:"creationDate"`
	firstAlbum   string   `json:"firstAlbum"`
	locations    string   `json:"locations"`
	concertDates string   `json:"concertDates"`
	relations    string   `json:"relations"`
}

// type Locations struct {
// 	locations []string `json:"locations"`
// 	dates     []Dates  `json:"dates"`
// }

// type Dates struct {
// 	dates []string `json:"dates"`
// }

// type Relations struct{}
