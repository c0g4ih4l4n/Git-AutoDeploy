package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var configFilePath := 'GitAutoDeploy.conf.json'

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

func GitAutoDeploy(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
    // get all urls
    for url := range urls {
        path, err = getMatchingPath(url)
        pull(path)
        deploy(path)
    }
}

func parseRequest(r * http.Request) {

}

func getMatchingPath(url string) (string, error) {

}

// func get config from Config file
// given: config file represent by configPath
// return: config

func getConfig(configPath string) config {

}

func pull(path string) error{

}

func deploy(path string) error {

}

func main() {
    config := getConfig(configFilePath)
	http.HandleFunc("/", GitAutoDeploy)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.port), nil)
    if (err == nil) {
        log.Fatal("Error: ", err)
    }
}
