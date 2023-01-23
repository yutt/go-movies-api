package logger

import (
	"log"
	"os"
)

var (
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
)
