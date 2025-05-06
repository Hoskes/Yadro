package competition

import (
	"fmt"
	"time"
	"yadro-test-assigment/internal"
	"yadro-test-assigment/internal/config"
	"yadro-test-assigment/internal/event"
	"yadro-test-assigment/internal/time_parser"
)

// Competition Представляет собой текущее соревнование
type Competition struct {
	Config      config.Config
	Competitors map[int]*internal.Competitor
	ShotsCount  int
}

// NewCompetition создает новый экземпляр соревнования
func NewCompetition(config config.Config) *Competition {
	return &Competition{
		Config:      config,
		Competitors: make(map[int]*internal.Competitor),
		ShotsCount:  5,
	}
}

// ProcessEvent Дает интерфейс к обработке события
func (receiver *Competition) ProcessEvent(event event.Event) string {
	processResult, err := receiver.process(event)
	if err != nil {
		return err.Error()
	}

	return processResult
}

// Обрабатывает перечень возможных событий
func (receiver *Competition) process(event event.Event) (msg string, err error) {
	competitor := receiver.findCompetitor(event.CompetitorID)
	layout := "15:04:05.000"
	msg = fmt.Sprintf("[%s] ", event.Time.Format(layout))
	switch event.EventID {
	case 1:
		msg += fmt.Sprintf("The competitor(%d) registered", competitor.ID)
		receiver.Competitors[event.CompetitorID] = &internal.Competitor{
			ID:         event.CompetitorID,
			Registered: true,
			Laps:       make([]internal.Lap, 0),
			Penalties:  make([]internal.Lap, 0),
			ShotLines:  internal.NewShotLine(),
		}
	case 2:
		msg += fmt.Sprintf("The start time for the competitor(%d) was set by a draw to %s", competitor.ID, event.ExtraParams)
		//TODO вынести в отдельный модуль
		timeFormat := "15:04:05.000"
		competitor.StartTime, err = time.Parse(timeFormat, event.ExtraParams)
		if err != nil {
			return "", err
		}
		//
	case 3:
		msg += fmt.Sprintf("The competitor(%d) is on the start line", competitor.ID)
	case 4:

		msg += fmt.Sprintf("The competitor(%d) has started", competitor.ID)
		competitor.ActualStartTime = event.Time
		b, err := time_parser.CheckTimeDeviation(competitor.StartTime, event.Time, receiver.Config.StartDelta)
		if err != nil {
			return "", err
		}
		if b {
			competitor.Status = "Not started"
		}
		competitor.AddLap(internal.Lap{
			StartTime: event.Time,
			Length:    float64(receiver.Config.LapLen),
		})

	case 5:
		msg += fmt.Sprintf("The competitor(%d) is on the firing range(%s)", competitor.ID, event.ExtraParams)

		competitor.ShotLines.AddLine(0)

	case 6:
		msg += fmt.Sprintf("The target(%s) has been hit by competitor(%d)", event.ExtraParams, competitor.ID)
		competitor.LineHits++
		competitor.ShotLines.Hits[len(competitor.ShotLines.Hits)-1] += 1
	case 7:
		msg += fmt.Sprintf("The competitor(%d) left the firing range", competitor.ID)

		competitor.AllHits += competitor.LineHits
		competitor.ShotLines.TotalHits += competitor.LineHits

	case 8:
		msg += fmt.Sprintf("The competitor(%d) entered the penalty laps", competitor.ID)

		penaltyLen := competitor.ShotLines.CalcPenaltyLength(receiver.ShotsCount, receiver.Config.PenaltyLen)

		competitor.AddPenalty(internal.Lap{
			StartTime: event.Time,
			Length:    float64(penaltyLen),
		})

	case 9:
		msg += fmt.Sprintf("The competitor(%d) left the penalty laps", competitor.ID)
		//TODO зафиксировать время отбытия

		competitor.GetCurrentPenalty().EndTime = event.Time
		//Фиксируем penaltyStartTime, считаем скорость проезда
	case 10:
		msg += fmt.Sprintf("The competitor(%d) ended the main lap", competitor.ID)
		competitor.LapsCount++
		// Длина известна из конфига, время надо замерить. Status проставлять если нет другого (Started<Finished<Disqualified<NotStarted)
		competitor.GetCurrentLap().EndTime = event.Time
		if competitor.LapsCount >= receiver.Config.Laps {
			if competitor.Status == "" {
				competitor.Status = "Finished"
			}
		} else {
			//КОПИПАСТА ИСПРАВИТЬ
			competitor.AddLap(internal.Lap{
				StartTime: event.Time,
				Length:    float64(receiver.Config.LapLen),
			})
		}
	case 11:
		msg += fmt.Sprintf("The competitor(%d) can`t continue: %s", competitor.ID, event.ExtraParams)
		competitor.Status = "Not finished"
		competitor.GetCurrentLap().EndTime = event.Time
	}
	return msg, err
}

// Ищет или создает нового участника
func (receiver *Competition) findCompetitor(CompetitorID int) *internal.Competitor {
	if receiver.Competitors[CompetitorID] == nil {
		receiver.Competitors[CompetitorID] = &internal.Competitor{
			ID: CompetitorID,
		}
	}
	return receiver.Competitors[CompetitorID]
}
