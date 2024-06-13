package log

import (
	"github.com/Dbinggo/HireSphere/server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

var Zap = new(_zap)

type _zap struct{}

//func GetZap() (logger *zap.Logger) {
//	if ok, _ := utils.PathExists(global.Config.Log.Director); !ok { // 判断是否有Director文件夹
//		fmt.Printf("create %v directory\n", global.Config.Zap.Director)
//		_ = os.Mkdir(global.Config.Zap.Director, os.ModePerm)
//	}
//
//	cores := internal.Zap.GetZapCores()
//	logger = zap.New(zapcore.NewTee(cores...))
//
//	if global.GVA_CONFIG.Zap.ShowLine {
//		logger = logger.WithOptions(zap.AddCaller())
//	}
//	return logger
//}

// GetEncoder 获取 zapcore.Encoder
func (z *_zap) GetEncoder() zapcore.Encoder {
	if global.Config.Log.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

// GetEncoderConfig 获取zapcore.EncoderConfig
func (z *_zap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, //大写编码器带颜色
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

// GetEncoderCore 获取Encoder的 zapcore.Core
func (z *_zap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {

	log := &lumberjack.Logger{
		Filename:   "",   // 日志文件的位置
		MaxSize:    10,   // 文件最大尺寸（以MB为单位）
		MaxBackups: 3,    // 保留的最大旧文件数量
		MaxAge:     28,   // 保留旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
		LocalTime:  true, // 使用本地时间创建时间戳
	}
	writer := zapcore.AddSync(log)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		writer,
		zap.InfoLevel,
	)
	logger := zap.New(core)
	defer logger.Sync() // 刷新任何缓冲的日志条目
	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}

// GetZapCores 根据配置文件的Level获取 []zapcore.Core
func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := zapcore.Level(global.Config.Log.Level); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}

// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}
