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

package assert

import (
	"sync"
	"testing"
)

// assertMutex 全局互斥锁，用于保护testing.T的并发访问
var assertMutex sync.Mutex

// CallerInfo This function is inspired by:
// https://github.com/stretchr/testify/blob/master/assert/assertions.go
func CallerInfo() []string { _ = "STUB: not implemented"; return nil }

// The breaks below failed to terminate the loop, and we ran off the
// end of the call stack.

// This is a huge edge case, but it will panic if this is the case, see #180

// testing.tRunner is the standard library function that calls
// tests. Subtests are called directly by tRunner, without going through
// the Test/Benchmark/Example function that contains the t.Run calls, so
// with subtests we should break when we hit tRunner, without adding it
// to the list of callers.

func Equal(t *testing.T, a, b interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func EqualCleanString(t *testing.T, a, b string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotEqual(t *testing.T, a, b interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func True(t *testing.T, value bool, msgAndArgs ...interface{}) { _ = "STUB: not implemented"; return }

func False(t *testing.T, value bool, msgAndArgs ...interface{}) { _ = "STUB: not implemented"; return }

func NotNil(t *testing.T, value interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Nil(t *testing.T, value interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func messageFromMsgAndArgs(msgAndArgs ...interface{}) string { _ = "STUB: not implemented"; return "" }

func isNil(v interface{}) bool { _ = "STUB: not implemented"; return false }
