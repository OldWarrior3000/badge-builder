package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var Region = ""
var RegistryId = ""

func retrieveImages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	repositoryId := vars["repositoryName"]

	fmt.Println("Endpoint Hit: retrieveImages")
	imageIds := GetSortedImageIds(Region, RegistryId, repositoryId)

	if imageIds != nil {
		_ = json.NewEncoder(w).Encode(imageIds[len(imageIds)-1])
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/{repositoryName}", retrieveImages)
	fmt.Print("Starting service")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	region, registryId, err := LoadConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	RegistryId = *registryId
	Region = *region

	handleRequests()
}
