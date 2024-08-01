package logger

import (
	"fmt"
	"github.com/eyslce/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type factory struct {
	logger *zap.SugaredLogger
}

func (f factory) Println(msg ...interface{}) {
	f.logger.Infoln(msg...)
}

var loggerFactory *factory

func InitLogger(logfile string) {
	encoder := getEncoder()
	ws := getLogWriter(logfile)
	core := zapcore.NewCore(encoder, ws, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())

	loggerFactory = &factory{logger: logger.Sugar()}

}

func getEncoder() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format(time.RFC3339))
	}
	return zapcore.NewJSONEncoder(cfg)
}

func getLogWriter(logfile string) zapcore.WriteSyncer {
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(&lumberjack.Logger{
		Filename:     logfile,
		MaxSize:      300,
		MaxAge:       30,
		MaxBackups:   10,
		DailyRolling: true,
		LocalTime:    true,
	}), zapcore.Lock(os.Stdout))

}

func Info(args ...interface{}) {
	loggerFactory.logger.Infoln(args...)
}

func Infof(msg string, args ...interface{}) {
	loggerFactory.logger.Infoln(fmt.Sprintf(msg, args...))
}

func Warn(args ...interface{}) {
	loggerFactory.logger.Warnln(args...)
}

func Warnf(msg string, args ...interface{}) {
	loggerFactory.logger.Warnln(fmt.Sprintf(msg, args...))
}

func Error(args ...interface{}) {
	loggerFactory.logger.Errorln(args...)
}

func Errorf(msg string, args ...interface{}) {
	loggerFactory.logger.Errorln(fmt.Sprintf(msg, args...))
}

func GetLoggerFactory() factory {
	return *loggerFactory
}

func Sync() {
	err := loggerFactory.logger.Sync()
	if err != nil {
		loggerFactory.logger.Errorln(fmt.Sprintf("log sync error:%s", err))
		return
	}
}
