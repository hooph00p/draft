package main

type ESPNTEAM struct {
	Id        int    `csv:"id"`
	TeamName  string `csv:"team_name"`
	Division  string `csv:"division"`
	OwnerName string `csv:"owner"`
	Email     string `csv:"email"`
}

type YAHOOTEAM struct {
}
