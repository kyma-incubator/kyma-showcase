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
