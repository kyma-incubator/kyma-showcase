package model

// Image defines a struct to use for the received image
type Image struct {
	ID      string `json:"id"`      //nanoID
	Content string `json:"content"` //base64
	GCP     string `json:"gcp"`     //returned JSON from gcp api
	Status  bool   `json:"status"`  //defines if the image was already processed
}

// ID defines a struct containing id filed
type ID struct {
	ID string `json:"id"` //id
}

// Event defines a struct that triggers lambdas
type Event struct {
	Source           string `json:"source"`
	SpecVersion      string `json:"specversion"`
	EventTypeVersion string `json:"eventtypeversion"`
	Data             string `json:"data"`
	DataContentType  string `json:"datacontenttype"`
	Id               string `json:"id"`
	EventType        string `json:"type"`
}
