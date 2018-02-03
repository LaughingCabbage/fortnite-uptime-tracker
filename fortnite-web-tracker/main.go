package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const statusEndpoint = "https://lightswitch-public-service-prod06.ol.epicgames.com/lightswitch/api/service/bulk/status?serviceId=Fortnite"

func main() {
	log.Println("fortnite-uptime-tracker web start")

	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":8080", r))

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("serving root")

	//get the json status from the server
	status, err := getServerStatus()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var serverStatus []ServerStatus
	if err := json.Unmarshal(status, &serverStatus); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := Data{"status": serverStatus}
	renderTemplate(w, "root", data)
}

func getServerStatus() ([]byte, error) {
	//create a request to the subroute to retreive a json response
	req, err := http.Get(statusEndpoint)
	if err != nil {
		return nil, err
	}
	//close request body on return
	defer req.Body.Close()
	fmt.Println(req.Header)
	//read the response from the body
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	//return the byte array of the body
	return body, nil
}
