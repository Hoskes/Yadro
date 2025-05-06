package internal

import (
	"time"
)

// Competitor представляет собой участника соревнования (биатлониста)
type Competitor struct {
	ID              int
	Registered      bool
	StartTime       time.Time
	ActualStartTime time.Time
	Laps            []Lap
	Penalties       []Lap
	ShotLines       ShotLine
	AllHits         int
	LineHits        int    // Надо для расчета пенальти
	Status          string // Следующие статусы должны быть в обработке "NotStarted", "NotFinished", "Finished", "Disqualified"
	Comment         string
	LapsCount       int
}
type Lap struct {
	StartTime time.Time
	EndTime   time.Time
	Length    float64
}

func (receiver Lap) CalcSpeed() float64 {
	duration := receiver.EndTime.Sub(receiver.StartTime)
	return receiver.Length / duration.Seconds()
}
func (receiver *Competitor) AddLap(lap Lap) {
	receiver.Laps = append(receiver.Laps, lap)
}
func (receiver *Competitor) AddPenalty(lap Lap) {
	receiver.Penalties = append(receiver.Penalties, lap)
}
func (receiver *Competitor) GetCurrentLap() *Lap {
	return &receiver.Laps[len(receiver.Laps)-1]
}

func (receiver *Competitor) GetCurrentPenalty() *Lap {
	return &receiver.Penalties[len(receiver.Penalties)-1]
}

type ShotLine struct {
	TotalHits int
	Hits      []int
}

func (receiver *ShotLine) CalcPenaltyLength(shotsPerLine, penaltyLen int) int {
	return (shotsPerLine - receiver.Hits[len(receiver.Hits)-1]) * penaltyLen
}
func (receiver *ShotLine) AddLine(pHits int) {
	receiver.Hits = append(receiver.Hits, pHits)
}
func NewShotLine() ShotLine {
	return ShotLine{
		TotalHits: 0,
		Hits:      make([]int, 0),
	}
}
