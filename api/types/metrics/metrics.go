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

package metrics

// EngineMetrics holds various metrics for the rule engine execution.
type EngineMetrics struct {
	Current int64 // Number of currently executing engine
	Total   int64 // Total number of engine executions
	Failed  int64 // Number of failed chains executions
	Success int64 // Number of successful chains executions
}

// NewEngineMetrics creates a new instance of EngineMetrics.
func NewEngineMetrics() *EngineMetrics { _ = "STUB: not implemented"; return nil }

// IncrementCurrent increases the count of current executions.
func (m *EngineMetrics) IncrementCurrent() { _ = "STUB: not implemented"; return }

// DecrementCurrent decreases the count of current executions.
func (m *EngineMetrics) DecrementCurrent() { _ = "STUB: not implemented"; return }

// IncrementTotal increases the total count of executions.
func (m *EngineMetrics) IncrementTotal() { _ = "STUB: not implemented"; return }

// IncrementFailed increases the count of failed executions.
func (m *EngineMetrics) IncrementFailed() { _ = "STUB: not implemented"; return }

// IncrementSuccess increases the count of successful executions.
func (m *EngineMetrics) IncrementSuccess() { _ = "STUB: not implemented"; return }

// Get returns a copy of the current metrics.
func (m *EngineMetrics) Get() EngineMetrics { _ = "STUB: not implemented"; return *new(EngineMetrics) }

// Reset resets all metrics to zero.
func (m *EngineMetrics) Reset() { _ = "STUB: not implemented"; return }
