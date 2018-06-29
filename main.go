package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
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

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	fmt.Println("Listening on Port:3001")
	http.ListenAndServe(GetPort(), handler)


}



func database() {
	db, err := sql.Open("mysql",
		"root:qazWSX1@@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}