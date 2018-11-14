package main
 
import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
 	. "github.com/fahadem/championship"
	"github.com/gorilla/mux"
)
 
func AllLeagueEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}
 
func FindLeagueEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}
 
func CreateLeagueEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var league League
	if err := json.NewDecoder(r.Body).Decode(&League); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	league.Country_id = bson.NewObjectId()
	league.Country_name = "England"
	league.League_id = bson.NewObjectId()
	league.League_name = "Premier League"
	if err := dao.Insert(league); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, league)
}
 
func UpdateLeagueEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}
 
func DeleteLeagueEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}
 
func main() {
	addr, err := determineListenAddress()
  	if err != nil {
    		log.Fatal(err)
  	}
	r := mux.NewRouter()
	r.HandleFunc("/league", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/league", CreateLeagueEndPoint).Methods("POST")
	r.HandleFunc("/league", UpdateLeagueEndPoint).Methods("PUT")
	r.HandleFunc("/league", DeleteLeagueEndPoint).Methods("DELETE")
	r.HandleFunc("/league/{id}", FindLeagueEndpoint).Methods("GET")
	
	log.Fatal(http.ListenAndServe(addr,nil))

	/*if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}*/
}
