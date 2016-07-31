package ply

import (
	"log"
	"os"
)

func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", 0),
	}
}

type Logger struct {
	*log.Logger
}

func (l *Logger) Info(output string) {
	l.Println(output)
}

func (l *Logger) Error(output string, err error) {
	l.Printf("err: %s %s", output, err)
}
