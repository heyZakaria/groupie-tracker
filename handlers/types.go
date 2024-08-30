package handlers

type Artist struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	Images       string   `json:"image"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Relations struct {
	Index []struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
