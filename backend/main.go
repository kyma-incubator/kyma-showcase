package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var (
	connection = connectToRedis()
	ctx        = context.Background()
)

type Image struct {
	URL string `json:"url"`
	GCP string `json:"GCP"`
	Img string `json:"Img"`
}

func connectToRedis() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil
	}

	return rdb
}
func insertToDB(key string, value string) error {
	_, err := connection.Set(ctx, key, value, 0).Result()
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func getFromDB(key string) (interface{}, error) {
	val, err := connection.Get(ctx, key).Result()

	switch {
	case err == redis.Nil:
		fmt.Printf("%s key does not exist", key)
	case err != nil:
		fmt.Println(err)
	case val == "":
		fmt.Println("value is empty")
	}

	return val, err //fmt.Sprintf("%s", val)
}

func allKeys() []string {
	keys, err := connection.Keys(ctx, "*").Result()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return keys
}

func dbPostHandler(w http.ResponseWriter, r *http.Request) {
	var img Image
	err := json.NewDecoder(r.Body).Decode(&img)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)

	if /*r.Method == "POST" */ params["id"] == "999" {
		fmt.Println("json =method POST is not supported for all keys!")
	} else {
		j, err := json.Marshal(img) //key: url | value: {url:"...", GCP:"..."}
		if err != nil {
			fmt.Println(err)
		}

		insertToDB(params["id"], string(j))
	}

}

func dbGetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var img Image

	//999 only for GET.
	//jeśli error to ma nic nie robić

	if params["id"] == "999" {
		keys := allKeys()
		for _, key := range keys {
			fromDB, err := getFromDB(key)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Fprintf(w, "json = %s\n", fromDB)
			}
		}
	}

	fromDB, err := getFromDB(params["id"])
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(fromDB.(string)), &img)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "json = %s\n", fromDB)
	fmt.Fprintf(w, "json URL = %s,  json GCP = %s, json image = %s\n", img.URL, img.GCP, img.Img)

}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func fromBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func loadImg(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	//imageType := http.DetectContentType(bytes)
	//var imgString string

	//czy to ma leciec do GCP???
	// switch imageType {
	// case "image/jpeg":
	// 	imgString += "data:image/jpeg;base64," //czy trzeba te znaczniki dodawać przy wyslaniu do GCP??
	// case "image/png":
	// 	imgString += "data:image/png;base64,"
	// }

	imgString := toBase64(bytes)

	return imgString, err
}

func saveImg(imgString string, outFile string) {

	img, _ := fromBase64(imgString)
	err := ioutil.WriteFile(outFile, img, 0666)
	if err != nil {
		fmt.Println(err)
	}
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	if connection != nil {
		inFile := "image.png"
		err := downloadFile("https://i.pinimg.com/originals/54/9b/11/549b114dad455ae154295ecad1d05f71.png", inFile)
		if err != nil {
			fmt.Println(err)
		}
		img, err := loadImg(inFile)
		if err != nil {
			fmt.Println(err)
		}

		j, err := json.Marshal(Image{URL: img, GCP: "gcp"}) //key: url | value: {url:"...", GCP:"..."}
		if err != nil {
			fmt.Println(err)
		}

		insertToDB("id", string(j))

		imgFromDB, err := getFromDB("id")
		if err != nil {
			fmt.Println(err)
		}

		var x Image
		err = json.Unmarshal([]byte(imgFromDB.(string)), &x)
		if err != nil {
			fmt.Println(err)
		}

		saveImg(x.URL, "outFile.png")

		router := mux.NewRouter()

		router.HandleFunc("/get/{id}", dbGetHandler).Methods("GET")
		router.HandleFunc("/post/{id}", dbPostHandler).Methods("POST")

		fmt.Printf("Starting server at port 8081\n")
		http.ListenAndServe(":8081", router)
	}
}
