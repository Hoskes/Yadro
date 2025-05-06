package event

import (
	"time"
)

// Event Структура для хранения событий
type Event struct {
	Time         time.Time
	EventID      int
	CompetitorID int
	ExtraParams  string
}
