package logger

import (
	"log"
)

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string, fields ...interface{}) {
	log.Println("[INFO]", msg, fields)
}

func (l *Logger) Error(msg string, fields ...interface{}) {
	log.Println("[ERROR]", msg, fields)
}

func (l *Logger) Fatal(msg string, fields ...interface{}) {
	log.Fatalln("[FATAL]", msg, fields)
}
