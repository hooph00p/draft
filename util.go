package main

import (
	"encoding/csv"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func strip(s string) (b string) {
	b = strings.TrimSpace(s)
	return
}

func LoadRankings(file string, position string) (players []Player) {
	csvfile, _ := os.Open(file)
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1
	rawCSVdata, _ := reader.ReadAll()
	players = make([]Player, len(rawCSVdata))
	for i, each := range rawCSVdata {
		name := ""
		team := ""
		bye := 0
		player_name_line := strings.Split(each[0], " ")
		if len(player_name_line) > 1 { // prevents empty row sets (thanks ffballers ;])
			if strings.Contains(each[0], ", FA") {
				name = strings.Replace(each[0], ", FA", "", -1)
				team = "FA"
				bye = 0
			} else {
				for split := range player_name_line {
					if split < len(player_name_line)-2 { // name
						if split > 0 {
							name += " "
						}
						name += player_name_line[split]
					}
					if split == len(player_name_line)-2 { // team
						team = player_name_line[split]
					}
					if split == len(player_name_line)-1 { // bye
						bye, _ = strconv.Atoi(player_name_line[split][1 : len(player_name_line[split])-1])
					}
				}
			}
			consensus, _ := strconv.Atoi(each[1])
			andy, _ := strconv.Atoi(each[2])
			jason, _ := strconv.Atoi(each[3])
			mike, _ := strconv.Atoi(each[4])
			players[i] = Player{Name: name, Bye: bye, NflTeam: team, Position: position, Ranks: FantasyFootballers{Consensus: consensus, Andy: andy, Jason: jason, Mike: mike}}
		}
	}
	return
}

func CreatePlayerFromHTMLRow(row string) (player Player) {
	const (
		POSITION_RANK      = 0
		OVERALL            = 1
		NAME               = 2
		TEAM               = 3
		BYE                = 4
		ADP                = 5
		HIGH               = 6
		LOW                = 7
		STANDARD_DEVIATION = 8
		DRAFT_PERCENT      = 9
	)
	columns := strings.Split(row, "|")

	player = Player{DW: DraftWizard{}}

	for i := range columns {
		column := strip(columns[i])
		switch i {
		case POSITION_RANK:
			player.Position = regexp.MustCompile("[a-zA-Z]+").FindString(column)
			player.DW.Position, _ = strconv.Atoi(regexp.MustCompile("[0-9]+").FindString(column))
		case OVERALL:
			player.DW.Overall, _ = strconv.Atoi(column)
		case NAME:
			player.Name = column
		case TEAM:
			player.NflTeam = strings.ToUpper(column)
		case BYE:
			player.Bye, _ = strconv.Atoi(strings.Replace(strings.Replace(column, ")", "", -1), "(", "", -1))
		case ADP:
			player.DW.Average, _ = strconv.ParseFloat(column, 32)
		case HIGH:
			player.DW.High, _ = strconv.ParseFloat(column, 32)
		case LOW:
			player.DW.Low, _ = strconv.ParseFloat(column, 32)
		case DRAFT_PERCENT:
			pct, _ := strconv.ParseFloat(strings.Replace(column, "%", "", -1), 32)
			player.DW.PercentDrafted = pct / 100.0
		case STANDARD_DEVIATION:
			player.DW.StandardDeviation, _ = strconv.ParseFloat(column, 32)
		}
	}
	return
}
