package main

import (
	"fmt"
	"net/http"
)

func initializeData() {
	response, err := http.Get("https://protondb.max-p.me/reports.json")
}

func getData(gameID string) {
}

func main() {

	fmt.Println("Input Steam game ID")
	var gameId string
	fmt.scanln(&gameId)

}
