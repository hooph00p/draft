package main

type League struct {
	LeagueId   int      `gorm:"primary_key:true"`
	Name       string   `json:"name"`
	Setting    Settings `gorm:"ForeignKey:SettingsId" json:"settings"`
	SettingsId int
	Teams      []Team `gorm:"ForeignKey:TeamId"`
}

type Settings struct {
	SettingsId int           `gorm:"primary_key:true"`
	Starters   []Slot        `gorm:"ForeignKey:SlotId" json:"starters"`
	Maximum    []Slot        `gorm:"ForeignKey:SlotId" json:"maximum"`
	Scoring    ScoringScheme `json:"league_scoring"`
	RosterSize int           `json:"roster_size"`
}

type Slot struct {
	SlotId   int    `gorm:"primary_key:true"`
	Position string `json:"position"`
	Capacity int    `json:"capacity"`
}

type Owner struct {
	OwnerId int      `gorm:"primary_key:true"`
	Leagues []League `gorm:"ForeignKey:LeagueId" json:"leagues"`
	Name    string   `json:"owner_name"`
}

type Pick struct {
	PickId  int `gorm:"primary_key:true;" json:"pick_id"`
	Round   int `json:"round"`
	Overall int `json:"overall"`
}

type Transaction struct {
	Alpha        Owner    `gorm:"ForeignKey:OwnerId" json:"alpha_owner"`
	AlphaPlayers []Player `gorm:"ForeignKey:PlayerId" json:"alpha_players"`
	AlphaPicks   []Pick   `gorm:"ForeignKey:PickId" json:"alpha_picks"`
	Beta         Owner    `gorm:"ForeignKey:OwnerId" json:"beta_owner"`
	BetaPlayers  []Player `gorm:"ForeignKey:PlayerId" json:"beta_players"`
	BetaPicks    []Pick   `gorm:"ForeignKey:PickId" json:"beta_picks"`
}
