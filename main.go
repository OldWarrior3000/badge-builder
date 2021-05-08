package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var Region = ""
var RegistryId = ""

func retrieveImages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	repositoryId := vars["repositoryName"]

	fmt.Println("Endpoint Hit: retrieveImages")
	var imageIds = GetSortedImageIds(Region, RegistryId, repositoryId)

	json.NewEncoder(w).Encode(imageIds)
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
		os.Exit(-1)
	}

	RegistryId = *registryId
	Region = *region

	handleRequests()
}
