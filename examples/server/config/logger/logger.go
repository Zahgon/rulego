package logger

import "log"

// DefaultLogger 默认日志记录器包装，实现了 types.Logger 接口
type DefaultLogger struct {
	*log.Logger
}

// Debugf 记录调试级别的日志
func (l *DefaultLogger) Debugf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Infof 记录信息级别的日志
func (l *DefaultLogger) Infof(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Warnf 记录警告级别的日志
func (l *DefaultLogger) Warnf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Errorf 记录错误级别的日志
func (l *DefaultLogger) Errorf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Logger 暴露给外部的日志实例
var Logger *DefaultLogger

// Set 设置全局日志实例
func Set(logger *log.Logger) { _ = "STUB: not implemented"; return }

// Get 获取全局日志实例
func Get() *DefaultLogger { _ = "STUB: not implemented"; return nil }
