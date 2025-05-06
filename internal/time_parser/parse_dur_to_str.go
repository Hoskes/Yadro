package time_parser

import (
	"fmt"
	"time"
)

// ParseDurToStr Парсит длительность из миллисекунд в формат 00:00:00.000
func ParseDurToStr(duration time.Duration) string {
	//Распил времени на составные
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	milliseconds := duration % time.Second / time.Millisecond

	timeString := fmt.Sprintf("%02d:%02d:%02d.%03d", hours, minutes, seconds, milliseconds)
	return timeString
}
