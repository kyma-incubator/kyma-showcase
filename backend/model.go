package main

// Image defines a struct to use for the received image
type Image struct {
	URL string `json:"url"`
	GCP string `json:"gcp"`
	IMG string `json:"img"`
}
