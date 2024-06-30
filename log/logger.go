package log

import (
	"context"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

// 为保持常用printf风格，采用zap的sugarLogger
var logger *zap.SugaredLogger

// OutputConfig 日志配置结构体
type OutputConfig struct {
	// 日志输出端 （file，console）
	Writer      string      `yaml:"writer"`
	WriteConfig WriteConfig `yaml:"write_config"`
	// 日志输出格式 （console， json）
	Formatter    string       `yaml:"formatter"`
	FormatConfig FormatConfig `yaml:"format_config"`
	Level        string       `yaml:"level"`
}

// FormatConfig Formater配置
type FormatConfig struct {
	// TimeFmt 日志输出时间格式，空默认为"2006-01-02 15:04:05.000"
	TimeFmt string `yaml:"time_fmt"`
	// TimeKey 日志输出时间key， 默认为"Timestamp"
	TimeKey string `yaml:"time_key"`
	// LevelKey 日志级别输出key， 默认为"Level"
	LevelKey string `yaml:"level_key"`
	// NameKey 日志名称key， 默认为"Name"
	NameKey string `yaml:"name_key"`
	// CallerKey 日志输出调用者key， 默认"Caller"
	CallerKey string `yaml:"caller_key"`
	// FunctionKey 日志输出调用者函数名， 默认""，表示不打印函数名
	FunctionKey string `yaml:"function_key"`
	// MessageKey 日志输出消息体key，默认"Message"
	MessageKey string `yaml:"message_key"`
	// StacktraceKey 日志输出堆栈trace key， 默认"Stacktrace"
	StacktraceKey string `yaml:"stacktrace_key"`
}

// WriteConfig Writer配置
type WriteConfig struct {
	// 文件日志路径（含文件名）
	LogPath string `yaml:"log_path"`
	// 日志最大大小 单位（MB）
	MaxSize int `yaml:"max_size"`
	// 日志最大保留时间，单位（天）
	MaxAge int `yaml:"max_age"`
	// 日志最大保留文件数
	MaxBackups int `yaml:"max_backups"`
	// 是否启用压缩
	Compress bool `yaml:"compress"`
}

// Levels zap日志级别映射表
var Levels = map[string]zapcore.Level{
	"":      zapcore.DebugLevel,
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

// GetLogger 获取logger
func GetLogger() *zap.SugaredLogger {
	return logger
}

// GetLoggerWithContext 获取logger 携带链路信息
func GetLoggerWithContext(ctx context.Context) *zap.SugaredLogger {
	if ctx != nil {
		return logger.With("TraceId", ctx.Value("TraceId"), "UserId", ctx.Value("UserId"))
	}
	return logger
}

// InitLogger 初始化日志组件
func InitLogger(opc []OutputConfig) {

	var cores []zapcore.Core
	for _, o := range opc {
		writer := newWriter(&o)
		encoder := newEncoder(&o)
		level := Levels[strings.ToLower(o.Level)]
		core := zapcore.NewCore(encoder, writer, level)
		cores = append(cores, core)
	}

	logger = zap.New(
		zapcore.NewTee(cores...),
		zap.AddCallerSkip(1),
		zap.AddCaller(),
	).Sugar()

}

// Fatal 日志
func Fatal(ctx context.Context, args ...interface{}) {
	GetLoggerWithContext(ctx).Fatal(args...)
}

// Fatalf printf风格的日志
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	GetLoggerWithContext(ctx).Fatalf(format, args...)
}

// Error 日志
func Error(ctx context.Context, args ...interface{}) {
	GetLoggerWithContext(ctx).Error(args...)
}

// Errorf printf风格的日志
func Errorf(ctx context.Context, format string, args ...interface{}) {
	GetLoggerWithContext(ctx).Errorf(format, args...)
}

// Warn 日志
func Warn(ctx context.Context, args ...interface{}) {
	GetLoggerWithContext(ctx).Warn(args...)
}

// Warnf printf风格的日志
func Warnf(ctx context.Context, format string, args ...interface{}) {
	GetLoggerWithContext(ctx).Warnf(format, args...)
}

// Info 日志
func Info(ctx context.Context, args ...interface{}) {
	GetLoggerWithContext(ctx).Info(args...)
}

// Infof printf风格的日志
func Infof(ctx context.Context, format string, args ...interface{}) {
	GetLoggerWithContext(ctx).Infof(format, args...)
}

// Debug 日志
func Debug(ctx context.Context, args ...interface{}) {
	GetLoggerWithContext(ctx).Debug(args...)
}

// Debugf printf风格的日志
func Debugf(ctx context.Context, format string, args ...interface{}) {
	GetLoggerWithContext(ctx).Debugf(format, args...)
}

// WithField 设置数据字典
func WithField(ctx context.Context, fields ...interface{}) *zap.SugaredLogger {
	return GetLoggerWithContext(ctx).WithOptions(zap.AddCallerSkip(-1)).With(fields...)
}

// GetLogKey 获取用户自定义log输出字段名，没有则使用默认的
func getLogKey(defKey, key string) string {
	if key == "" {
		return defKey
	}
	return key
}

// DefaultTimeFormat 默认时间格式
func DefaultTimeFormat(t time.Time) []byte {
	t = t.Local()
	year, month, day := t.Date()
	hour, minute, second := t.Clock()
	micros := t.Nanosecond() / 1000

	buf := make([]byte, 23)
	buf[0] = byte((year/1000)%10) + '0'
	buf[1] = byte((year/100)%10) + '0'
	buf[2] = byte((year/10)%10) + '0'
	buf[3] = byte(year%10) + '0'
	buf[4] = '-'
	buf[5] = byte((month)/10) + '0'
	buf[6] = byte((month)%10) + '0'
	buf[7] = '-'
	buf[8] = byte((day)/10) + '0'
	buf[9] = byte((day)%10) + '0'
	buf[10] = ' '
	buf[11] = byte((hour)/10) + '0'
	buf[12] = byte((hour)%10) + '0'
	buf[13] = ':'
	buf[14] = byte((minute)/10) + '0'
	buf[15] = byte((minute)%10) + '0'
	buf[16] = ':'
	buf[17] = byte((second)/10) + '0'
	buf[18] = byte((second)%10) + '0'
	buf[19] = '.'
	buf[20] = byte((micros/100000)%10) + '0'
	buf[21] = byte((micros/10000)%10) + '0'
	buf[22] = byte((micros/1000)%10) + '0'
	return buf
}

// CustomTimeFormat 自定义时间格式
func CustomTimeFormat(t time.Time, format string) string {
	return t.Format(format)
}

// NewTimeEncoder 创建时间格式encoder
func NewTimeEncoder(format string) zapcore.TimeEncoder {
	switch format {
	case "":
		return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendByteString(DefaultTimeFormat(t))
		}
	case "seconds":
		return zapcore.EpochTimeEncoder
	case "milliseconds":
		return zapcore.EpochMillisTimeEncoder
	case "nanoseconds":
		return zapcore.EpochNanosTimeEncoder
	default:
		return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(CustomTimeFormat(t, format))
		}
	}
}

