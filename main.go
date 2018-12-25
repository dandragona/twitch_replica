package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//DirectoryHandler handles requests for the `/directory` resource
func DirectoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Twitch Directory\n"))
}

//TwitchHomeHandler handles requests for the `/` resource
func TwitchHomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home of Twitch_Replica.tv\n"))
}

//GameHandler handles requests for the `/directory/{game}` resource
func GameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game := vars["game"]
	header := fmt.Sprintf("Twitch data for %v game.\n", game[0])
	w.Write([]byte(header))
}

//StreamHandler handles requests for the `/directory/{game}/{stream}` resource
func StreamHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game := vars["game"]
	stream := vars["stream"]
	header := fmt.Sprintf("Twitch data for %v game and %v streamer.\n", game, stream)
	w.Write([]byte(header))
}

func main() {
	//get the value of the ADDR environment variable
	addr := os.Getenv("ADDR")

	//if it's blank, default to ":80", which means
	//listen port 80 for requests addressed to any host
	if len(addr) == 0 {
		addr = ":80"
	}

	//create a new mux (router)
	//the mux calls different functions for
	//different resource paths
	r := mux.NewRouter()

	//tell it to call the HelloHandler() function
	//when someone requests the resource path `/hello`
	r.HandleFunc("/", TwitchHomeHandler).Methods("GET")
	r.HandleFunc("/directory", DirectoryHandler).Methods("GET")
	r.HandleFunc("/directory/{game}", GameHandler).Methods("GET")
	r.HandleFunc("/directory/{game}/{stream}", StreamHandler).Methods("GET")

	//start the web server using the mux as the root handler,
	//and report any errors that occur.
	//the ListenAndServe() function will block so
	//this program will continue to run until killed
	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
