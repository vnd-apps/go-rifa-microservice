package logger

import (
	"context"
	"io"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormlogger "gorm.io/gorm/logger"
)

type Logger struct {
	*zap.SugaredLogger
}

type GormLogger struct {
	*Logger
	gormlogger.Config
}

type FxLogger struct {
	*Logger
}

type GinLogger struct {
	*Logger
}

func NewLogger(logLevel, env string) *Logger {
	config := zap.NewProductionConfig()

	if env == "dev" {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var level zapcore.Level

	switch logLevel {
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
		level = zap.PanicLevel
	}

	config.Level.SetLevel(level)

	zapLogger, err := config.Build()
	if err != nil {
		panic("Failed to create logger")
	}

	globalLog := zapLogger.Sugar()

	return &Logger{
		SugaredLogger: globalLog,
	}
}

func newSugaredLogger(logger *zap.Logger) *Logger {
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

func (l *Logger) GetGormLogger() gormlogger.Interface {
	skipFrameCount := 3

	logger := l.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(skipFrameCount),
	).Desugar()

	return &GormLogger{
		Logger: newSugaredLogger(logger),
		Config: gormlogger.Config{
			LogLevel: gormlogger.Info,
		},
	}
}

func (l *Logger) GetGinLogger() io.Writer {
	logger := l.WithOptions(
		zap.WithCaller(false),
	).Desugar()

	return GinLogger{
		Logger: newSugaredLogger(logger),
	}
}

// ------ GORM logger interface implementation -----

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level

	return &newlogger
}

func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Info {
		l.Debugf(str, args...)
	}
}

func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Warn {
		l.Warnf(str, args...)
	}
}

func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Error {
		l.Errorf(str, args...)
	}
}

func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}

	elapsed := time.Since(begin)

	if l.LogLevel >= gormlogger.Info {
		sql, rows := fc()
		l.Debug("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)

		return
	}

	if l.LogLevel >= gormlogger.Warn {
		sql, rows := fc()
		l.SugaredLogger.Warn("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)

		return
	}

	if l.LogLevel >= gormlogger.Error {
		sql, rows := fc()
		l.SugaredLogger.Error("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)

		return
	}
}

func (l FxLogger) Printf(str string, args ...interface{}) {
	if len(args) > 0 {
		l.Debugf(str, args)
	}

	l.Debug(str)
}

func (l GinLogger) Write(p []byte) (n int, err error) {
	l.Info(string(p))

	return len(p), nil
}
