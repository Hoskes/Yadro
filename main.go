package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	. "yadro-test-assigment/models"
)

func main() {

	conf, err := loadJSONConfig("sunny_5_skiers/config.json")
	if err != nil {
		log.Fatal("Ошибка при обработке JSON-conf", err)
	}
	competition := NewCompetition(conf)
	file := loadEventFile("sunny_5_skiers/events")
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		t, parseErr := parseEvent(scan.Text())
		fmt.Println(competition.ProcessEvent(t))
		//competition.Events = append(competition.Events, t)
		if parseErr != nil {
			log.Fatal("Ошибка при парсинге события:", parseErr)
		}
	}
	for _, competitor := range competition.Competitors {
		fmt.Println(competitor)
	}
}
func parseEvent(str string) (Event, error) {
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
func loadJSONConfig(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Ошибка при обработке:", err)
	}
	var conf Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatal("Ошибка при парсинге:", err)
	}
	closeErr := file.Close()
	if closeErr != nil {
		log.Fatal("Ошибка при закрытии:", err)
	}
	return conf, nil
}

// Загружает файл событий
func loadEventFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Ошибка при обработке event.txt", err)

	}
	return file
}
