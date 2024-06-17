package zlog

import (
	"context"
	"fmt"
	"github.com/Dbinggo/HireSphere/server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logKey string

const loggerKey = "logger"
const loggerExKey = "loggerEx"

// 分隔符
const formatSeparator = "%v\t"

var myLogger struct {
	*zap.Logger
	TraceId string
}
var logger *zap.Logger

func formatJson() bool {
	return global.Config.App.Env == "json"
}

// NewContext
//
//	@Description:给指定context添加字段 实现类似traceid作用
//	@param ctx
//	@param fields
//	@return context.Context
func NewContext(ctx context.Context, fields ...zapcore.Field) context.Context {
	if formatJson() {
		return context.WithValue(ctx, loggerKey, withContext(ctx).With(fields...))
	} else {
		ctx = context.WithValue(ctx, loggerKey, withContext(ctx))
		if ex, ok := ctx.Value(loggerExKey).([]zapcore.Field); ok {
			ex = append(ex, fields...)
			context.WithValue(ctx, loggerExKey, ex)
		} else {
			ctx = context.WithValue(ctx, loggerExKey, fields)
		}
		return ctx
	}
}

func InitLogger(zapLogger *zap.Logger) {
	logger = zapLogger
}

// 从指定的context返回一个zap实例
func withContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return logger
	}
	if ctxLogger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return ctxLogger
	}
	return logger
}

func Infof(format string, v ...interface{}) {
	logger.Info(fmt.Sprintf(format, v...))
}

func Errorf(format string, v ...interface{}) {
	logger.Error(fmt.Sprintf(format, v...))
}

func Warnf(format string, v ...interface{}) {
	logger.Warn(fmt.Sprintf(format, v...))
}

func Debugf(format string, v ...interface{}) {
	logger.Debug(fmt.Sprintf(format, v...))
}

func Panicf(format string, v ...interface{}) {
	logger.Fatal(fmt.Sprintf(format, v...))
}

func Fatalf(format string, v ...interface{}) {
	logger.Fatal(fmt.Sprintf(format, v...))
}

func addExField(ctx context.Context, format string, v []interface{}) (string, []interface{}) {
	if formatJson() {
		return format, v
	} else {
		if exField, ok := ctx.Value(loggerExKey).([]zap.Field); ok {
			_v := make([]interface{}, 0)
			for _, field := range exField {
				_v = append(_v, field.String)
				format = formatSeparator + format
			}
			_v = append(_v, v...)
			v = _v
		}
		return format, v
	}
}

// 下面的logger方法会携带trace id

func CtxInfof(ctx context.Context, format string, v ...interface{}) {
	format, v = addExField(ctx, format, v)
	withContext(ctx).Info(fmt.Sprintf(format, v...))
}

func CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	format, v = addExField(ctx, format, v)
	withContext(ctx).Error(fmt.Sprintf(format, v...))
}

func CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	format, v = addExField(ctx, format, v)
	withContext(ctx).Warn(fmt.Sprintf(format, v...))
}

func CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	format, v = addExField(ctx, format, v)
	withContext(ctx).Debug(fmt.Sprintf(format, v...))
}

func CtxPanicf(ctx context.Context, format string, v ...interface{}) {
	format, v = addExField(ctx, format, v)
	withContext(ctx).Panic(fmt.Sprintf(format, v...))
}

func CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	format, v = addExField(ctx, format, v)
	withContext(ctx).Fatal(fmt.Sprintf(format, v...))
}

//func TraceInfof(ctx context.Context, format string, v ...interface{}) {
//	logger.Info(ctx.Value(loggerTraceIdKey).(string), fmt.Sprintf(format, v...))
//}
