package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"vinhis.me/quote-share/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/quote", handlers.HandleQuote).Methods("GET")

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}
