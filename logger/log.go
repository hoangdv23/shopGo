package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func PrintLog() (*os.File, error) {
	// Mở file log
	filePath := os.Getenv("LOG_FILE_PATH")
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	// Cấu hình logger toàn cục
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	return file, nil
}
