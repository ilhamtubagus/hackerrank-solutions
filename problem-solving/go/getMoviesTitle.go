package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

var client = &http.Client{}
var res map[string]interface{}

func getMovieTitles(substr string) []string {
	request, err := http.NewRequest("GET", fmt.Sprintf("https://jsonmock.hackerrank.com/api/movies/search/?Title=%s", substr), nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		log.Fatal(err)
	}
	data := res["data"].([]interface{})
	result := []string{}
	for _, x := range data {
		y := x.(map[string]interface{})
		result = append(result, fmt.Sprintf("%v", y["Title"]))
	}
	sort.Strings(result)
	return result
}
