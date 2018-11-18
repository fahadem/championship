package leaguedb

type LeaguesStorage interface {
	Init()
	Add(l League) error
}

type Team struct {
	Name 	string `json:"name"`
	TeamID 	string `json:"teamid"`
}

type League struct {
	Name     string `json:"name"`
	Country  string `json:"country"`
	LeagueID string `json:"leagueid"`
	Teams 	 Team   `json:"teams"`
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
