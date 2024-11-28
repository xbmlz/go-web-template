package logger

import (
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

type Config struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
	Compress   bool   `json:"compress"`
}

func Init(c *Config) {
	// cores
	cores := []zapcore.Core{}

	// level
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(ParseLevel(c.Level))

	// console log
	// set encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder
	encoderConfig.ConsoleSeparator = "\t"

	cores = append(cores, zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(colorable.NewColorableStdout()),
		atomicLevel,
	))

	// file log
	if c.Filename != "" {
		fileEncoderConfig := encoderConfig
		fileEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		fileEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		fileEncoder := zapcore.NewJSONEncoder(fileEncoderConfig)
		fileCore := zapcore.NewCore(
			fileEncoder,
			zapcore.AddSync(&lumberjack.Logger{
				Filename:   c.Filename,
				MaxSize:    c.MaxSize,
				MaxBackups: c.MaxBackups,
				MaxAge:     c.MaxAge,
				Compress:   c.Compress,
			}),
			atomicLevel,
		)
		cores = append(cores, fileCore)
	}

	zapOpts := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	}
	if c.Level == "debug" {
		zapOpts = append(zapOpts, zap.Development(), zap.AddStacktrace(zapcore.ErrorLevel))
	}

	logger = zap.New(zapcore.NewTee(cores...), zapOpts...)

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
