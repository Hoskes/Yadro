package models

import "time"

// Competitor / Представляет собой участника соревнования
type Competitor struct {
	ID               int
	Registered       bool
	StartTime        time.Time
	ActualStartTime  time.Time
	LapTimes         []time.Duration
	PenaltyTime      []time.Duration
	Hits             int
	Status           string // Следующие статусы должны быть в обработке "NotStarted", "NotFinished", "Finished", "Disqualified"
	Comment          string
	PenaltyStartTime time.Time
	LapsCount        int
}
