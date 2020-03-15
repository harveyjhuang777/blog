package logger

import (
	"os"

	"github.com/jwjhuang/blog/service/utils/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	appLog *zap.SugaredLogger
)

func Log() *zap.SugaredLogger {
	return appLog
}

func Start() {
	appLog = initLogger(getZapLevel(conf.GetLogLevel()))
}

func initLogger(level zapcore.Level) *zap.SugaredLogger {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// package/file:line
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func getZapLevel(l string) zapcore.Level {
	switch l {
	case zapcore.DebugLevel.String(): // "debug"
		return zapcore.DebugLevel
	case zapcore.InfoLevel.String(): // "info"
		return zapcore.InfoLevel
	case zapcore.WarnLevel.String(): // "warn"
		return zapcore.WarnLevel
	case zapcore.ErrorLevel.String(): // "error"
		return zapcore.ErrorLevel
	case zapcore.DPanicLevel.String(): // "dpanic"
		return zapcore.DPanicLevel
	case zapcore.PanicLevel.String(): // "panic"
		return zapcore.PanicLevel
	case zapcore.FatalLevel.String(): // "fatal"
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
