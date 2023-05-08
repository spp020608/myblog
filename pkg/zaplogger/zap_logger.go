package zaplogger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"myblog/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// logTimeLayout 自定义时间格式
const logTimeLayout = "2006-01-02 15:04:05"

// 对应 zapcore.ISO8601TimeEncoder 这种时间编码的日期格式
const iso8601TimeLayout = "2006-01-02T15:04:05.000Z0700"


//
// New 创建并初始化 zap 日志框架
//	@return	error
//
func init() {
	initZapLogger()
}

//
// InitZapLogger 初始化 zap 日志库
//
func initZapLogger() {
	consoleEncoder := customConsoleEncoder() // 自定义编码器

	// 向控制台窗口输出的 zapcore.Core 对象
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)

	// 读取配置文件中记录的配置参数
	zapConfig := config.Get().LogConfig

	// 日志切割对象
	lumberJackLogger := &lumberjack.Logger{
		Filename:   zapConfig.Filename,   // 日志文件名称
		MaxSize:    zapConfig.MaxSize,    // 日志文件所占空间大小(单位:M)
		MaxAge:     zapConfig.MaxAge,     // 日志文件存活时长(单位:天)
		MaxBackups: zapConfig.MaxBackups, // 日志备份的时候, 最大备份数量
		LocalTime:  zapConfig.LocalTime,  // 日志备份的时候, 是否使用本地时间作为备份文件名, 如果不使用本地时间的话, 会有 8 个小时的时差
		Compress:   zapConfig.Compress,   // 日志是否压缩
	}

	// 向日志文件输出的 zapcore.Core 对象
	logFileCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(lumberJackLogger), zapcore.DebugLevel)

	// 最终可以同时向控制台和日志文件输出的 zapcore.Core 对象
	finalCore := zapcore.NewTee(consoleCore, logFileCore)

	// 创建 Logger 对象
	zapLogger := zap.New(finalCore, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	// 使用自己创建的这个 Logger 取代 zap 库的全局 Logger,
	// 在应用程序其他地方可以直接通过 zap.L() 和 zap.S() 来使用我们自定义的 zaplogger 独享
	zap.ReplaceGlobals(zapLogger)
}

//
// customConsoleEncoder 自定义 ConsoleEncoder
//	@return	zapcore.Encoder 接口
//
func customConsoleEncoder() zapcore.Encoder {
	// 自定义编码器配置
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",                         // 以 Json 格式输出的时候, 日志消息对应的 key
		LevelKey:       "level",                       // 以 Json 格式输出的时候, 日志级别对应的 key
		TimeKey:        "ts",                          // 以 Json 格式输出的时候, 日志时间对应的 key
		NameKey:        "zaplogger",                   // 以 Json 格式输出的时候, 日志名称对应的 key
		CallerKey:      "caller_line",                 // 以 Json 格式输出的时候, 调用者对应的 key
		FunctionKey:    zapcore.OmitKey,               // 以 Json 格式输出的时候,
		StacktraceKey:  "stacktrace",                  // 以 Json 格式输出的时候,
		SkipLineEnding: false,                         // 是否跳过行尾
		LineEnding:     zapcore.DefaultLineEnding,     // 行尾字符
		EncodeLevel:    consoleLevelEncode,            // 自定义日志级别的输出格式
		EncodeTime:     consoleTimeEncoder,            // 自定义时间格式, 这里也可以使用 zapcore.ISO8601TimeEncoder
		EncodeDuration: zapcore.StringDurationEncoder, // 调用消耗时长
		EncodeCaller:   consoleCallerEncoder,          // 自定义函数调用输出格式
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

//
// consoleTimeEncoder 自定义时间格式显示
//	@parameter	ts 日志时间
//	@parameter	enc 编码器接口
//
func consoleTimeEncoder(ts time.Time, enc zapcore.PrimitiveArrayEncoder) {
	//enc.AppendString("[" + ts.Format(logTimeLayout) + "]")
	enc.AppendString("[" + ts.Format(iso8601TimeLayout) + "]")
}

//
// consoleLevelEncode 自定义日志级别显示
//	@parameter	level 日志级别
//	@parameter	enc 编码器接口
//
func consoleLevelEncode(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	formattedLevel := fmt.Sprintf("[%5s]", level.CapitalString())
	//enc.AppendString("[" + level.CapitalString() + "]")
	enc.AppendString(formattedLevel)
}

//
// consoleCallerEncoder 自定义函数调用显示
//	@parameter	entryCaller 记录函数调用者的结构体对象
//	@parameter	enc 编码器接口
//
func consoleCallerEncoder(entryCaller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	formattedCaller := fmt.Sprintf("[%30s]", entryCaller.TrimmedPath())
	//enc.AppendString("[" + entryCaller.TrimmedPath() + "]")
	enc.AppendString(formattedCaller)
}