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

// Package runtime provides utilities for runtime-related operations.
// This package includes functions for retrieving stack traces and other
// runtime information that can be useful for debugging and logging purposes.
//
// The Stack function in this package returns a formatted string containing
// the stack trace of the current goroutine, excluding the first two stack frames.
// This is particularly useful for error reporting and debugging complex call chains.
//
// Usage example:
//
//	stackTrace := runtime.Stack()
//	fmt.Println("Current stack trace:", stackTrace)
//
// Note: The stack trace includes file names and line numbers for each call in the stack.
package runtime

// Stack 获取堆栈信息
func Stack() string { _ = "STUB: not implemented"; return "" }
