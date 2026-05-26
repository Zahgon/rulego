/*
 * Copyright 2024 The RuleGo Authors.
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

// Package funcs provides built-in function implementations for the RuleGo rule engine.
//
// This package implements various functions that can be used within rule chains,
// allowing for custom logic to be integrated into rule chains. It includes pre-defined
// functions such as string manipulation and mathematical operations, as well as the
// ability to create custom functions.
//
// Key components:
// - FuncMap: A generic map for storing functions
// - TemplateFunc: A map for storing template functions
// - ScriptFunc: A map for storing script functions
//
// The package supports features such as:
// - Registering and accessing functions by name
// - Retrieving all registered functions
// - Unregistering functions
// - Retrieving function names
package funcs

import (
	"strings"
	"sync"
)

// TemplateFunc 内置模板函数
var TemplateFunc = NewFuncMap[any]()

// ScriptFunc 内置Js用户函数
var ScriptFunc = NewFuncMap[any]()

func init() {
	TemplateFunc.Register("escape", func(s string) string {
		var replacer = strings.NewReplacer(
			"\\", "\\\\", // 反斜杠
			"\"", "\\\"", // 双引号
			"\n", "\\n", // 换行符
			"\r", "\\r", // 回车符
			"\t", "\\t", // 制表符
		)
		return replacer.Replace(s)
	})
}

// FuncMap 是一个泛型映射，用于存储函数
type FuncMap[T any] struct {
	v map[string]T
	sync.RWMutex
}

// NewFuncMap 创建一个新的FuncMap实例
func NewFuncMap[T any]() *FuncMap[T] { _ = "STUB: not implemented"; return nil }

func (x *FuncMap[T]) Register(name string, value T) { _ = "STUB: not implemented"; return }

func (x *FuncMap[T]) RegisterAll(values map[string]T) { _ = "STUB: not implemented"; return }

func (x *FuncMap[T]) UnRegister(name string) { _ = "STUB: not implemented"; return }

func (x *FuncMap[T]) Get(name string) (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (x *FuncMap[T]) GetAll() map[string]T { _ = "STUB: not implemented"; return nil }

func (x *FuncMap[T]) Names() []string { _ = "STUB: not implemented"; return nil }
