package main

import (
	"github.com/fahadem/championship/leaguedb"
	"fmt"
	"net/http"
	"os"
)

func champ(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "api for championship league")
}

func main() {
	leaguedb.InitWh()

	//in memory storage
	leaguedb.Global_db = &leaguedb.LeaguesMongoDB{
		DatabaseURL:           "mongodb://fahadem:269093f@ds051658.mlab.com:51658/championship",
		DatabaseName:          "championship",
		LeaguesCollectionName: "league",
	}

	//mongodb storage
	leaguedb.Global_db.Init()

	port := os.Getenv("PORT")

	http.HandleFunc("/champ", champ)
	http.HandleFunc("/champ/league", leaguedb.LeagueHandler)
	http.HandleFunc("/champ/webhook/", leaguedb.WebhookHandler) //POST & GET
	http.ListenAndServe(":"+port, nil)
}
