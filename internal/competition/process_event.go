package competition

import (
	"fmt"
	"time"
	"yadro-test-assigment/internal"
	"yadro-test-assigment/internal/event"
	"yadro-test-assigment/internal/time_parser"
)

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
		// Создание нового участника в заезде
		receiver.Competitors[event.CompetitorID] = &internal.Competitor{
			ID:         event.CompetitorID,
			Registered: true,
			Laps:       make([]internal.Lap, 0),
			Penalties:  make([]internal.Lap, 0),
			ShotLines:  internal.NewShotLine(),
		}
	case 2:
		msg += fmt.Sprintf("The start time for the competitor(%d) was set by a draw to %s", competitor.ID, event.ExtraParams)

		// Фиксация времени старта
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

		//Фиксация реального времени старта
		competitor.ActualStartTime = event.Time
		b, err := time_parser.CheckTimeDeviation(competitor.StartTime, event.Time, receiver.Config.StartDelta)
		if err != nil {
			return "", err
		}
		if b {
			competitor.Status = "Not started"
			//Удаление данные о заезде при дисквалификации
			deleteLastMoves(competitor)
		} else {
			competitor.AddLap(internal.Lap{
				StartTime: event.Time,
				Length:    float64(receiver.Config.LapLen),
			})
		}
	case 5:
		msg += fmt.Sprintf("The competitor(%d) is on the firing range(%s)", competitor.ID, event.ExtraParams)
		// Добавление новой линии для поражения мишеней
		competitor.ShotLines.AddLine(0)
	case 6:
		msg += fmt.Sprintf("The target(%s) has been hit by competitor(%d)", event.ExtraParams, competitor.ID)

		//Фиксация попаданий
		competitor.LineHits++
		competitor.ShotLines.Hits[len(competitor.ShotLines.Hits)-1] += 1
	case 7:
		msg += fmt.Sprintf("The competitor(%d) left the firing range", competitor.ID)

		//Фиксация попаданий
		competitor.AllHits += competitor.LineHits
		competitor.ShotLines.TotalHits += competitor.LineHits
		competitor.TotalShots += receiver.ShotsCount
	case 8:
		msg += fmt.Sprintf("The competitor(%d) entered the penalty laps", competitor.ID)

		//Расчет длины пенальти
		penaltyLen := competitor.ShotLines.CalcPenaltyLength(receiver.ShotsCount, receiver.Config.PenaltyLen)

		//Создание структуры для хранения данных о пенальти
		competitor.AddPenalty(internal.Lap{
			StartTime: event.Time,
			Length:    float64(penaltyLen),
		})

	case 9:
		msg += fmt.Sprintf("The competitor(%d) left the penalty laps", competitor.ID)

		//Фиксация времени выезда из зоны пенальти
		competitor.GetCurrentPenalty().EndTime = event.Time

	case 10:
		msg += fmt.Sprintf("The competitor(%d) ended the main lap", competitor.ID)

		competitor.LapsCount++
		competitor.GetCurrentLap().EndTime = event.Time
		if competitor.LapsCount >= receiver.Config.Laps {
			if competitor.Status == "" {
				competitor.Status = "Finished"
			}
		} else {
			competitor.AddLap(internal.Lap{
				StartTime: event.Time,
				Length:    float64(receiver.Config.LapLen),
			})
		}
	case 11:
		msg += fmt.Sprintf("The competitor(%d) can`t continue: %s", competitor.ID, event.ExtraParams)

		competitor.Status = "Not finished"
		// Удалить заход на новый круг или пенальти если потерялся в процессе
		deleteLastMoves(competitor)

	}
	return msg, err
}

// Удаляет данные о последнем пройденном участке
func deleteLastMoves(competitor *internal.Competitor) {
	if competitor.Laps != nil {
		if competitor.GetCurrentLap() != nil {
			if (competitor.GetCurrentLap().EndTime == time.Time{}) {
				competitor.DeleteLastLap()
			}
		}
	}
	if competitor.Penalties != nil {
		if competitor.GetCurrentPenalty() != nil {

			if (competitor.GetCurrentPenalty().EndTime == time.Time{}) {
				competitor.DeleteLastPenalty()
			}
		}
	}
}
