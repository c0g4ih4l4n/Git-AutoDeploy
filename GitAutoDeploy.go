package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	//"reflect"
)

const configFilePath = "conf.json"

// hold config infomation
// when get from configFile
type config struct {
	port         float64
	repositories []map[string]interface{}
}

// objective: replace interface{} in repositories
// not done yet
type repo struct {
	url  string
	path string
}

// must delete this
type test_struct struct {
	test string
}

// serve func
// flow: req -> parse request to get param (urls) ->
//          -> loop all urls -> get path of each url -> pull -> deploy

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
	var result config                      // hold result
	var fileContent map[string]interface{} // hold file content when get read file
	data := make([]byte, 1000)             /* hold data in fileConfig
	value 1000 is len of slice (going to change it)
	config file have length unknown and must loop to read all data */

	// open file
	file, err := os.Open(configPath)
	check(err)

	// open stream and read from file to data

	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("read %d bytes: %q\n", count, data[:count])

	// convert data to fileContent
	err = json.Unmarshal(data[:count], &fileContent)
	check(err)

	result.port = fileContent["port"].(float64)
	repositories := fileContent["repositories"].([]interface{})

	for _, value := range repositories {
		element := value.(map[string]interface{})
		result.repositories = append(result.repositories, element)
	}

	return result
}

func pull(path string, cfg config) error {

	return nil
}

func deploy(path string) error {
	return nil
}

func respond() error {
	return nil
}

func main() {
	fmt.Println("WebService - GitAutoDeploy Starting ... ")
	var cfg config
	cfg = getConfig(configFilePath)
	http.HandleFunc("/", GitAutoDeploy)

	// fmt.Println(cfg)

	fmt.Printf("GitAutoDeploy Listening At Port %v ...", cfg.port)

	// pass port into string
	// string need to pass to ListenandServe
	strPort := strconv.Itoa(int(cfg.port))
	strPort = ":" + strPort

	err := http.ListenAndServe(strPort, nil)
	// err := http.ListenAndServe(":80", nil)
	if err == nil {
		log.Fatal("Error: ", err)
	}

}
