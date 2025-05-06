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
	report := report2.Report{
		Competition: *competition,
	}
	report.GenerateFinalReport()
	report.Show()
	// нужно потом убрать
	//for _, competitor := range competition.Competitors {
	//	//fmt.Println(competitor)
	//	fmt.Println("ID", competitor.ID)
	//	fmt.Println("Shots:", competitor.ShotLines)
	//	fmt.Println("Status:", competitor.Status)
	//	fmt.Println("Hits:", competitor.LineHits)
	//	fmt.Println("Laps:", competitor.Laps)
	//	fmt.Println("Penalties", competitor.Penalties)
	//	fmt.Println("Comment:", competitor.Comment)
	//}
}
