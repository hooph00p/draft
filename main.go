package main

import (
	"bufio"
	_ "bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

const (
	etc_path    = "/Users/jadaho/work/go/src/bitbucket.org/hooph00p/draft/etc"
	teams_file  = etc_path + "/stat/teams.json"
	keeper_file = etc_path + "/stat/keepers.json"

	composite_file = etc_path + "/gen/composite.json"
	target_file    = etc_path + "/gen/data.json"
	draft_wizard   = etc_path + "/stat/dw.html"

	_db    = "sqlite3"
	_dbloc = "db/draft.db"
)

func main() {
	things()
	// reset()
	// auto()
	// populate()
}

func things() {
	loc := "etc/stat/league/"
	ln, _ := ioutil.ReadDir(loc)
	for i := range ln {
		if ln[i].IsDir() {
			dir := ln[i].Name()
			settings, err := os.Open(loc + dir + "/raw/settings")
			if err == nil {

			}
			defer settings.Close()

			// teams, err := os.Open(loc + dir + "/raw/teams")
			// if err == nil {
			// 	scanner := bufio.NewScanner(teams)
			// 	for scanner.Scan() {
			// 		line := scanner.Text()
			// 		header := strings.Contains(line, "#")
			// 		if header {
			// 			fmt.Println(strings.Repeat("-", 120))
			// 		} else {
			// 			fmt.Println(line)
			// 		}
			// 	}
			// }
			// defer teams.Close()

			picks, err := os.Open(loc + dir + "/raw/picks")
			if err == nil {
				scanner := bufio.NewScanner(picks)
				teams := make(map[string][]string)
				for scanner.Scan() {
					line := scanner.Text()
					team := strings.Split(line, ",")[0]
					teams[team] = strings.Split(line, ",")[1:]
				}
				for i := 1; i <= 160; i++ {
					if i%10 == 1 {
						fmt.Println("")
					}
					for k := range teams {
						for _, b := range teams[k] {
							bi, _ := strconv.Atoi(b)
							if bi == i {
								fmt.Printf("%s\t", k)
							}
						}
					}
				}
			}
			defer picks.Close()

		}
	}
}

func populate() {
	db, err := gorm.Open(_db, _dbloc)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

}

func reset() {
	db, err := gorm.Open(_db, _dbloc)

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	// rankings.go
	db.DropTableIfExists(&FantasyFootballers{}, &DraftWizard{})
	// management.go
	db.DropTableIfExists(&League{}, &Settings{}, &Slot{}, &Owner{}, &Transaction{}, &Pick{})
	// roster.go
	db.DropTableIfExists(&Team{}, &Player{})
}

func auto() {
	db, err := gorm.Open(_db, _dbloc)

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	// rankings.go
	db.AutoMigrate(&FantasyFootballers{})
	db.AutoMigrate(&DraftWizard{})
	// management.go
	db.AutoMigrate(&League{})
	db.AutoMigrate(&Owner{})
	db.AutoMigrate(&Transaction{})
	db.AutoMigrate(&Settings{})
	db.AutoMigrate(&Slot{})
	db.AutoMigrate(&Pick{})
	// roster.go
	db.AutoMigrate(&Team{})
	db.AutoMigrate(&Player{})
}

func start() {
	application := Application{}
	application.Load()
	application.StartDraft()
	application.EndDraft()
}

func DefaultSettings() Settings {
	return Settings{Scoring: PPR, Starters: []Slot{
		Slot{Position: QB, Capacity: 1},
		Slot{Position: RB, Capacity: 2},
		Slot{Position: WR, Capacity: 2},
		Slot{Position: TE, Capacity: 1},
		Slot{Position: FLEX, Capacity: 1},
		Slot{Position: DST, Capacity: 1},
		Slot{Position: K, Capacity: 1},
	}, Maximum: []Slot{
		Slot{Position: QB, Capacity: 4},
		Slot{Position: RB, Capacity: 8},
		Slot{Position: WR, Capacity: 8},
		Slot{Position: TE, Capacity: 3},
		Slot{Position: K, Capacity: 3},
		Slot{Position: DST, Capacity: 3},
	}, RosterSize: 16}
}

func (a *Application) LoadKeepers(from_file string) {
	file, _ := ioutil.ReadFile(from_file)
	var keepers []Keeper
	json.Unmarshal(file, &keepers)
	for i := range a.Teams {
		for j := range keepers {
			k := keepers[j]
			player, err := a.FindPlayer(k.Name)
			if err != nil {
				log.Fatal(err)
			}
			(*player).Keeper = true
			if a.Teams[i].Abbrev == k.KeepingTeam {
				a.Teams[i].Roster = append(a.Teams[i].Roster, (*player))
			}
		}
	}
}

func (a *Application) LoadTeams(from_file string) {
	file, _ := ioutil.ReadFile(from_file)
	var board []Team
	json.Unmarshal(file, &board)
	a.Teams = board
}

func (a *Application) FindPlayer(playerName string) (player *Player, err error) {
	for p := range a.Players {
		if a.Players[p].Name == playerName {
			player = &a.Players[p]
			return
		}
	}
	err = errors.New("Could not find the player: " + playerName)
	return
}

func (a *Application) AddPlayer(player Player) {
	find, err := a.FindPlayer(player.Name)
	if err != nil {
		a.Players = append(a.Players, player)
	} else {
		a.MergePlayer(player, find)
	}
}

func (a *Application) MergePlayer(invader Player, target *Player) {
	fmt.Println("Merging Player Data For " + invader.Name)
	if (*target).DW.Empty() {
		fmt.Println("Adding Draft Wizard Data")
		(*target).DW = invader.DW
	}

	if (*target).Ranks.Empty() {
		fmt.Println("Adding FFBallers Rankings")
		(*target).Ranks = invader.Ranks
	}
}
