package main

import (
	"bufio"
	"fmt"
	"log"
	. "yadro-test-assigment/internal"
	. "yadro-test-assigment/internal/competition"
	. "yadro-test-assigment/internal/config"
	. "yadro-test-assigment/internal/event"
)

func main() {

	conf, err := LoadJSONConfig("sunny_5_skiers/config.json")
	if err != nil {
		log.Fatal("Ошибка при обработке JSON-conf", err)
	}

	competition := NewCompetition(conf)
	file := LoadEventFile("sunny_5_skiers/events")

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		t, parseErr := ParseEvent(scan.Text())
		fmt.Println(competition.ProcessEvent(t))
		if parseErr != nil {
			log.Fatal("Ошибка при парсинге события:", parseErr)
		}
	}

	// нужно потом убрать
	for _, competitor := range competition.Competitors {
		fmt.Println(competitor)
	}
}
