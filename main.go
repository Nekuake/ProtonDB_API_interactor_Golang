package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func initializeData() map[string]interface{} {
	resp, err := http.Get("https://protondb.max-p.me/reports.json")
	if err != nil {
		fmt.Printf("Error while making request: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	// Read content of request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error while reading content: %v\n", err)
		return nil
	}

	// JSON -> map
	var protonMap map[string]interface{}
	err = json.Unmarshal(body, &protonMap)
	if err != nil {
		fmt.Printf("Error while paring JSON: %v\n", err)
		return nil
	}
	return protonMap
}

func getData(gameID string, gameDB map) {

}

func main() {
	fmt.Println("Downloading protonDB Json...")

	fmt.Println("Input Steam game ID")
	var gameId string
	fmt.Scanln(&gameId)
}
