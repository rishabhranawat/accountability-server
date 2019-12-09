package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: homepage")

	data := map[string]interface{} {"status" : "200", "message" : "rishabh"}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}


func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main(){
	handleRequests()
}