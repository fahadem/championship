package leaguedb

type LeaguesStorage interface {
	Init()
	Add(l League) error
}

type Team struct {
	Key 	string `json:"key"`
	Name 	string `json:"name"`
	Code 	string `json:"code"`
}

type Match struct {
	Date 	string `json:"date"`
	Team1	Team   `json:"team1"`
	Team2	Team   `json:"team2"`
	Score1	int64  `json:"score1"`
	Score2	int64  `json:"score2"`
}

type Rounds struct {
	Name 	string  `json:"name"`
	Matches []Match `json:"matches"`
}

type MatchesL struct {
	Name 	string  `json:"name"`
	Rounds	Rounds  `json:"rounds"`
}
type League struct {
	Name     string `json:"name"`
	Country  string `json:"country"`
	LeagueID string `json:"leagueid"`
	Teams 	 []Team `json:"teams"`
}

type LeaguesDB struct {
	leagues map[string]League
}

func (db *LeaguesDB) Init() {
	db.leagues = make(map[string]League)
}

func (db *LeaguesDB) Add(l League) error {
	db.leagues[l.LeagueID] = l
	return nil
}
