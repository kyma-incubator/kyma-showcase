package model

// Image defines a struct to use for the received image
type Image struct {
	ID      string   `json:"id"`      //nanoID
	Content string   `json:"content"` //base64
	GCP     []string `json:"gcp"`     //returned JSON from gcp api
	Time    string   `json:"time"`    //defines the time when image was posted
}

// ID defines a struct containing id filed
type ID struct {
	ID string `json:"id"` //id
}
