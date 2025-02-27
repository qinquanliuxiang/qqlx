package logger

import (
	"context"
	"os"
	"qqlx/base/conf"
	"qqlx/base/constant"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Caller() *zap.SugaredLogger {
	return zap.S().WithOptions(zap.AddCaller())
}

// customTimeEncoder 用于在日志中打印指定时区的时间
// func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
// 	// cst, err := time.LoadLocation("Asia/Shanghai")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// enc.AppendString(t.In(cst).Format("2006-01-02 15:04:05"))
// 	enc.AppendString(t.Format("2006-01-02 15:04:05"))
// }

func InitLogger() {
	config := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeCaller:  zapcore.FullCallerEncoder,
	}
	var logFormat = conf.GetLogFormat()
	var encoder zapcore.Encoder
	encoder = zapcore.NewJSONEncoder(config)

	writer := zapcore.AddSync(os.Stdout)
	var logLevelStr = conf.GetLogLevel()
	var logLevel zapcore.Level
	switch logLevelStr {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "err":
		logLevel = zap.ErrorLevel
	default:
		logLevel = zap.InfoLevel
	}
	core := zapcore.NewCore(encoder, writer, logLevel)

	logger := zap.New(core)
	zap.ReplaceGlobals(logger)
	zap.S().Infof("log initialization successful, log format: %s, log level: %s", logFormat, logLevelStr)
}

func WithContext(ctx context.Context, addCaller bool) *zap.SugaredLogger {
	if addCaller {
		lg := zap.S().WithOptions(zap.AddCaller())
		if traceID := ctx.Value(constant.TraceID).(string); traceID != "" {
			return lg.With(constant.TraceID, traceID)
		}
		return lg
	}
	return zap.S()
}
