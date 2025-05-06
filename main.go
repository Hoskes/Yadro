package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"yadro-test-assigment/internal/create_file"
	report2 "yadro-test-assigment/internal/report"

	. "yadro-test-assigment/internal/competition"
	. "yadro-test-assigment/internal/config"
	. "yadro-test-assigment/internal/event"
)

func main() {
	const shotsCount int = 5 //Взято из регламента проведения соревнований

	// Задать параметры ввода-вывода
	jsonConfig := flag.String("config", "sunny_5_skiers/config.json", "Path to JSON-config file")
	inputPath := flag.String("input", "sunny_5_skiers/events", "Path to input events file")
	pathOutput := flag.String("output", "output.txt", "Path to log file")
	pathResult := flag.String("result", "result.txt", "Path to result file")
	flag.Parse()

	// Подгрузка файла конфигурации
	conf, err := LoadJSONConfig(*jsonConfig)
	if err != nil {
		log.Fatal("Ошибка при обработке JSON-conf", err)
	}

	// Создание сущности соревнования
	competition := NewCompetition(conf, shotsCount)
	file := LoadEventFile(*inputPath)

	scan := bufio.NewScanner(file)
	logs := ""

	fmt.Println("Логи соревнований:")
	for scan.Scan() {
		t, parseErr := ParseEvent(scan.Text())
		if parseErr != nil {
			log.Fatal("Ошибка при парсинге события:", parseErr)
		}
		str := competition.ProcessEvent(t)
		logs += str + "\n"
		fmt.Println(str)

	}

	_, err = create_file.FileCreateAndWrite(*pathOutput, logs)
	if err != nil {
		log.Fatal("Ошибка при создании файла логов", err)
	}
	report := report2.Report{
		Competition: *competition,
	}
	result := report.GenerateFinalReport().Show()
	fmt.Println("\nДанные о результатах соревнований:")
	fmt.Println(result)
	_, err = create_file.FileCreateAndWrite(*pathResult, result)
	if err != nil {
		log.Fatal("Ошибка при создании файла результатов", err)
	}
}
