package initializers

import (
	"log"
	"os"

	"github.com/yutt/go-movies-api/logger"
)

func CreateLogs() {
	logger.Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
