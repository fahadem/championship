# championship

Build RESTful API in Go and MongoDB
The REST API service will expose endpoints to manage a store of leagues and teams.
The API will allow to see data on leagues, teams and matchs.The API will also allow to receive alerts when new data are added using Slack and webhooks via a registration.

The app url is https://chmpship.herokuapp.com/ .

The home page is https://chmpship.herokuapp.com/champ

To have information about the leagues just use https://chmpship.herokuapp.com/champ/leagues/ 

To have information about the teams just use https://chmpship.herokuapp.com/champ/leagues/{id}/teams 

To have information about the matches just use https://chmpship.herokuapp.com/champ/matches/ 

Webhook League url : https://chmpship.herokuapp.com/champ/webhook/ 

Webhook Match url : https://chmpship.herokuapp.com/champ/webhookM/ 


To create a league : curl -d "@file.json" -X POST https://chmpship.herokuapp.com/champ/league 

To create matches associate to a league : curl -d "@file.matches.json" -X POST https://chmpship.herokuapp.com/champ/league

(replace "file" by the three first letters of a country)