func newEncoder(c *OutputConfig) zapcore.Encoder {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        getLogKey("Timestamp", c.FormatConfig.TimeKey),
		LevelKey:       getLogKey("Level", c.FormatConfig.LevelKey),
		NameKey:        getLogKey("Name", c.FormatConfig.NameKey),
		CallerKey:      getLogKey("Caller", c.FormatConfig.CallerKey),
		FunctionKey:    getLogKey(zapcore.OmitKey, c.FormatConfig.FunctionKey),
		MessageKey:     getLogKey("Message", c.FormatConfig.MessageKey),
		StacktraceKey:  getLogKey("Stacktrace", c.FormatConfig.StacktraceKey),
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     NewTimeEncoder(c.FormatConfig.TimeFmt),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	switch c.Formatter {
	case "console":
		return zapcore.NewConsoleEncoder(encoderCfg)
	case "json":
		return zapcore.NewJSONEncoder(encoderCfg)
	default:
		return zapcore.NewConsoleEncoder(encoderCfg)
	}
}

func newWriter(c *OutputConfig) zapcore.WriteSyncer {
	switch c.Writer {
	case "file":
		luberJackLogger := &lumberjack.Logger{
			Filename:   getLogKey("./run.log", c.WriteConfig.LogPath),
			MaxSize:    c.WriteConfig.MaxSize,
			MaxAge:     c.WriteConfig.MaxAge,
			MaxBackups: c.WriteConfig.MaxBackups,
			LocalTime:  true,
			Compress:   c.WriteConfig.Compress,
		}
		return zapcore.AddSync(luberJackLogger)
	case "console":
		return zapcore.Lock(os.Stdout)
	default:
		return zapcore.Lock(os.Stdout)
	}
}
