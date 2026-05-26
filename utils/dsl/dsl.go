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

package dsl

import (
	"github.com/rulego/rulego/api/types"
)

const FieldNameScript = "script"

// ParseCrossNodeDependencies 解析规则链中的跨节点依赖关系，返回每个节点依赖的节点ID列表（仅包含在规则链中定义的节点）
// ParseCrossNodeDependencies parses cross-node dependencies in the rule chain and returns dependent node IDs for each node (only includes nodes defined in the rule chain)
func ParseCrossNodeDependencies(def types.RuleChain) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

// Remove duplicates and filter only defined nodes

// 只添加在规则链中实际定义的节点ID
// Only add node IDs that are actually defined in the rule chain

// GetReferencedNodeIds 获取规则链中所有被引用且在规则链中定义的节点ID列表（去重）
// GetReferencedNodeIds gets all referenced node IDs that are defined in the rule chain (deduplicated)
func GetReferencedNodeIds(def types.RuleChain) []string { _ = "STUB: not implemented"; return nil }

// 只添加在规则链中实际定义的节点ID
// Only add node IDs that are actually defined in the rule chain

// Convert set to slice

// IsNodeIdDefined 检查给定的nodeId是否在规则链的节点定义中
// IsNodeIdDefined checks if the given nodeId is defined in the rule chain nodes
func IsNodeIdDefined(def types.RuleChain, nodeId string) bool {
	_ = "STUB: not implemented"
	return false
}

// ExtractReferencedNodeIds 从节点配置中提取被引用的节点ID列表（支持嵌套字段）
// ExtractReferencedNodeIds extracts referenced node IDs from node configuration (supports nested fields)
func ExtractReferencedNodeIds(configuration types.Configuration) []string {
	_ = "STUB: not implemented"
	return nil
}

// extractNodeIdsFromValue 递归提取任意类型值中的节点引用
// extractNodeIdsFromValue recursively extracts node references from values of any type
func extractNodeIdsFromValue(value interface{}, uniqueNodeIds map[string]bool, nodeIds *[]string) {
	_ = "STUB: not implemented"
	return
}

// 提取字符串中的节点引用，支持 ${nodeId.msg.xx} 和 nodeId.msg.xx 格式
// Extract node references from string, supports ${nodeId.msg.xx} and nodeId.msg.xx formats

// 递归处理map类型
// Recursively process map type

// 递归处理slice类型
// Recursively process slice type

// 递归处理Configuration类型
// Recursively process Configuration type

// 对于其他类型（int, bool, float等），不包含节点引用，直接忽略
// For other types (int, bool, float, etc.), no node references, ignore

// 不处理其他类型
// Do not process other types

// BuiltinVars 内置变量列表，这些不应该被识别为节点ID
// Built-in variables list, these should not be recognized as node IDs
var BuiltinVars = map[string]bool{
	"msg":      true,
	"metadata": true,
	"msgType":  true,
	"global":   true,
	"vars":     true,
	"len":      true,
	"string":   true,
	"int":      true,
	"float":    true,
	"bool":     true,
	"true":     true,
	"false":    true,
}

// ExtractNodeReferencesFromExpression 从表达式内容中提取节点引用
// ExtractNodeReferencesFromExpression extracts node references from expression content
func ExtractNodeReferencesFromExpression(expression string) []string {
	_ = "STUB: not implemented"
	return nil
}

// 使用正则表达式来匹配节点引用
// Use regex to match node references
// 匹配 nodeId.data, nodeId.msg, nodeId.metadata, nodeId.id, nodeId.ts, nodeId.dataType, nodeId.global, nodeId.vars 的模式，确保前面不是点号
// nodeId 支持字母、数字、下划线、中划线和斜杠
// Match nodeId.data, nodeId.msg, nodeId.metadata, etc. patterns, ensuring not preceded by a dot
// nodeId supports letters, numbers, underscores, hyphens and slashes

// 排除内置变量，只处理真正的跨节点引用
// Exclude built-in variables, only process real cross-node references

// ParseVars 解析规则链中的变量
func ParseVars(varPrefix string, def types.RuleChain, includeNodeId ...string) []string {
	_ = "STUB: not implemented"
	return nil
}

//脚本通过 {varPrefix}.xx 方式解析

//通过 ${{varPrefix}.xx} 方式解析

// IsFlowNode 判断是否是子规则链
func IsFlowNode(def types.RuleChain, nodeId string) bool { _ = "STUB: not implemented"; return false }

// ProcessVariables replaces placeholders in the node configuration with global and chain-specific variables.
func ProcessVariables(config types.Config, ruleChainDef types.RuleChain, from types.Configuration) types.Configuration {
	_ = "STUB: not implemented"
	return *new(types.Configuration)
}

func GetInitNodeEnv(config types.Config, ruleChainDef types.RuleChain) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}
