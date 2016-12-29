package main

type Team struct {
	TeamId       int      `json:"team_id" gorm:"primary_key:true;"`
	Abbrev       string   `json:"abbreviation"`
	Name         string   `json:"name"`
	Owner        Owner    `json:"owner"`
	StartingPick int      `json:"pick"`
	Roster       []Player `gorm:"ForeignKey:PlayerId;" json:"all"`
	Picks        []Pick   `json:"picks"`
}

type Player struct {
	PlayerId             int    `gorm:"primary_key:true;"`
	Name                 string `gorm:"unique" json:"name"`
	Position             string `json:"position"`
	NflTeam              string `json:"nfl_team"`
	Keeper               bool   `json:"keeper"`
	Bye                  int    `json:"bye"`
	FantasyFootballersId int
	Ranks                FantasyFootballers `gorm:"ForeignKey:FantasyFootballersId" json:"fantasyfootballers"`
	DraftWizardId        int
	DW                   DraftWizard `gorm:"ForeignKey:DraftWizardId" json:"draftwizard"`
}

type Keeper struct {
	Player
	KeepingTeam string `json:"abbreviation"`
	Team        Team   `json:"team"`
}

func (r *Team) getPosition(p string) (players []Player) {
	for i := range r.Roster {
		if r.Roster[i].Position == p {
			players = append(players, r.Roster[i])
		}
	}
	return
}
