package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"twitch_replica/jsons"
	"twitch_replica/utils"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Used to fetch from twitch_api
const TWITCH_API_KEY = "0t78gp7euwf2p2nb17cimvf3sxd3ji"

//DirectoryHandler handles requests for the `/directory` resource
func DirectoryHandler(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", "https://api.twitch.tv/kraken/games/top?limit=5", nil)
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
	games := data.Top
	var gameResponses []jsons.GameResponse
	for _, game := range games {
		var gameResponse jsons.GameResponse
		gameResponse.Name = game.Game.Name
		gameResponse.Img = game.Game.Box.Small
		gameResponse.Viewers = game.Viewers
		gameResponse.Link = "localhost:8000/directory/" + game.Game.Name
		gameResponses = append(gameResponses, gameResponse)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(gameResponses)
	return
}

//GameHandler handles requests for the `/directory/{game}` resource
func StreamsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game := vars["game"]
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
	var data jsons.Streams
	json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal("Json decoding failed.")
	}
	var streamResponses []jsons.StreamResponse
	for _, stream := range data.Data {
		var streamResponse jsons.StreamResponse
		streamResponse.UserName = stream.User_Name
		var img = strings.Split(stream.Thumbnail_URL, "{")[0]
		img += "100x100.jpg"
		streamResponse.Img = img
		streamResponse.Viewers = stream.Viewer_Count
		streamResponse.Title = stream.Title
		streamResponses = append(streamResponses, streamResponse)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(streamResponses)
	return
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
	r.HandleFunc("/directory", DirectoryHandler).Methods("GET")
	r.HandleFunc("/directory/{game}", StreamsHandler).Methods("GET")

	//start the web server using the mux as the root handler,
	//and report any errors that occur.
	//the ListenAndServe() function will block so
	//this program will continue to run until killed
	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, handlers.CORS()(r)))
}
