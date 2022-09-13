package log

import (
	"fmt"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugaredLog *zap.SugaredLogger
var log *zap.Logger

type logger zap.SugaredLogger

func (s *logger) Printf(msg string, args ...interface{}) {
	sugaredLog.Infof(msg, args...)
}

func GetLog() *logger {
	return &logger{}
}

func InitLogger(args ...LogOption) error {
	params := ParseLogParameter(args...)
	logDirectory := "./logs"
	if params.logDirectory != "" {
		logDirectory = params.logDirectory
	}
	if _, err := os.Stat(logDirectory); os.IsNotExist(err) {
		fmt.Printf("create %v directory\n", logDirectory)
		err = os.Mkdir(logDirectory, os.ModePerm)
		if err != nil {
			return err
		}
	}

	linkName := "latest_log"
	if params.linkName != "" {
		linkName = params.linkName
	}

	writeSyncer, err := getLogWriter(logDirectory, params.logFilePrefix,
		linkName, params.logInConsole, params.maxAge, params.rotationTime)
	if err != nil {
		return err
	}
	encoder := getEncoder(params)
	var level zapcore.Level
	switch params.level { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)

	log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	if params.showCaller {
		log = log.WithOptions(zap.AddCaller())
	}
	sugaredLog = log.Sugar()
	return nil
}

func getEncoder(params *LogParameter) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	switch params.encodeLevel {
	case "LowercaseLevelEncoder": // 小写编码器(默认)
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder": // 小写编码器带颜色
		encoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder": // 大写编码器
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder": // 大写编码器带颜色
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	if params.jsonFormat {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(logDirectory, logFilePrefix, linkName string, logInConsole bool, maxAge, rotationTime int) (
	zapcore.WriteSyncer, error) {

	fileWriter, err := rotatelogs.New(
		path.Join(logDirectory, logFilePrefix+"%Y-%m-%d.log"),
		rotatelogs.WithLinkName(linkName),
		rotatelogs.WithMaxAge(time.Duration(maxAge)*24*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Hour),
	)
	if logInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
			zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

func Debug(msg string, args ...zap.Field) {
	log.Debug(msg, args...)
}

func Info(msg string, args ...zap.Field) {
	log.Info(msg, args...)
}

func Warn(msg string, args ...zap.Field) {
	log.Warn(msg, args...)
}

func Error(msg string, args ...zap.Field) {
	log.Error(msg, args...)
}

func DPanic(msg string, args ...zap.Field) {
	log.DPanic(msg, args...)
}

func Panic(msg string, args ...zap.Field) {
	log.Panic(msg, args...)
}

func Fatal(msg string, args ...zap.Field) {
	log.Fatal(msg, args...)
}

func Debugw(msg string, args ...interface{}) {
	sugaredLog.Debugw(msg, args...)
}

func Infow(msg string, args ...interface{}) {
	sugaredLog.Infow(msg, args...)
}

func Warnw(msg string, args ...interface{}) {
	sugaredLog.Warnw(msg, args...)
}

func Errorw(msg string, args ...interface{}) {
	sugaredLog.Errorw(msg, args...)
}

func DPanicw(msg string, args ...interface{}) {
	sugaredLog.DPanicw(msg, args...)
}

func Panicw(msg string, args ...interface{}) {
	sugaredLog.Panicw(msg, args...)
}

func Fatalw(msg string, args ...interface{}) {
	sugaredLog.Fatalw(msg, args...)
}

func Debugf(template string, args ...interface{}) {
	sugaredLog.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	sugaredLog.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	sugaredLog.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	sugaredLog.Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	sugaredLog.DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	sugaredLog.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	sugaredLog.Fatalf(template, args...)
}
