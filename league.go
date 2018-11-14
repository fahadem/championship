package models

import "gopkg.in/mgo.v2/bson"

type League struct {
	Country_id	bson.ObjectId	`bson:"CountryID" json:"CountryID,omitempty"`
	Country_name	string		`bson:"Country" json:"Country,omitempty"`
	League_id	bson.ObjectId	`bson:"LeagueID" json:"LeagueID,omitempty"`
	League_name	string		`bson:"League" json:"League,omitempty"`
}

/*type Team struct {
	InfoLeague	League	`bson:"InfoLeague" json:"InfoLeague,omitempty"`
	Team_id		bson.ObjectId	`bson:"TeamID" json:"TeamID,omitempty"`
	Team_name	string	`bson:"Team" json:"Team,omitempty"`
}

type Match struct {
	DateMatch 	Time	`bson:"Date" json:"Date,omitempty"`
	FirstTeam	Team	`bson:"FirstTeam" json:"FirstTeam,omitempty"`
	SecondTeam 	Team	`bson:"SecondTeam" json:"SecondTeam,omitempty"`
}*/


