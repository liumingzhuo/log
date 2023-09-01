package log

import "context"

// InfoLogger infoLogger 记录非错误日志
type InfoLogger interface {
	Info(msg string, fields ...Field)
	Infof(template string, args ...interface{})
	Infow(format string, keysAndValues ...interface{})

	Enable() bool
}

// Logger 日志信息，包括错误日志以及非错误日志
type Logger interface {
	InfoLogger

	Debug(msg string, fields ...Field)
	Debugf(template string, args ...interface{})
	Debugw(format string, keysAndValues ...interface{})

	Warn(msg string, fields ...Field)
	Warnf(templace string, args ...interface{})
	Warnw(format string, keysAndValues ...interface{})

	Error(msg string, fields ...Field)
	Errorf(templace string, args ...interface{})
	Errorw(format string, keysAndValues ...interface{})

	Panic(msg string, fields ...Field)
	Panicf(templace string, args ...interface{})
	Panicw(format string, keysAndValues ...interface{})

	Fatal(msg string, fields ...Field)
	Fatalf(templace string, args ...interface{})
	Fatalw(format string, keysAndValues ...interface{})

	// V  返回一个特定等级level的InfoLogger对象，
	// level值越大代表日志级别越低
	// level需要大于0
	V(level int) InfoLogger

	// WithValue 为Logger添加keyvalue对
	WithValue(keysAndValues ...interface{}) Logger

	// WithName 向Logger的添加新的名字元素
	WithName(name string) Logger

	// WithContext 拷贝一个context，并且设置了日志值
	WithContext(ctx context.Context) context.Context

	// Flush 调用底层Core的Sync方法 刷新缓存日志
	Flush()
}
