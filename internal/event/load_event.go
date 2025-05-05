package event

import (
	"log"
	"os"
)

// LoadEventFile Загружает файл событий
func LoadEventFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Ошибка при обработке event.txt", err)

	}
	return file
}
