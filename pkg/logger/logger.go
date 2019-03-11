package logger

import (
	"fmt"
	"time"
	"travel/internal/pkg/config"
)

//Logger ...
type Logger struct {
	Config *config.Config
}

//NewLogger ...
func NewLogger(config *config.Config) *Logger {
	return &Logger{Config: config}
}

//AddDebugLog ...
func (l *Logger) AddDebugLog(str string) {
	fmt.Printf("[%s] [DEBUG] [%s] %s\n", time.Now().UTC().Format("2006-01-02T15:04:05.999Z"), l.Config.EnvLevel, str)
}
