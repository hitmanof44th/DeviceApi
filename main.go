package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/mux"
)
//====================================================================
func deviceAdd(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]
    fmt.Fprintf(w, "Key: " + key)
}
//====================================================================
func deviceStatus(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]
    fmt.Fprintf(w, "Key: " + key)
}

//====================================================================
func apiStatus(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Rest API v1.0 - Device api mockup")
}

//====================================================================
func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", apiStatus).Methods("GET")

	
	devicesrouter := myRouter.PathPrefix("/device").Subrouter()
	devicesrouter.HandleFunc("", deviceAdd).Methods("POST")
	devicesrouter.HandleFunc("/status/{id}", deviceStatus).Methods("GET")
	
	
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

//====================================================================
func main() {

	fmt.Println("Rest API v1.0 - Device api mockup")
	handleRequests();
}