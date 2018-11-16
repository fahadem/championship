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
	//in memory storage
	leaguedb.Global_db = &leaguedb.LeaguesMongoDB{
		DatabaseURL:           "mongodb://fahadem:Basket-ball11@ds051658.mlab.com:51658/championship",
		DatabaseName:          "championship",
		LeaguesCollectionName: "league",
	}

	//mongodb storage
	leaguedb.Global_db.Init()

	port := os.Getenv("PORT")

	http.HandleFunc("/champ", champ)
	//http.HandleFunc("/champ/league", leaguedb.LeagueHandler)
	http.ListenAndServe(":"+port, nil)
}
