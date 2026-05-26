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
	"github.com/rulego/rulego/api/types"
)

// init 注册IteratorNode组件
// init registers the IteratorNode component with the default registry.
func init() {
	//Registry.Add(&IteratorNode{}) //弃用，使用 for 组件代替
}

// IteratorNodeConfiguration IteratorNode配置结构
// IteratorNodeConfiguration defines the configuration structure for the IteratorNode component.
type IteratorNodeConfiguration struct {
	// FieldName 要遍历的字段名称，支持点符号嵌套访问
	// FieldName specifies the field name to iterate over.
	// If empty, iterates over the entire message.
	// Supports nested field access using dot notation (e.g., "items.value", "items").
	FieldName string

	// JsScript 可选的JavaScript过滤脚本
	// JsScript specifies optional JavaScript code for filtering items.
	// If empty, all items pass the filter.
	// The script should implement: function ItemFilter(item, index, metadata)
	//   - item: Current item being iterated
	//   - index: Array index (for arrays) or key (for objects)
	//   - metadata: Message metadata for context
	// Returns: boolean indicating whether item should be routed via True relation
	JsScript string
}

// IteratorNode 遍历消息数据中数组或对象的动作组件
// IteratorNode is an action component that iterates over arrays or objects in message data.
//
// 弃用通知 - Deprecation Notice:
//   - 此组件已弃用，请使用ForNode获得更好的性能和功能 - This component is deprecated, use ForNode for better performance and features
//
// 核心算法：
// Core Algorithm:
// 1. 提取并解析要遍历的数据（支持JSON） - Extract and parse data to iterate (supports JSON)
// 2. 根据FieldName提取特定字段或使用整个消息 - Extract specific field by FieldName or use entire message
// 3. 遍历数组或对象的每个元素 - Iterate over each element in array or object
// 4. 对每个元素应用JavaScript过滤器（如果配置） - Apply JavaScript filter to each element (if configured)
// 5. 根据过滤结果路由到True/False关系 - Route to True/False relations based on filter results
// 6. 遍历完成后通过Success关系发送原始消息 - Send original message via Success relation after iteration
//
// 支持的数据类型 - Supported data types:
//   - []interface{}: 数组，使用数字索引 - Arrays with numeric indices
//   - map[string]interface{}: 对象，使用字符串键 - Objects with string keys
//
// JavaScript过滤函数签名 - JavaScript filter function signature:
//   - function ItemFilter(item, index, metadata) -> boolean
type IteratorNode struct {
	// Config 节点配置
	// Config holds the node configuration including field name and JavaScript filter
	Config IteratorNodeConfiguration

	// jsEngine JavaScript引擎实例
	// jsEngine holds the JavaScript engine instance for item filtering
	jsEngine types.JsEngine
}

// Type 返回组件类型
// Type returns the component type identifier.
func (x *IteratorNode) Type() string {
	_ = "STUB: not implemented"

	// New 创建新实例
	// New creates a new instance.
	return ""
}

func (x *IteratorNode) New() types.Node { _ = "STUB: not implemented"; return *new(types.Node) }

// Init 初始化组件
// Init initializes the component.
func (x *IteratorNode) Init(ruleConfig types.Config, configuration types.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

// OnMsg 处理消息，遍历指定字段或整个消息，应用JavaScript过滤器
// OnMsg processes incoming messages by iterating over the specified field or entire message.
func (x *IteratorNode) OnMsg(ctx types.RuleContext, msg types.RuleMsg) {
	_ = "STUB: not implemented"
	return
}

// 遍历指定字段

//出现错误中断遍历

//出现错误中断遍历

// Destroy 清理资源
// Destroy cleans up resources.
func (x *IteratorNode) Destroy() {
	_ = "STUB: not implemented"
	// 无资源需要清理
	// No resources to clean up
	return
}

// executeItem 处理每个遍历项，应用JavaScript过滤器并路由
// executeItem processes each individual item during iteration.
func (x *IteratorNode) executeItem(ctx types.RuleContext, msg types.RuleMsg, item interface{}, index interface{}) error {
	_ = "STUB: not implemented"
	return nil

	// 使用零拷贝GetReadOnlyValues，JS引擎只读取metadata
}

//出现错误中断遍历
