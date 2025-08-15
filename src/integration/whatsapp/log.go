package whatsapp

import (
	"log"
	"os"

	"github.com/DanielHGimenez/yell/src/config"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type FileLogger struct {
	logger *log.Logger
	module *string
}

func NewFileLogger(logFilePath string) (*FileLogger, error) {
	logFilePath, err := config.GetFilePathInExecutableFolder(logFilePath)
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	module := "whatsapp"
	logger := log.New(file, module+": ", log.LstdFlags)
	return &FileLogger{logger: logger, module: &module}, nil
}

func (fl *FileLogger) Warnf(msg string, args ...interface{}) {
	fl.logger.Printf("WARNING: "+msg, args...)
}

func (fl *FileLogger) Errorf(msg string, args ...interface{}) {
	fl.logger.Printf("ERROR: "+msg, args...)
}

func (fl *FileLogger) Infof(msg string, args ...interface{}) {
	fl.logger.Printf("INFO: "+msg, args...)
}

func (fl *FileLogger) Debugf(msg string, args ...interface{}) {
	fl.logger.Printf("DEBUG: "+msg, args...)
}

func (fl *FileLogger) Sub(module string) waLog.Logger {
	newModule := *fl.module + ": " + module
	return &FileLogger{
		logger: log.New(fl.logger.Writer(), newModule+": ", fl.logger.Flags()),
		module: &newModule,
	}
}
