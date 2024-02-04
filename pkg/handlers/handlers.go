package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Player struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Role string `json:"role"`
}

var PlayersList = []Player{
	{"Cristiano Ronaldo", 36, "Forward"},
	{"Sergio Ramos", 35, "Defender"},
	{"Karim Benzema", 33, "Forward"},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/players", GetPlayersList).Methods("GET")
	router.HandleFunc("/players/{name}", GetPlayerDetails).Methods("GET")
	router.HandleFunc("/health", HealthCheck).Methods("GET")

	return router
}

func GetPlayersList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PlayersList)
}

func GetPlayerDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	playerName := params["name"]

	for _, player := range PlayersList {
		if player.Name == playerName {
			json.NewEncoder(w).Encode(player)
			return
		}
	}
	http.Error(w, "Player not found", http.StatusNotFound)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Web App is healthy!(Author: Dimash)"))
}