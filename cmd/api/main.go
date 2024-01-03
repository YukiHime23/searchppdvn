package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	searchinppdvn "github.com/YukiHime23/search-in-ppdvn"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "Missing 'name' parameter", http.StatusBadRequest)
		return
	}

	nameQuery := strings.Replace(name, " ", "+", -1)
	resultJson := searchinppdvn.Search(nameQuery)
	jsonData, err := json.Marshal(resultJson)
	if err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	// Create ServeMux
	mux := http.NewServeMux()

	// Define routes and handle requests.
	mux.HandleFunc("/api/search-in-ppdvn", searchHandler)

	// Run the server on port 8787.
	server := &http.Server{
		Addr:    ":8787",
		Handler: mux,
	}

	fmt.Println("Server is running on :8787")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
