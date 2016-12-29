package main

type ScoringScheme string

const (
	PPR      ScoringScheme = "PPR"
	Standard ScoringScheme = "Standard"
)

const (
	CONSENSUS = 0 + iota
	ANDY
	JASON
	MIKE
)

const (
	QB   string = "QB"
	RB          = "RB"
	WR          = "WR"
	TE          = "TE"
	FLEX        = "FLEX"
	DST         = "DST"
	K           = "K"
)
