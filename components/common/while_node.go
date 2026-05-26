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

package common

import (
	"context"

	"github.com/rulego/rulego/api/types"
	"github.com/rulego/rulego/utils/el"
)

func init() {
	// Register the WhileNode with the component registry on initialization.
	Registry.Add(&WhileNode{})
}

// WhileNodeConfiguration defines the configuration for the WhileNode.
type WhileNodeConfiguration struct {
	// Condition is the expression to check before each iteration.
	// If the expression evaluates to true, the loop continues.
	// Uses 'el' expression language (e.g., "${msg.count} < 10").
	Condition string
	// Do specifies the node or sub-rule chain to process in each iteration,
	// e.g., "s1" or "chain:rule01".
	Do string
	// Mode 0:不处理msg，1：合并遍历msg.Data，2：替换msg
	Mode int
}

// WhileNode provides a while-loop structure.
// It executes the 'Do' node/chain repeatedly as long as the 'Condition' evaluates to true.
//
// WhileNode 提供 while 循环结构。
// 只要 'Condition' 评估为真，它就会重复执行 'Do' 节点/链。
//
// Configuration:
// 配置说明：
//
//	{
//		"condition": "${msg.count} < 5", // Expression to check  检查表达式
//		"do": "s3",                      // Target node ID or sub-chain  目标节点ID或子链
//		"mode": 1                        // Processing mode: 0=DoNotProcess (default), 1=MergeValues, 2=ReplaceValues  处理模式
//	}
type WhileNode struct {
	// Config contains the node configuration.
	Config WhileNodeConfiguration
	// ruleNodeId is the parsed target for the 'Do' action.
	ruleNodeId types.RuleNodeId
	// conditionTemplate is the compiled template for the condition.
	conditionTemplate el.Template
}

// Type returns the component type.
func (x *WhileNode) Type() string { _ = "STUB: not implemented"; return "" }

func (x *WhileNode) New() types.Node { _ = "STUB: not implemented"; return *new(types.Node) }

// Init initializes the node.
func (x *WhileNode) Init(_ types.Config, configuration types.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *WhileNode) toMap(data string) interface{} { _ = "STUB: not implemented"; return nil }

func (x *WhileNode) toList(dataType types.DataType, itemDataList []string) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// OnMsg processes the message.
func (x *WhileNode) OnMsg(ctx types.RuleContext, msg types.RuleMsg) {
	_ = "STUB: not implemented"

	// Create a context with cancel for the loop execution
	return
}

// Update loop index in metadata

// Check condition

// Parse result to boolean

// Execute the iteration

// Always pass the updated msg to the next iteration so the condition can evaluate it

// Check for break signal

// msg is already lastMsg (final state).

// Destroy cleans up resources.
func (x *WhileNode) Destroy() {
	_ = "STUB: not implemented"

	// executeItem processes the 'Do' node/chain.
	return
}

func (x *WhileNode) executeItem(ctxWithCancel context.Context, ctx types.RuleContext, fromMsg types.RuleMsg) (types.RuleMsg, []string, error) {
	_ = "STUB: not implemented"
	return *new(types.RuleMsg), nil, nil
}

// Prepare callback

func (x *WhileNode) formDoVar() error { _ = "STUB: not implemented"; return nil }

// castToBool converts interface{} to bool.
func castToBool(val interface{}) bool { _ = "STUB: not implemented"; return false }
