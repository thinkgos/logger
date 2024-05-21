package logger

import (
	"io"

	"go.uber.org/zap/zapcore"
)

type LumberjackFile struct {
	// Path 日志保存路径, 默认 empty, 即当前路径
	Path string `yaml:"path" json:"path"`
	// see https://github.com/natefinch/lumberjack
	// lumberjack.Log
	// Filename 空字符使用默认, 默认<processname>-lumberjack.log
	Filename string `yaml:"filename" json:"filename"`
	// MaxSize 每个日志文件最大尺寸(MB), 默认100MB
	MaxSize int `yaml:"maxSize" json:"maxSize"`
	// MaxAge 日志文件保存天数, 默认0 不删除
	MaxAge int `yaml:"maxAge" json:"maxAge"`
	// MaxBackups 日志文件保存备份数, 默认0 都保存
	MaxBackups int `yaml:"maxBackups" json:"maxBackups"`
	// LocalTime 是否格式化时间戳, 默认UTC时间
	LocalTime bool `yaml:"localTime" json:"localTime"`
	// Compress 是否使用gzip压缩文件, 采用默认不压缩
	Compress bool `yaml:"compress" json:"compress"`
}

// Config 日志配置
type Config struct {
	// Level 日志等级, debug,info,warn,error,dpanic,panic,fatal, 默认warn
	Level string `yaml:"level" json:"level"`
	// Format: 编码格式: json,console 默认json
	Format string `yaml:"format" json:"format"`
	// 编码器类型, 默认: LowercaseLevelEncoder
	// LowercaseLevelEncoder: 小写编码器
	// LowercaseColorLevelEncoder: 小写编码器带颜色
	// CapitalLevelEncoder: 大写编码器
	// CapitalColorLevelEncoder: 大写编码器带颜色
	EncodeLevel string `yaml:"encodeLevel" json:"encodeLevel"`
	// Adapter 输出适配器, file,console,multi,custom,file-custom,console-custom,multi-custom 默认 console
	Adapter string `yaml:"adapter" json:"adapter"`
	// Stack 是否使能栈调试输出, 默认false
	Stack bool `yaml:"stack" json:"stack"`
	// Writer 输出
	// 当adapter有附带custom时, 如果为writer为空, 将使用os.Stdout
	Writer []io.Writer `yaml:"-" json:"-"`
	// EncoderConfig 如果配置该项,则 EncodeLevel 将被覆盖
	EncoderConfig *zapcore.EncoderConfig `yaml:"-" json:"-"`
	// 文件配置, 仅Adapter有file时有效
	File LumberjackFile `yaml:"file" json:"file"`
}

// Option An Option configures a Log.
type Option func(c *Config)

// WithConfig with config
func WithConfig(cfg Config) Option {
	return func(c *Config) { *c = cfg }
}

// WithLevel with level
// debug(default),info,warn,error,dpanic,panic,fatal
func WithLevel(level string) Option {
	return func(c *Config) { c.Level = level }
}

// WithFormat with format
// json(default) or console
func WithFormat(format string) Option {
	return func(c *Config) { c.Format = format }
}

// WithEncodeLevel with EncodeLevel
// LowercaseLevelEncoder(default): 小写编码器
// LowercaseColorLevelEncoder: 小写编码器带颜色
// CapitalLevelEncoder: 大写编码器
// CapitalColorLevelEncoder: 大写编码器带颜色
func WithEncodeLevel(encodeLevel string) Option {
	return func(c *Config) { c.EncodeLevel = encodeLevel }
}

// EncoderConfig 如果配置该项,则 EncodeLevel 将被覆盖
func WithEncoderConfig(encoderConfig *zapcore.EncoderConfig) Option {
	return func(c *Config) { c.EncoderConfig = encoderConfig }
}

// WithAdapter with adapter
// file,console(default),multi,custom,file-custom,console-custom,multi-custom
// writer: 当 adapter=custom 使用,如果为writer为空,将使用os.Stdout
func WithAdapter(adapter string, writer ...io.Writer) Option {
	return func(c *Config) {
		c.Adapter = adapter
		c.Writer = writer
	}
}

// WithStack with stack
// Stack 是否使能栈调试输出, 默认false
func WithStack(stack bool) Option {
	return func(c *Config) { c.Stack = stack }
}

// WithPath with path
// 日志保存路径, 默认 empty, 即当前路径
func WithPath(path string) Option {
	return func(c *Config) { c.File.Path = path }
}

/******************************** lumberjack.Log **************************************/

// WithFilename with filename
// 空字符使用默认, 默认<processname>-lumberjack.log
func WithFilename(filename string) Option {
	return func(c *Config) { c.File.Filename = filename }
}

// WithMaxSize with max size
// 每个日志文件最大尺寸(MB), 默认100MB
func WithMaxSize(maxSize int) Option {
	return func(c *Config) { c.File.MaxSize = maxSize }
}

// WithMaxAge with max age
// 日志文件保存天数, 默认0 不删除
func WithMaxAge(maxAge int) Option {
	return func(c *Config) { c.File.MaxAge = maxAge }
}

// WithMaxBackups with max backup
// 日志文件保存备份数, 默认0 都保存
func WithMaxBackups(maxBackups int) Option {
	return func(c *Config) { c.File.MaxBackups = maxBackups }
}

// WithEnableLocalTime with local time
// 是否格式化时间戳, 默认UTC时间
func WithEnableLocalTime() Option {
	return func(c *Config) { c.File.LocalTime = true }
}

// WithEnableCompress with compress
// 是否使用gzip压缩文件, 采用默认不压缩
func WithEnableCompress() Option {
	return func(c *Config) { c.File.Compress = true }
}
