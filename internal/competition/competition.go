package competition

import (
	"errors"
	"fmt"
	"time"
	"yadro-test-assigment/internal"
	"yadro-test-assigment/internal/config"
	"yadro-test-assigment/internal/event"
)

// Competition Представляет собой текущее соревнование
type Competition struct {
	Config      config.Config
	Competitors map[int]*internal.Competitor
	Events      []event.Event
}

// NewCompetition создает новый экземпляр соревнования
func NewCompetition(config config.Config) *Competition {
	return &Competition{
		Config:      config,
		Competitors: make(map[int]*internal.Competitor),
		Events:      make([]event.Event, 0),
	}
}
func (receiver *Competition) ProcessEvent(event event.Event) string {
	//TODO обработать логику из md

	processResult, err := receiver.process(event)
	if err != nil {
		return err.Error()
	}
	receiver.Events = append(receiver.Events, event) //в конце запихнуть в логи
	return processResult
}

// Обрабатывает перечень возможных событий
func (receiver *Competition) process(event event.Event) (msg string, err error) {
	competitor := receiver.findCompetitor(event.CompetitorID)

	switch event.EventID {
	case 1:
		msg = fmt.Sprintf("The competitor(%d) registered", competitor.ID)
		receiver.Competitors[event.CompetitorID] = &internal.Competitor{
			ID:         event.CompetitorID,
			Registered: true,
		}
	case 2:
		msg = fmt.Sprintf("The start time for the competitor(%d) was set by a draw to %s", competitor.ID, event.ExtraParams)
		timeFormat := "15:04:05.000"
		competitor.StartTime, err = time.Parse(timeFormat, event.ExtraParams)
		if err != nil {
			return "", err
		}

	case 3:
		msg = fmt.Sprintf("The competitor(%d) is on the start line", competitor.ID)
	case 4:

		msg = fmt.Sprintf("The competitor(%d) has started", competitor.ID)
		competitor.ActualStartTime = event.Time
		b, err := checkTimeDeviation(competitor.StartTime, event.Time, receiver.Config.StartDelta)
		if err != nil {
			return "", err
		}
		if b {
			competitor.Status = "Not started"
		}
	case 5:
		msg = fmt.Sprintf("The competitor(%d) is on the firing range(%s)", competitor.ID, event.ExtraParams)
		//TODO зафиксировать время прибытия
	case 6:
		msg = fmt.Sprintf("The target(%s) has been hit by competitor(%d)", event.ExtraParams, competitor.ID)
		competitor.Hits++

	case 7:
		msg = fmt.Sprintf("The competitor(%d) left the firing range", competitor.ID)
		//TODO зафиксировать время отбытия
	case 8:
		msg = fmt.Sprintf("The competitor(%d) entered the penalty laps", competitor.ID)
		//TODO зафиксировать время прибытия
	case 9:
		msg = fmt.Sprintf("The competitor(%d) left the penalty laps", competitor.ID)
		//TODO зафиксировать время отбытия
	case 10:
		msg = fmt.Sprintf("The competitor(%d) ended the main lap", competitor.ID)
		if competitor.LapsCount >= receiver.Config.Laps {
			competitor.Status = "Finished"
		} else {
			competitor.LapsCount++
		}
	case 11:
		msg = fmt.Sprintf("The competitor(%d) can`t continue: %s", competitor.ID, event.ExtraParams)
		competitor.Status = "Not finished"
	}
	return msg, err
}

func checkTimeDeviation(startTime time.Time, endTime time.Time, delta string) (bool, error) {

	duration, err := parseDuration(delta)
	if err != nil {
		return false, err
	}
	return endTime.Sub(startTime) > duration, nil
}

// Функция для парсинга длительности из строки формата "00:00:00"
// TODO Вытяни в сериализатор/десериализатор куда-нибудь отдельно
func parseDuration(s string) (time.Duration, error) {
	var hours, minutes, seconds int
	_, err := fmt.Sscanf(s, "%02d:%02d:%02d", &hours, &minutes, &seconds)
	if err != nil {
		return 0, errors.New("Parse duration error")
	}
	return time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second, nil
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
