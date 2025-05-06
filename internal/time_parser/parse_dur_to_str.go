package time_parser

import (
	"fmt"
	"time"
)

func ParseDurToStr(duration time.Duration) string {
	// Получаем часы, минуты и секунды
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	milliseconds := int(duration.Milliseconds()) % 1000

	// Форматируем строку
	timeString := fmt.Sprintf("%02d:%02d:%02d.%03d", hours, minutes, seconds, milliseconds)

	// Выводим результат
	return timeString // Вывод: 10:32:22.000
}
