package logging

import (
	"os"

	"github.com/code-ready/crc/pkg/crc/constants"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	logfile       *os.File
	LogLevel      string
	originalHooks = logrus.LevelHooks{}
)

func OpenLogFile(path string) (*os.File, error) {
	err := constants.EnsureBaseDirExists()
	if err != nil {
		return nil, err
	}
	l, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func CloseLogFile() {
	logfile.Close()
}

func CloseLogging() {
	CloseLogFile()
	logrus.StandardLogger().ReplaceHooks(make(logrus.LevelHooks))
}

func InitLogrus(logLevel, logFilePath string) {
	logFile, err := OpenLogFile(logFilePath)
	if err != nil {
		logrus.Fatal("Unable to open log file: ", err)
	}
	// send logs to file
	logrus.SetOutput(logFile)

	logrus.SetLevel(logrus.TraceLevel)

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.InfoLevel
	}

	// Add hook to send non error logs to stdout
	logrus.AddHook(newstdOutHook(level, &logrus.TextFormatter{
		// Setting ForceColors is necessary because logrus.TextFormatter determines
		// whether or not to enable colors by looking at the output of the logger.
		// In this case, the output is ioutil.Discard, which is not a terminal.
		// Overriding it here allows the same check to be done, but against the
		// hook's output instead of the logger's output.
		ForceColors:            terminal.IsTerminal(int(os.Stderr.Fd())),
		DisableTimestamp:       true,
		DisableLevelTruncation: false,
	}))

	// Add hook to send error/fatal to stderr
	logrus.AddHook(newstdErrHook(level, &logrus.TextFormatter{
		ForceColors:            terminal.IsTerminal(int(os.Stderr.Fd())),
		DisableTimestamp:       true,
		DisableLevelTruncation: false,
	}))

	for k, v := range logrus.StandardLogger().Hooks {
		originalHooks[k] = v
	}
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Infof(s string, args ...interface{}) {
	logrus.Infof(s, args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Warnf(s string, args ...interface{}) {
	logrus.Warnf(s, args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func Fatalf(s string, args ...interface{}) {
	logrus.Fatalf(s, args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Errorf(s string, args ...interface{}) {
	logrus.Errorf(s, args...)
}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Debugf(s string, args ...interface{}) {
	logrus.Debugf(s, args...)
}
