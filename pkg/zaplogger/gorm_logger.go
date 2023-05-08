package zaplogger

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type GormLogger struct {
	ZapLogger                 *zap.Logger     // zap.Logger 对象的指针
	SlowThreshold             time.Duration   // 慢 SQL 周期
	LogLevel                  logger.LogLevel // 日志级别
	IgnoreRecordNotFoundError bool            // 是否忽略 RecordNotFoundError 错误
	SkipCallerLookup          bool            // 是否跳过函数调用查找
}

func NewGormLogger(zapLogger *zap.Logger) *GormLogger {
	return &GormLogger{
		ZapLogger:                 zapLogger,
		LogLevel:                  logger.Info,
		SlowThreshold:             time.Millisecond * 100,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: false,
	}
}

//
// LogMode 设置日志记录模式
//	@parameter	level 日志级别
//	@return		logger.Interface
//
func (g GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return GormLogger{
		ZapLogger:                 g.ZapLogger,
		SlowThreshold:             g.SlowThreshold,
		LogLevel:                  g.LogLevel,
		IgnoreRecordNotFoundError: g.IgnoreRecordNotFoundError,
		SkipCallerLookup:          g.SkipCallerLookup,
	}
}

//
// Info 输出 Info 级别的日志信息
//	@parameter	ctx
//	@parameter	s 要输出的日志文本
//	@parameter	i
//
func (g GormLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	if g.LogLevel < logger.Info {
		return
	}
	g.ZapLogger.Sugar().Debugf(msg, args)
}

func (g GormLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	if g.LogLevel < logger.Warn {
		return
	}
	g.ZapLogger.Sugar().Warnf(msg, args...)
}

func (g GormLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	if g.LogLevel < logger.Error {
		return
	}
	g.ZapLogger.Sugar().Errorf(msg, args...)
}

func (g GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)

	switch {
	case err != nil && g.LogLevel >= logger.Error && (!g.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		g.ZapLogger.Error("trace",
			zap.Error(err),
			zap.Duration("elapsedTime", elapsed),
			zap.Int64("rowsAffected", rows),
			zap.String("sql", sql),
		)
	case g.SlowThreshold != 0 && elapsed > g.SlowThreshold && g.LogLevel >= logger.Warn: //警告禁止
		sql, rows := fc()
		g.ZapLogger.Warn("warn",
			zap.Duration("elapsedTime", elapsed),
			zap.Int64("rowsAffected", rows),
			zap.String("sql", sql),
		)
	case g.LogLevel >= logger.Info:
		sql, rows := fc()
		g.ZapLogger.Debug("info",
			zap.Duration("elapsedTime", elapsed),
			zap.Int64("rowsAffected", rows),
			zap.String("sql", sql),
		)

	}
}
