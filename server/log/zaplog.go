package log

import (
	"github.com/Dbinggo/HireSphere/server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

/*
[感谢伟人 让我彻底搞懂 zap]https://juejin.cn/post/7313979344561242162?searchId=20240613163846377BACC6CC0FB80CC369
*/

func GetZap() *zap.Logger {
	var logger *zap.Logger
	var z zapConfig

	switch global.Config.App.Env {
	case "pro":
		logger, _ = zap.NewProduction()
	case "dev":
	default:
		logger = zap.New(*z.Core, z.Options...)
	}
	defer logger.Sync()
	return logger
}

type zapConfig struct {
	Core         *zapcore.Core
	Encoder      *zapcore.Encoder
	WriteSyncer  *zapcore.WriteSyncer
	LevelEnabler *zapcore.LevelEnabler
	Options      []zap.Option
}

// 定制core
func (z *zapConfig) setCore() *zapConfig {
	*z.Core = zapcore.NewCore(*z.Encoder, *z.WriteSyncer, *z.LevelEnabler)
	return z
}

// Encoder 是编码器，以什么样的格式写入日志。
// 目前，zap只支持两种编码器——JSON Encoder和Console Encoder
func (z *zapConfig) setEncoder(encoder func(cfg zapcore.EncoderConfig) zapcore.Encoder) *zapConfig {
	*z.Encoder = encoder(zapcore.EncoderConfig{
		MessageKey:          "message",
		LevelKey:            "level",
		TimeKey:             "time",
		NameKey:             "name",
		CallerKey:           "caller",
		FunctionKey:         "function",
		StacktraceKey:       "stacktrace",
		SkipLineEnding:      false,
		EncodeName:          nil,
		NewReflectedEncoder: nil,
		ConsoleSeparator:    "\t",

		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, //大写彩色
		EncodeTime:     newTimeEncoder(),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // path/file
	})
	return z
}

func (z *zapConfig) setWriteSyncer() *zapConfig {
	return z
}
func (z *zapConfig) setLevelEnabler(enabler zapcore.LevelEnabler) *zapConfig {
	*z.LevelEnabler = enabler
	return z
}
func (z *zapConfig) setOptions() *zapConfig {
	return z
}

func newTimeEncoder() zapcore.TimeEncoder {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006/1/2 15:04:05.000"))
	}
}
