package create_file

import (
	"fmt"
	"os"
)

// FileCreateAndWrite Создает и заполняет файлы для вывода
func FileCreateAndWrite(path, msg string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", path, err)
		return nil, err
	}
	_, err = file.WriteString(msg)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Ошибка закрытия файла", err)
		}
	}(file)
	return file, nil
}
