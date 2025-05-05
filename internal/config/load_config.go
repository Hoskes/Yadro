package config

import (
	"encoding/json"
	"log"
	"os"
)

func LoadJSONConfig(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Ошибка при обработке:", err)
	}
	var conf Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatal("Ошибка при парсинге:", err)
	}
	closeErr := file.Close()
	if closeErr != nil {
		log.Fatal("Ошибка при закрытии:", err)
	}
	return conf, nil
}
