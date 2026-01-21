package logger

import (
	"context"
	"os"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.SugaredLogger

func InitLogger(){
	
	//配置Encoder
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 日志等级大写
		EncodeTime:     zapcore.ISO8601TimeEncoder,  // 时间格式
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)

	//考虑到会使用日志监控体系进行监控,所以不需要过多分级
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zap.InfoLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zap.ErrorLevel
	})

	//多核心，不同地方输出日志
	cores := [...]zapcore.Core{
		zapcore.NewCore(encoder, os.Stdout, errorLevel),    //写入终端	
		zapcore.NewCore(
			encoder,
			getWriteSyncer("./logs/info.log"),
			infoLevel,
		),									//写入文件
	}


	//合并核心
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

func Sync() {
	if logger != nil {
		_ = logger.Sync()
	}
}


//创建日志轮转器，写入文件/进行轮转
func getWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,
		MaxSize:    100,  // MB
		MaxBackups: 30,   // 最大备份数
		MaxAge:     30,   // 保存天数
		Compress:   true, // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

func WithContext(ctx context.Context) *zap.SugaredLogger {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		return logger.With("trace_id", span.SpanContext().TraceID().String())
	}
	return logger
}




//封装日志方法
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}


func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

func CtxDebug(ctx context.Context, args ...interface{}) {
	WithContext(ctx).Debug(args...)
}

func CtxDebugf(ctx context.Context, template string, args ...interface{}) {
	WithContext(ctx).Debugf(template, args...)
}

func CtxInfo(ctx context.Context, args ...interface{}) {
	WithContext(ctx).Info(args...)
}

func CtxInfof(ctx context.Context, template string, args ...interface{}) {
	WithContext(ctx).Infof(template, args...)
}

func CtxWarn(ctx context.Context, args ...interface{}) {
	WithContext(ctx).Warn(args...)
}

func CtxWarnf(ctx context.Context, template string, args ...interface{}) {
	WithContext(ctx).Warnf(template, args...)
}

func CtxError(ctx context.Context, args ...interface{}) {
	WithContext(ctx).Error(args...)
}

func CtxErrorf(ctx context.Context, template string, args ...interface{}) {
	WithContext(ctx).Errorf(template, args...)
}

func CtxFatal(ctx context.Context, args ...interface{}) {
	WithContext(ctx).Fatal(args...)
}

func CtxFatalf(ctx context.Context, template string, args ...interface{}) {
	WithContext(ctx).Fatalf(template, args...)
}



