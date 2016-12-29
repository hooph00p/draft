package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

type Application struct {
	Players  []Player
	Teams    []Team
	Settings Settings
}

func (a *Application) Load() {

}

func (a *Application) StartDraft() {
	fmt.Println("Starting the Draft...")
}

func (a *Application) EndDraft() {
	fmt.Println("Ending the Draft...")
}

func (a *Application) LoadRankings() {
	quarterBacks := LoadRankings(qb_file, QB)
	runningBacks := LoadRankings(rb_file, RB)
	tightEnds := LoadRankings(te_file, TE)
	wideReceivers := LoadRankings(wr_file, WR)

	// Compile Rankings
	var entireRankings []Player
	entireRankings = append(entireRankings, quarterBacks...)
	entireRankings = append(entireRankings, runningBacks...)
	entireRankings = append(entireRankings, tightEnds...)
	entireRankings = append(entireRankings, wideReceivers...)
	for i := range entireRankings {
		a.AddPlayer(entireRankings[i])
	}
}

func (a *Application) LoadDraftWizardAdpFromHtml(from_file string) {
	file, _ := ioutil.ReadFile(from_file)
	reader := strings.NewReader(string(file))
	tok := html.NewTokenizer(reader)
	player := ""
	start := false
	for {
		tt := tok.Next()
		token := tok.Token()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			if token.Data == "tbody" {
				start = true
			}
			if start {
				if token.Data == "tr" {
					if len(player) > 0 {
						add_player := CreatePlayerFromHTMLRow(player)
						fmt.Println("Adding Player: " + player)
						a.AddPlayer(add_player)
					}
					player = ""
				}
			}
		case tt == html.TextToken:
			if start {
				content := strip(token.String())
				if len(content) > 0 {
					if len(player) > 0 {
						player += " | "
					}
					player += content
				}
			}
		}
	}
}
