package competition

import (
	"yadro-test-assigment/internal"
	"yadro-test-assigment/internal/config"
)

// Competition Представляет собой текущее соревнование по биатлону
type Competition struct {
	Config      config.Config
	Competitors map[int]*internal.Competitor
	ShotsCount  int
}

// NewCompetition создает новый экземпляр соревнования
func NewCompetition(config config.Config, shotsCount int) *Competition {
	return &Competition{
		Config:      config,
		Competitors: make(map[int]*internal.Competitor),
		ShotsCount:  shotsCount, //Вне
	}
}
