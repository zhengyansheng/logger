package logger

import (
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

type ZapLogger struct {
	LogFile    string
	Level      string
	MaxSize    int  // 每个日志文件保存的最大尺寸 单位：M
	MaxAge     int  // 文件最多保存多少天
	MaxBackups int  // 日志文件最多保存多少个备份
	Compress   bool // 是否压缩
}

func NewZapLogger(logFile, logLevel string, maxSize, maxBackups, maxAge int) *ZapLogger {
	return &ZapLogger{
		LogFile:    logFile,
		Level:      logLevel,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   false,
	}
}

func (z *ZapLogger) InitLog() error {
	writeSyncer := z.getLogWriter()
	encoder := z.getEncoder()
	// set log level
	var logLevel zapcore.Level
	switch z.Level {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	case "fatal":
		logLevel = zapcore.FatalLevel
	default:
		logLevel = zapcore.DebugLevel
	}
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger = logger.Sugar()
	return nil
}

func (z *ZapLogger) getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func (z *ZapLogger) getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   z.LogFile,
		MaxSize:    z.MaxSize,
		MaxBackups: z.MaxBackups,
		MaxAge:     z.MaxAge,
		Compress:   z.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// Debug ...
func Debug(args ...interface{}) {
	sugarLogger.Debug(args)
}

// Info ...
func Info(args ...interface{}) {
	sugarLogger.Info(args)
}

// Warn ...
func Warn(args ...interface{}) {
	sugarLogger.Warn(args)
}

// Error ...
func Error(args ...interface{}) {
	sugarLogger.Error(args)
}

// Debugf ...
func Debugf(msg string, args ...interface{}) {
	sugarLogger.Debugf(msg, args...)
}

// Infof ...
func Infof(msg string, args ...interface{}) {
	sugarLogger.Infof(msg, args...)
}

// Warnf ...
func Warnf(msg string, args ...interface{}) {
	sugarLogger.Warnf(msg, args...)
}

// Errorf ...
func Errorf(msg string, args ...interface{}) {
	sugarLogger.Errorf(msg, args...)
}
