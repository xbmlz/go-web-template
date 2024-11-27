package logger

import (
	"os"

	"github.com/fatih/color"
	"github.com/xbmlz/go-web-template/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Init(c *config.Config) {
	// First, define our level-handling logic.
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		if c.Server.IsDev() {
			return lvl < zapcore.ErrorLevel
		} else {
			return lvl >= zapcore.InfoLevel && lvl < zapcore.ErrorLevel
		}
	})

	// Directly output to stdout and stderr, and add caller information.
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.ConsoleSeparator = "\t"
	encoderConfig.EncodeLevel = colorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the two cores together.
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	logger = zap.New(core, zap.AddCaller()).WithOptions(zap.AddCallerSkip(1))

	defer logger.Sync()
}

func ParseLevel(l string) (level zapcore.Level) {
	switch l {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.DebugLevel
	}
	return
}

func colorLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	colorLevel := color.New()

	switch l {
	case zapcore.DebugLevel:
		colorLevel.Add(color.FgCyan)
	case zapcore.InfoLevel:
		colorLevel.Add(color.FgGreen)
	case zapcore.WarnLevel:
		colorLevel.Add(color.FgYellow)
	case zapcore.ErrorLevel, zapcore.DPanicLevel:
		colorLevel.Add(color.FgHiRed)
	case zapcore.PanicLevel, zapcore.FatalLevel:
		colorLevel.Add(color.FgRed)
	default:
		colorLevel.Add(color.Reset)
	}

	enc.AppendString(colorLevel.Sprint(l.CapitalString()))
}

func Debug(v ...interface{}) {
	logger.Sugar().Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	logger.Sugar().Debugf(format, v...)
}

func Info(v ...interface{}) {
	logger.Sugar().Info(v...)
}

func Infof(format string, v ...interface{}) {
	logger.Sugar().Infof(format, v...)
}

func Warn(v ...interface{}) {
	logger.Sugar().Warn(v...)
}

func Warnf(format string, v ...interface{}) {
	logger.Sugar().Warnf(format, v...)
}

func Error(v ...interface{}) {
	logger.Sugar().Error(v...)
}

func Errorf(format string, v ...interface{}) {
	logger.Sugar().Errorf(format, v...)
}

func Fatal(v ...interface{}) {
	logger.Sugar().Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	logger.Sugar().Fatalf(format, v...)
}

func Panic(v ...interface{}) {
	logger.Sugar().Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	logger.Sugar().Panicf(format, v...)
}
