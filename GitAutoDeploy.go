package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const configFilePath = "GitAutoDeploy.conf.json"

type config struct {
	port         int
	repositories []repo
}

type repo struct {
	url    string
	path   string
	deploy string
}

// serve func
// flow: req -> parse request to get param (urls) ->
//          -> loop all urls -> get path of each url -> pull -> deploy

type test_struct struct {
	key   string
	value interface{}
}

func GitAutoDeploy(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r)

	var t test_struct
	for key, _ := range r.Form {
		log.Println(key)
		//LOG: {"test": "that"}
		err := json.Unmarshal([]byte(key), &t)
		if err != nil {
			log.Println(err.Error())
		}
	}
	log.Println(t)

	// var result map[string]interface{}

	// data := json.Unmarshal(r, &result);

	// repository := data["repository"](map[string]string)
	// url := repository["url"](string)

	// fmt.Println("url")
	// get all urls
	// for url := range urls {
	//     path, err = getMatchingPath(url)
	//     pull(path)
	//     deploy(path)
	// }
}

func parseRequest(r *http.Request) {

}

func getMatchingPath(url string) (string, error) {

	return "", nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// func get config from Config file
// given: config file represent by configPath
// return: config

func getConfig(configPath string) config {
	var result config

	file, err := os.Open(configPath) // For read access.
	check(err)

	data := make([]byte, 1000)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])

	var t test_struct

	err = json.Unmarshal(data[:count], &t)
	check(err)

	fmt.Printf("%+v", t)

	return result
}

func pull(path string) error {
	return nil
}

func deploy(path string) error {
	return nil
}

func main() {
	fmt.Println("WebService - GitAutoDeploy Starting ... ")
	config := getConfig(configFilePath)
	http.HandleFunc("/", GitAutoDeploy)

	fmt.Println(config)

	// config.port = 80
	// fmt.Printf("GitAutoDeploy Listening At Port %v ...", config.port)
	// err := http.ListenAndServe(fmt.Sprintf(":%d", config.port), nil)
	// if err == nil {
	// 	log.Fatal("Error: ", err)
	// }
}
