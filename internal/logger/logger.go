package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() *os.File {
	// создаем папку logs если нет
	logDir := "logs"
	_ = os.MkdirAll(logDir, os.ModePerm)

	// имя файла по дате
	fileName := time.Now().Format("2006-01-02") + ".log"
	filePath := filepath.Join(logDir, fileName)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Ошибка открытия лог файла: %v", err)
	}

	// пишем и в консоль, и в файл
	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	return file
}
