package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("fortnite-uptime-tracker web start")

	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":8080", r))

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("serving root")
	renderTemplate(w, "root", nil)
}
