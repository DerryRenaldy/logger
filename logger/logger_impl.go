package logger

import (
	"fmt"
	"github.com/DerryRenaldy/logger/constant"
	log "github.com/rs/zerolog"
	"os"
	"runtime"
	"strings"
)

type Logger struct {
	service string
	env     string
	log     *log.Logger
}

var Log *Logger

func New(serviceName, env, logLevel string) *Logger {
	Log = Init(serviceName, env, logLevel)
	return Log
}

func (l Logger) Debug(msg string) {
	pc, file, line, _ := runtime.Caller(1)

	logger := l.log.With().
		Str("service", l.service).
		Str("env", l.env).
		Str("file", file).
		Str("function",
			runtime.FuncForPC(pc).Name()).
		Int("line", line).Logger()

	if strings.EqualFold(l.env, constant.EnvDev) || strings.EqualFold(l.env, constant.EnvStage) {
		logger.Debug().Msg(msg)
	} else {
		logger.Info().Msg(msg)
	}
}

func (l Logger) Info(msg string) {
	pc, file, line, _ := runtime.Caller(1)

	logger := l.log.With().
		Str("service", l.service).
		Str("env", l.env).
		Str("file", file).
		Str("function", runtime.FuncForPC(pc).Name()).
		Int("line", line).Logger()

	logger.Info().Msg(msg)
}

func (l Logger) Warn(msg string) {
	pc, file, line, _ := runtime.Caller(1)

	logger := l.log.With().
		Str("service", l.service).
		Str("env", l.env).
		Str("file", file).
		Str("function", runtime.FuncForPC(pc).Name()).
		Int("line", line).Logger()

	logger.Warn().Msg(msg)
}

func (l Logger) Error(msg string) {
	pc, file, line, _ := runtime.Caller(1)

	logger := l.log.With().
		Str("service", l.service).
		Str("env", l.env).
		Str("file", file).
		Str("function", runtime.FuncForPC(pc).Name()).
		Int("line", line).Logger()

	logger.Warn().Msg(msg)
}

func (l Logger) Fatal(msg string) {
	pc, file, line, _ := runtime.Caller(1)

	logger := l.log.With().
		Str("service", l.service).
		Str("env", l.env).
		Str("file", file).
		Str("function", runtime.FuncForPC(pc).Name()).
		Int("line", line).Logger()

	logger.Fatal().Msg(msg)
}

func (l Logger) Panic(msg string) {
	pc, file, line, _ := runtime.Caller(1)

	logger := l.log.With().
		Str("service", l.service).
		Str("env", l.env).
		Str("file", file).
		Str("function", runtime.FuncForPC(pc).Name()).
		Int("line", line).Logger()

	logger.Panic().Msg(msg)
}

func (l Logger) Debugf(format string, args ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)

	logger := l.log.With().
		Str("service", l.service).
		Str("env", l.env).
		Str("file", file).
		Str("function", runtime.FuncForPC(pc).Name()).
		Int("line", line).Logger()

	formatMsg := fmt.Sprintf(format, args...)

	if strings.EqualFold(l.env, constant.EnvDev) || strings.EqualFold(l.env, constant.EnvStage) {
		logger.Debug().Msg(formatMsg)
	} else {
		logger.Info().Msg(formatMsg)
	}
}

func (l Logger) Infof(format string, args ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)

	logger := l.log.With().
		Str("service", l.service).
		Str("env", l.env).
		Str("file", file).
		Str("function", runtime.FuncForPC(pc).Name()).
		Int("line", line).Logger()

	formatMsg := fmt.Sprintf(format, args...)
	logger.Info().Msg(formatMsg)
}

func (l Logger) Warnf(format string, args ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)

	logger := l.log.With().
		Str("service", l.service).
		Str("env", l.env).
		Str("file", file).
		Str("function", runtime.FuncForPC(pc).Name()).
		Int("line", line).Logger()

	formatMsg := fmt.Sprintf(format, args...)
	logger.Warn().Msg(formatMsg)
}

func (l Logger) Errorf(format string, args ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)

	logger := l.log.With().
		Str("service", l.service).
		Str("env", l.env).
		Str("file", file).
		Str("function", runtime.FuncForPC(pc).Name()).
		Int("line", line).Logger()

	formatMsg := fmt.Sprintf(format, args...)
	logger.Error().Msg(formatMsg)
}

func (l Logger) Fatalf(format string, args ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)

	logger := l.log.With().
		Str("service", l.service).
		Str("env", l.env).
		Str("file", file).
		Str("function", runtime.FuncForPC(pc).Name()).
		Int("line", line).Logger()

	formatMsg := fmt.Sprintf(format, args...)
	logger.Fatal().Msg(formatMsg)
}

func (l Logger) Panicf(format string, args ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)

	logger := l.log.With().
		Str("service", l.service).
		Str("env", l.env).
		Str("file", file).
		Str("function", runtime.FuncForPC(pc).Name()).
		Int("line", line).Logger()

	formatMsg := fmt.Sprintf(format, args...)
	logger.Panic().Msg(formatMsg)
}

func Init(service, env, logLevel string) *Logger {
	ll := log.New(os.Stdout).With().Timestamp().Logger()

	if logLevel == "" {
		logLevel = "error"
	}
	lv, err := log.ParseLevel(logLevel)
	if err != nil {
		lv = log.ErrorLevel
	}

	// Only log the warning severity or above.
	log.SetGlobalLevel(lv)

	return &Logger{
		service: service,
		env:     env,
		log:     &ll,
	}
}
