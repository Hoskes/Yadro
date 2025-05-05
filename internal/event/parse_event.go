package event

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func ParseEvent(str string) (Event, error) {
	args := strings.Split(str, " ")
	var event Event
	if len(args) < 3 {
		return event, fmt.Errorf("Нарушено условие минимального кол-ва аргументов")
	}

	timeFormat := "15:04:05.000"
	t, err := time.Parse(timeFormat, strings.Trim(args[0], "[]"))
	if err != nil {
		log.Fatal("Ошибка при парсинге Time:", err)
		return event, err
	}

	eventID, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal("Ошибка парсинга eventID")
		return event, err
	}

	competitorID, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal("Ошибка парсинга eventID")
		return event, err
	}
	event = Event{
		Time:         t,
		EventID:      eventID,
		CompetitorID: competitorID,
	}
	if len(args) > 3 {
		event.ExtraParams = args[3]
	}
	return event, nil
}
