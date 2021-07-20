package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	//val, err := connection.Do(ctx, "GET", key).Result()
	val, err := connection.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
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
	j, err := json.Marshal(Image{URL: "url", GCP: "gcp"}) //key: url | value: {url:"...", GCP:"..."}
	if err != nil {
		fmt.Println(err)
	}

	params := mux.Vars(r)
	insertToDB(params["id"], string(j))

}

func dbGetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var x Image

	//999 only for GET.

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
	} else {
		fromDB, err := getFromDB(params["id"])
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal([]byte(fromDB.(string)), &x)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Fprintf(w, "json = %s\n", fromDB)
		fmt.Fprintf(w, "json URL = %s,  json GCP = %s\n", x.URL, x.GCP)
	}
}

func main() {

	if connection != nil {

		router := mux.NewRouter()

		router.HandleFunc("/get/{id}", dbGetHandler).Methods("GET")
		router.HandleFunc("/add/{id}", dbPostHandler).Methods("POST")

		fmt.Printf("Starting server at port 8081\n")
		http.ListenAndServe(":8081", router)
	}
}
