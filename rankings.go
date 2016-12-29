package main

import "github.com/jinzhu/gorm"

const (
	qb_file = etc_path + "/stat/ffb_qb.csv"
	wr_file = etc_path + "/stat/ffb_wrs.csv"
	rb_file = etc_path + "/stat/ffb_rbs.csv"
	te_file = etc_path + "/stat/ffb_tes.csv"
)

type FantasyFootballers struct {
	gorm.Model
	Consensus int `json:"consensus"`
	Andy      int `json:"andy"`
	Jason     int `json:"jason"`
	Mike      int `json:"mike"`
}

type DraftWizard struct {
	gorm.Model
	Position          int     `json:"position"`
	Overall           int     `json:"overall"`
	Average           float64 `json:"average"`
	High              float64 `json:"high"`
	Low               float64 `json:"low"`
	StandardDeviation float64 `json:"std"`
	PercentDrafted    float64 `json:"draft_pct"`
}

func (dw *DraftWizard) Empty() bool {
	return dw.Overall == 0
}

func (rk *FantasyFootballers) Empty() bool {
	return rk.Consensus == 0
}
