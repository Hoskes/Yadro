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
		lapsDur := report.calcLapDuration(comp.Laps, report.Competition.Config.LapLen)

		// Расчет скорости на пенальти
		penaltyDur := report.calcLapDuration(comp.Penalties, report.Competition.Config.PenaltyLen)

		// Расчет всех выстрелов производится на основе УЖЕ сделанных выстрелов на площадке
		hitsPerShots := strconv.Itoa(comp.LineHits) + "/" + strconv.Itoa(comp.TotalShots)

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

func (report *Report) calcLapDuration(lap []internal.Lap, lapLength int) []LapAndSpeed {
	lapsDur := make([]LapAndSpeed, len(lap))
	for i := 0; i < len(lap); i++ {
		lapTime := lap[i].EndTime.Sub(lap[i].StartTime)
		lapSp := float64(lapLength) / float64(int(lapTime.Seconds()))
		lapsDur[i] = LapAndSpeed{
			lapTime:  lapTime,
			lapSpeed: lapSp,
		}
	}
	return lapsDur
}

func (report *Report) Show() string {
	//	TODO доработать вывод по условиям

	sort.Sort(ByDuration(report.Report))
	res := ""
	for _, reportObject := range report.Report {
		// TODO влепить статус если не закончено
		compTime := time_parser.ParseDurToStr(reportObject.Duration)
		if reportObject.Status != "Finished" {
			compTime = reportObject.Status
		}

		s1 := fmt.Sprintf("[%s] ", compTime)
		s1 += fmt.Sprintf("%d ", reportObject.ID)
		s1 += "[ "
		for i, obj := range reportObject.lapsDuration {
			s1 += fmt.Sprintf("{%s, %.3f}", time_parser.ParseDurToStr(obj.lapTime), obj.lapSpeed)
			if i != len(reportObject.lapsDuration)-1 {
				s1 += ", "
			}
		}
		s1 += "] "
		// TODO скобочки
		s3 := ""
		for i, obj := range reportObject.penaltiesDuration {

			s3 += fmt.Sprintf("{%s, %.3f}", time_parser.ParseDurToStr(obj.lapTime), obj.lapSpeed)
			if i != len(reportObject.penaltiesDuration)-1 {
				s3 += ", "
			}
		}
		s1 += fmt.Sprintf("{%s} ", s3)
		s1 += fmt.Sprintf("%s", reportObject.HitsPerShots)

		str := fmt.Sprintf("%s \n", s1)
		res += str
	}
	return res
}
