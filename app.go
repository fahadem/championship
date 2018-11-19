package main

import (
	"github.com/fahadem/championship/leaguedb"
	"github.com/fahadem/championship/matchdb"
	"fmt"
	"net/http"
	"os"
)

func champ(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "api for championship league")
}

func main() {
	leaguedb.InitWh()
	matchdb.InitWh()

	//in memory storage
	leaguedb.Global_db = &leaguedb.LeaguesMongoDB{
		DatabaseURL:           "mongodb://fahadem:269093f@ds051658.mlab.com:51658/championship",
		DatabaseName:          "championship",
		LeaguesCollectionName: "league",
	}

	//mongodb storage
	leaguedb.Global_db.Init()

	//in memory storage
	matchdb.Global_db = &matchdb.MatchesMongoDB{
		DatabaseURL:           "mongodb://fahadem:269093f@ds051658.mlab.com:51658/championship",
		DatabaseName:          "championship",
		MatchesCollectionName: "matches",
	}

	//mongodb storage
	matchdb.Global_db.Init()

	port := os.Getenv("PORT")

	http.HandleFunc("/champ", champ)
	http.HandleFunc("/champ/leagues/", leaguedb.LeagueHandler)
	http.HandleFunc("/champ/webhook/", leaguedb.WebhookHandler) //POST & GET
	http.HandleFunc("/champ/matches/", matchdb.MatchHandler)
	http.HandleFunc("/champ/webhookM/", leaguedb.WebhookHandler) //POST & GET
	http.ListenAndServe(":"+port, nil)
}
