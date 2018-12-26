package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"twitch_replica/jsons"
)

//TWITCH_API_KEY is Used to fetch from twitch_api
const TWITCH_API_KEY = "0t78gp7euwf2p2nb17cimvf3sxd3ji"

//GetGameID converts names of games to their Twitch IDs
func GetGameID(gameName string) string {

	urlRequest := fmt.Sprintf("https://api.twitch.tv/helix/games?name=%s", gameName)
	req, err := http.NewRequest("GET", urlRequest, nil)
	if err != nil {
		log.Fatal("Could not create a new api request to find game_Id of ", gameName)
	}
	req.Header.Set("Client-Id", TWITCH_API_KEY)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Could not complete the api request.")
	}
	defer resp.Body.Close()

	var data jsons.GameData
	json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal("Json decoding failed.")
	}
	return data.Data[0].ID
}
