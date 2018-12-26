package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"twitch_replica/jsons"
	"twitch_replica/utils"

	"github.com/gorilla/mux"
)

// Used to fetch from twitch_api
const TWITCH_API_KEY = "0t78gp7euwf2p2nb17cimvf3sxd3ji"

//DirectoryHandler handles requests for the `/directory` resource
func DirectoryHandler(w http.ResponseWriter, r *http.Request) {
	header := "Games Directory\n\n"
	req, err := http.NewRequest("GET", "https://api.twitch.tv/kraken/games/top?limit=100", nil)
	if err != nil {
		log.Fatal("Could not create the new api request.")
	}
	req.Header.Set("Client-Id", TWITCH_API_KEY)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Could not complete the api request.")
	}
	defer resp.Body.Close()
	var data jsons.TopGames
	json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal("Json decoding failed.")
	}
	w.Write([]byte(header))
	games := data.Top
	responseString := ""
	for i, game := range games {
		gameString := fmt.Sprintf("(%d) %s -- Viewers: %d", i, game.Game.Name, game.Viewers)
		responseString = fmt.Sprintf("%s%s\n\n", responseString, gameString)
	}
	w.Write([]byte(responseString))

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("%s", b)))
}

//TwitchHomeHandler handles requests for the `/` resource
func TwitchHomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home of Twitch_Replica.tv\n"))
}

//GameHandler handles requests for the `/directory/{game}` resource
func GameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game := vars["game"]
	header := fmt.Sprintf("Twitch data for %v.\n\n", game)
	game_name := utils.GetGameID(game)
	urlRequest := fmt.Sprintf("https://api.twitch.tv/helix/streams?game_id=%s", game_name)
	req, err := http.NewRequest("GET", urlRequest, nil)
	if err != nil {
		log.Fatal("Could not create a new api request for ", game)
	}
	req.Header.Set("Client-Id", TWITCH_API_KEY)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Could not complete the api request.")
	}
	defer resp.Body.Close()
	w.Write([]byte(header))
	var data jsons.Streams
	json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal("Json decoding failed.")
	}

	responseString := ""
	for i, stream := range data.Data {
		streamString := fmt.Sprintf("(%d) Title: %s\nUser: %s\nViewers: %d", i, stream.Title, stream.User_Name, stream.Viewer_Count)
		responseString = fmt.Sprintf("%s%s\n\n", responseString, streamString)
	}
	w.Write([]byte(responseString))
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
