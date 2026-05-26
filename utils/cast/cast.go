/*
 * Copyright 2025 The RuleGo Authors.
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

package cast

import (
	"time"
)

// ToInt converts an interface{} to int.
// It returns 0 if conversion fails.
func ToInt(value interface{}) int { _ = "STUB: not implemented"; return 0 }

// ToIntE converts an interface{} to int with error handling.
// Returns the converted int value and nil error if successful.
// Returns 0 and an error if conversion fails.
func ToIntE(value interface{}) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// ToInt64 converts an interface{} to int64.
// It returns 0 if conversion fails.
func ToInt64(value interface{}) int64 { _ = "STUB: not implemented"; return 0 }

// ToInt64E converts an interface{} to int64 with error handling.
// Returns the converted int64 value and nil error if successful.
// Returns 0 and an error if conversion fails.
func ToInt64E(value interface{}) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// ToDurationE converts an interface{} to time.Duration with error handling.
// Returns the converted duration value and nil error if successful.
// Returns 0 and an error if conversion fails.
func ToDurationE(value interface{}) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// ToBool converts an interface{} to bool.
// It returns false if conversion fails.
func ToBool(value interface{}) bool { _ = "STUB: not implemented"; return false }

// ToBoolE converts an interface{} to bool with error handling.
// Returns the converted bool value and nil error if successful.
// Returns false and an error if conversion fails.
func ToBoolE(value interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ToFloat64 converts an interface{} to float64.
// It returns 0 if conversion fails.
func ToFloat64(value interface{}) float64 { _ = "STUB: not implemented"; return 0 }

// ToFloat64E converts an interface{} to float64 with error handling.
// Returns the converted float64 value and nil error if successful.
// Returns 0 and an error if conversion fails.
func ToFloat64E(value interface{}) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// ToString converts an interface{} to string.
// It returns empty string if conversion fails.
func ToString(input interface{}) string { _ = "STUB: not implemented"; return "" }

// ToStringE converts an interface{} to string with error handling.
// Returns the converted string value and nil error if successful.
// Returns empty string and an error if conversion fails.
func ToStringE(input interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

// 转换为 map[string]interface{}

// ConvertIntToTime 将整数时间戳转换为 time.Time
func ConvertIntToTime(timestampInt int64, timeUnit time.Duration) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// 默认按秒处理
