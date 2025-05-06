package main

import (
	"bufio"
	"fmt"
	"log"
	report2 "yadro-test-assigment/internal/report"

	. "yadro-test-assigment/internal/competition"
	. "yadro-test-assigment/internal/config"
	. "yadro-test-assigment/internal/event"
)

func main() {

	conf, err := LoadJSONConfig("internal/test/config.json")
	if err != nil {
		log.Fatal("Ошибка при обработке JSON-conf", err)
	}

	competition := NewCompetition(conf)
	file := LoadEventFile("internal/test/events")

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		t, parseErr := ParseEvent(scan.Text())
		fmt.Println(competition.ProcessEvent(t))
		if parseErr != nil {
			log.Fatal("Ошибка при парсинге события:", parseErr)
		}
	}
	report := report2.Report{
		Competition: *competition,
	}
	report.GenerateFinalReport().Show()
	// TODO добавить вывод в файлы

}
