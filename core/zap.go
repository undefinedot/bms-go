package core

import (
	"bms-go/global"
	"bms-go/utils"
	"log"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
)

type _zap struct{}

// Zap zap日志 - zap.Logger
func Zap() (logger *zap.Logger) {
	var z = new(_zap)
	// 判断配置文件指定的日志目录是否存在
	if ok, _ := utils.DirExists(global.SYS_CONFIG.Zap.Dir); !ok {
		// 不存在时自动创建
		_ = os.Mkdir(global.SYS_CONFIG.Zap.Dir, os.ModePerm)
	}

	logger = zap.New(z.GetZapCores())
	log.Println("----- zap logger init succeed -----")
	return
}

// GetEncoder 根据配置设置日志输出格式，默认console
func (z *_zap) GetEncoder() zapcore.Encoder {
	if global.SYS_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

// GetEncoderConfig 获取encoder的配置项
func (z *_zap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "log-level",
		TimeKey:     "time",
		NameKey:     "logger",
		CallerKey:   "caller",
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime:  z.CustomTime, // 函数类型
	}
}

// CustomTime 自定义zap日志的时间格式
func (z *_zap) CustomTime(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format("2006-01-02 15:04:05"))
}

// GetEncoderCore 切割json日志
func (z *_zap) GetEncoderCore(level zapcore.Level) zapcore.Core {
	// todo: 日志切割
	name := global.SYS_CONFIG.Zap.Dir + "/" + global.SYS_CONFIG.Zap.Name
	writer := &lumberjack.Logger{
		Filename: name,
		MaxSize:  10,
	}
	return zapcore.NewCore(z.GetEncoder(), zapcore.Lock(zapcore.AddSync(writer)), level)
}

// GetZapCores 日志同时在多个位置打印
func (z *_zap) GetZapCores() zapcore.Core {
	var l zapcore.Level
	l.UnmarshalText([]byte(global.SYS_CONFIG.Zap.Level))

	if global.SYS_CONFIG.Zap.Release == "dev" {
		// 开发环境下，同时使用json和console格式的日志
		core1 := z.GetEncoderCore(l)
		core2 := zapcore.NewCore(z.GetEncoder(), zapcore.Lock(os.Stdout), l)
		return zapcore.NewTee(core1, core2)
	}
	return z.GetEncoderCore(l)
}

// GetLoggerLevel 根据配置文件设置日志级别
//func (z *_zap) GetLoggerLevel(level string) zap.LevelEnablerFunc {
//	switch level {
//	case zapcore.DebugLevel.String():
//		return func(level zapcore.Level) bool { // 调试级别
//			return level == zap.DebugLevel
//		}
//	case zapcore.InfoLevel.String():
//		return func(level zapcore.Level) bool { // 日志级别
//			return level == zap.InfoLevel
//		}
//	case zapcore.WarnLevel.String():
//		return func(level zapcore.Level) bool { // 警告级别
//			return level == zap.WarnLevel
//		}
//	case zapcore.ErrorLevel.String():
//		return func(level zapcore.Level) bool { // 错误级别
//			return level == zap.ErrorLevel
//		}
//	case zapcore.DPanicLevel.String():
//		return func(level zapcore.Level) bool { // dpanic级别
//			return level == zap.DPanicLevel
//		}
//	case zapcore.PanicLevel.String():
//		return func(level zapcore.Level) bool { // panic级别
//			return level == zap.PanicLevel
//		}
//	case zapcore.FatalLevel.String():
//		return func(level zapcore.Level) bool { // 终止级别
//			return level == zap.FatalLevel
//		}
//	default:
//		return func(level zapcore.Level) bool { // 调试级别
//			return level == zap.DebugLevel
//		}
//	}
//}
