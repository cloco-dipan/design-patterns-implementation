package logger

import (
	"encoding/json"
	"log"
	"os"
)

type FileLogger struct {
}

// Log writes to file
func (fl *FileLogger) Log(msg interface{}) {
	
	f, err := os.OpenFile(fl.GetFileLocation(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// convert to string
	jsonStr, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	jsonStr = append(jsonStr, '\n')

	if _, err := f.Write([]byte(string(jsonStr))); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func (fl *FileLogger) GetFileLocation() string {
	return FILE_LOGGER_PATH
}
