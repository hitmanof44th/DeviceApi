package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/mux"
)

//====================================================================
func deviceStatus(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]

    fmt.Fprintf(w, "Key: " + key)
}

//====================================================================
func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/status/{id}", deviceStatus)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

//====================================================================
func main() {

	fmt.Println("Rest API v1.0 - Device api mockup")
	handleRequests();
}