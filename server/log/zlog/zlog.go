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
const loggerTraceIdKey = "traceId"

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
		if len(fields) == 1 && fields[0].Key == loggerTraceIdKey {
			ctx = context.WithValue(ctx, loggerTraceIdKey, fields[0].String)
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

func addTraceId(ctx context.Context, format string, v []interface{}) (string, []interface{}) {
	if formatJson() {
		return format, v
	} else {
		if traceId, ok := ctx.Value(loggerTraceIdKey).(string); ok {
			_v := make([]interface{}, 0)
			_v = append(_v, traceId)
			_v = append(_v, v...)
			return "%s \t" + format, _v
		}
		return format, v
	}
}

// 下面的logger方法会携带trace id

func CtxInfof(ctx context.Context, format string, v ...interface{}) {
	format, v = addTraceId(ctx, format, v)
	withContext(ctx).Info(fmt.Sprintf(format, v...))
}

func CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	format, v = addTraceId(ctx, format, v)
	withContext(ctx).Error(fmt.Sprintf(format, v...))
}

func CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	format, v = addTraceId(ctx, format, v)
	withContext(ctx).Warn(fmt.Sprintf(format, v...))
}

func CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	format, v = addTraceId(ctx, format, v)
	withContext(ctx).Debug(fmt.Sprintf(format, v...))
}

func CtxPanicf(ctx context.Context, format string, v ...interface{}) {
	format, v = addTraceId(ctx, format, v)
	withContext(ctx).Panic(fmt.Sprintf(format, v...))
}

func CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	format, v = addTraceId(ctx, format, v)
	withContext(ctx).Fatal(fmt.Sprintf(format, v...))
}

//func TraceInfof(ctx context.Context, format string, v ...interface{}) {
//	logger.Info(ctx.Value(loggerTraceIdKey).(string), fmt.Sprintf(format, v...))
//}
