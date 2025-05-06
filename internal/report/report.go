package report

import (
	"fmt"
	"sort"
	"strconv"
	"time"
	"yadro-test-assigment/internal"
	"yadro-test-assigment/internal/competition"
	"yadro-test-assigment/internal/time_parser"
)

type Report struct {
	Competition competition.Competition
	Report      []ReportObject
}

// GenerateFinalReport Преобразует данные из Competition для генерации отчета
func (report *Report) GenerateFinalReport() *Report {
	for _, comp := range report.Competition.Competitors {
		// Расчет между стартом и последним действием на трассе (финиш круга/потерялся)
		compDuration := comp.Laps[len(comp.Laps)-1].EndTime.Sub(comp.StartTime)
		// Расчет скорости на каждом круге
		lapsDur := report.calcLapDuration(comp, comp.Laps)
		// Расчет скорости на пенальти
		penaltyDur := report.calcLapDuration(comp, comp.Penalties)
		allShots := report.Competition.ShotsCount * report.Competition.Config.Laps
		hitsPerShots := strconv.Itoa(comp.LineHits) + "/" + strconv.Itoa(allShots)
		obj := &ReportObject{
			ID:                comp.ID,
			Duration:          compDuration,
			lapsDuration:      lapsDur,
			penaltiesDuration: penaltyDur,
			HitsPerShots:      hitsPerShots,
			Status:            comp.Status,
		}
		report.Report = append(report.Report, *obj)

	}
	return report
}

type ReportObject struct {
	ID                int
	Status            string
	Duration          time.Duration
	lapsDuration      []LapAndSpeed
	penaltiesDuration []LapAndSpeed
	HitsPerShots      string
}
type ByDuration []ReportObject

func (a ByDuration) Len() int           { return len(a) }
func (a ByDuration) Less(i, j int) bool { return a[i].Duration < a[j].Duration }
func (a ByDuration) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type LapAndSpeed struct {
	lapTime  time.Duration
	lapSpeed float64
}

func (report *Report) calcLapDuration(comp *internal.Competitor, lap []internal.Lap) []LapAndSpeed {
	lapsDur := make([]LapAndSpeed, len(lap))
	for i := 0; i < len(lap); i++ {
		lapTime := lap[i].EndTime.Sub(comp.StartTime)
		lapSp := float64(report.Competition.Config.LapLen) / float64(int(lapTime.Seconds()))
		lapsDur[i] = LapAndSpeed{
			lapTime:  lapTime,
			lapSpeed: lapSp,
		}
	}
	return lapsDur
}

func (report *Report) Show() {
	//	TODO доработать вывод по условиям

	sort.Sort(ByDuration(report.Report))
	for _, reportObject := range report.Report {
		// TODO влепить статус если не закончено
		compTime := time_parser.ParseDurToStr(reportObject.Duration)
		if reportObject.Status != "Finished" {
			compTime = reportObject.Status
		}

		s1 := fmt.Sprintf("[%s] ", compTime)
		s1 += fmt.Sprintf("%d ", reportObject.ID)
		s1 += "[ "
		for _, obj := range reportObject.lapsDuration {
			s1 += fmt.Sprintf("{%s, %.3f} ", time_parser.ParseDurToStr(obj.lapTime), obj.lapSpeed)
		}
		s1 += "] "
		// TODO скобочки
		for _, obj := range reportObject.penaltiesDuration {

			s1 += fmt.Sprintf("{%s, %.3f} ", time_parser.ParseDurToStr(obj.lapTime), obj.lapSpeed)
		}
		s1 += fmt.Sprintf("%s", reportObject.HitsPerShots)

		str := fmt.Sprintf("[ %s ]\n", s1)
		fmt.Print(str)
	}
}
