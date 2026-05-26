/*
 * Copyright 2023 The RuleGo Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package types

import (
	"io"
	"log"
	"sync"
)

// ============================================
// 日志级别
// ============================================

// LogLevel 日志级别
type LogLevel int8

const (
	// DebugLevel 调试级别
	DebugLevel LogLevel = iota - 1
	// InfoLevel 信息级别
	InfoLevel
	// WarnLevel 警告级别
	WarnLevel
	// ErrorLevel 错误级别
	ErrorLevel
)

// String 返回日志级别的字符串表示
func (l LogLevel) String() string { _ = "STUB: not implemented"; return "" }

// ============================================
// Field 结构化字段（可选扩展）
// ============================================

// Field 结构化日志字段
type Field struct {
	Key   string
	Value any
}

// F 创建日志字段的快捷方法
func F(key string, value any) Field { _ = "STUB: not implemented"; return *new(Field) }

// ============================================
// Logger 接口
// ============================================

// Logger 日志接口
// 应用层需要实现此接口以接入自己的日志框架
type Logger interface {
	// Printf 兼容旧接口
	// Deprecated: 请使用 Debugf/Infof/Warnf/Errorf 代替
	Printf(format string, v ...interface{})
	// Debugf 调试日志
	Debugf(format string, v ...interface{})
	// Infof 信息日志
	Infof(format string, v ...interface{})
	// Warnf 警告日志
	Warnf(format string, v ...interface{})
	// Errorf 错误日志
	Errorf(format string, v ...interface{})
}

// ============================================
// 默认实现
// ============================================

// 确保 StdLogger 实现 Logger 接口
var _ Logger = (*StdLogger)(nil)

// 确保标准库 log.Logger 兼容 Printf 接口
var _ interface{ Printf(string, ...interface{}) } = (*log.Logger)(nil)

// StdLogger 默认日志实现
type StdLogger struct {
	mu    sync.Mutex
	w     io.Writer
	level LogLevel
}

// NewStdLogger 创建标准日志器
func NewStdLogger(w io.Writer) *StdLogger { _ = "STUB: not implemented"; return nil }

// SetLevel 设置日志级别
func (l *StdLogger) SetLevel(level LogLevel) { _ = "STUB: not implemented"; return }

// GetLevel 获取日志级别
func (l *StdLogger) GetLevel() LogLevel { _ = "STUB: not implemented"; return *new(LogLevel) }

// SetOutput 设置输出目标
func (l *StdLogger) SetOutput(w io.Writer) { _ = "STUB: not implemented"; return }

func (l *StdLogger) output(level LogLevel, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Printf 实现 Logger 接口
func (l *StdLogger) Printf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Debugf 调试日志
func (l *StdLogger) Debugf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Infof 信息日志
func (l *StdLogger) Infof(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Warnf 警告日志
func (l *StdLogger) Warnf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Errorf 错误日志
func (l *StdLogger) Errorf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// ============================================
// 工厂函数
// ============================================

// DefaultLogger 返回默认日志器
func DefaultLogger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

// NewLogger 创建日志器
func NewLogger(custom Logger) Logger { _ = "STUB: not implemented"; return *new(Logger) }

// IsNilLogger 判断日志器是否为空或使用空实现
// 用于在打印日志前进行判断，避免空指针
func IsNilLogger(logger Logger) bool { _ = "STUB: not implemented"; return false }
