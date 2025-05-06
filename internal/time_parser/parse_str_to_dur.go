package time_parser

import (
	"errors"
	"fmt"
	"time"
)

// ParseStrToDuration Функция для парсинга длительности из строки формата  "00:00:00.000"
func ParseStrToDuration(s string) (time.Duration, error) {
	var hours, minutes, seconds, milliseconds int
	_, err := fmt.Sscanf(s, "%02d:%02d:%02d.%03d", &hours, &minutes, &seconds, &milliseconds)
	if err != nil {
		_, err = fmt.Sscanf(s, "%02d:%02d:%02d", &hours, &minutes, &seconds)
		if err != nil {
			return 0, errors.New("parse duration error")
		}
	}
	return time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second +
		time.Duration(milliseconds)*time.Millisecond, nil
}
