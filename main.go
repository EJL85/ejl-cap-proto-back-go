package main

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/rs/cors"
	"os"
)

type Map struct {
	Src string `json:"Src"`
	Name string `json:"Name"`
}

type Maps []Map

func allInfo(w http.ResponseWriter, r *http.Request) {
	maps := Maps{
		Map{Src: "Test Source", Name:"Test Name"},
	}

	fmt.Println("Endpoint Hit: All Maps endpoint")
	json.NewEncoder(w).Encode(maps)
}

func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func handleRequests() {
	http.HandleFunc("/", allInfo)
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func main() {

	fmt.Println("Starting server")

	port := os.Getenv("PORT")

	if port == "" {
		port = "3001"
		fmt.Println("Defaulting to port 3001")
	}

	fmt.Println(port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", allInfo)

	handler := cors.Default().Handler(mux)
	fmt.Println("Listening on Port:3001")
	http.ListenAndServe(GetPort(), handler)
}
