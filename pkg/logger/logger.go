package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// Interface -.
type Logger interface {
	Debugf(format interface{}, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format interface{}, args ...interface{})
	Fatalf(format interface{}, args ...interface{})
}

type logger struct {
	logger *zerolog.Logger
}

var _ Logger = (*logger)(nil)

// New -.
func New(level string) Logger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	skipFrameCount := 3
	//logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()
	lg := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &logger{
		logger: &lg,
	}
}

// Debug -.
func (l *logger) Debugf(format interface{}, args ...interface{}) {
	l.msg("debug", format, args...)
}

// Info -.
func (l *logger) Infof(format string, args ...interface{}) {
	l.log(format, args...)
}

// Warn -.
func (l *logger) Warnf(format string, args ...interface{}) {
	l.log(format, args...)
}

// Error -.
func (l *logger) Errorf(format interface{}, args ...interface{}) {
	if l.logger.GetLevel() == zerolog.DebugLevel {
		l.Debugf(format, args...)
	}

	l.msg("error", format, args...)
}

// Fatal -.
func (l *logger) Fatalf(format interface{}, args ...interface{}) {
	l.msg("fatal", format, args...)

	os.Exit(1)
}

func (l *logger) log(format string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Info().Msg(format)
	} else {
		l.logger.Info().Msgf(format, args...)
	}
}

func (l *logger) msg(level string, format interface{}, args ...interface{}) {
	switch msg := format.(type) {
	case error:
		l.log(msg.Error(), args...)
	case string:
		l.log(msg, args...)
	default:
		l.log(fmt.Sprintf("%s format %v has unknown type %v", level, format, msg), args...)
	}
}
