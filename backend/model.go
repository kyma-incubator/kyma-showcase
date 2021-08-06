package main

// Image defines a struct to use for the received image
type Image struct {
	EncodedBase64 string `json:"base64"`
	GCP string `json:"gcp"`
	IMG string `json:"img"`
}

type ImageID struct {
	ID string `json:"id"`
}
