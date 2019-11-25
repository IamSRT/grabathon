package logger

import (
	log "github.com/sirupsen/logrus"
	"grabathon/config"
	"path"
	"runtime"
)

// init the logger
func init()  {
	// Setup hook
	log.AddHook(ContextHook{})

	// Set log formatter
	log.SetFormatter(&log.JSONFormatter{})

	// Set log level
	logLevel, _ := log.ParseLevel(config.LogLevel)
	log.SetLevel(logLevel)
}

// ContextHook ...
type ContextHook struct{}

// Levels ...
func (hook ContextHook) Levels() []log.Level {
	return log.AllLevels
}

// Fire ...
func (hook ContextHook) Fire(entry *log.Entry) error {
	if _, file, line, ok := runtime.Caller(10); ok {
		entry.Data["file"] = path.Base(file)
		entry.Data["line"] = line
	}
	return nil
}