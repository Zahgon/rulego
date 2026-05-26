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

package external

import (
	"github.com/rulego/rulego/utils/el"

	"github.com/rulego/rulego/api/types"
)

// 注册节点
func init() {
	Registry.Add(&CacheGetNode{})
	Registry.Add(&CacheSetNode{})
	Registry.Add(&CacheDeleteNode{})
}

const (
	CacheLevelChain                = "chain"
	CacheLevelGlobal               = "global"
	CacheOutputModeMergeToMetadata = 0   // 合并到当前消息元数据
	CacheOutputModeMergeToMsg      = 1   //合并到当前消息负荷
	CacheOutputModeNewMsg          = 2   //覆盖原消息负荷输出
	KeyMatchAll                    = "*" //通配符
)

// LevelKey 缓存key
type LevelKey struct {
	// Level 缓存级别，chain或global
	// chain: 规则链缓存，在当前规则链命名空间下操作，用于规则链实例内不同执行上下文之间的数据共享。如果规则链实例被销毁，会自动删除该规则链命名空间下所有缓存
	// global: 全局缓存，在全局命名空间下操作，用于跨规则链间的数据共享
	Level string `json:"level"`
	// 可以使用 ${metadata.key} 读取元数据中的变量或者使用 ${msg.key} 读取消息负荷中的变量进行替换
	// 支持 * 通配符查找，例如：test:* 表示以test: 开头的所有key
	Key string `json:"key"`
}

// LevelKeyTemplate 缓存key模板
type LevelKeyTemplate struct {
	level       string      // 缓存级别
	keyTemplate el.Template // 缓存key模板
}

const (
	// WhenKeyNotFoundFailure 找不到key时走失败链
	WhenKeyNotFoundFailure = "failure"
	// WhenKeyNotFoundSuccess 找不到key时走成功链
	WhenKeyNotFoundSuccess = "success"
)

// CacheGetNodeConfiguration 缓存获取节点配置
type CacheGetNodeConfiguration struct {
	// Keys 获取key列表
	Keys []LevelKey `json:"keys"`
	// OutputMode 输出模式
	// 0:查询结果，合并到当前消息元数据
	// 1:查询结果，合并到当前消息负荷。要求输入消息负荷`DataType`必须是JSON类型，并且消息负荷`Data`可以解析为map结构
	// 2:查询结果，转成JSON，覆盖原消息负荷输出
	OutputMode int `json:"outputMode"`
	// WhenKeyNotFound 找不到key时的处理策略
	// "": 保持原有行为（Mode 0/1 走成功链，Mode 2 全 miss 走失败链）
	// "failure": 找不到key时走失败链
	// "success": 找不到key时走成功链
	WhenKeyNotFound string `json:"whenKeyNotFound"`
}

// CacheGetNode retrieves data from cache storage at different levels (chain or global).
// It supports wildcard pattern matching and multiple output modes for flexible data integration.
//
// CacheGetNode 从不同级别（链级或全局）的缓存存储中检索数据。
// 支持通配符模式匹配和多种输出模式以实现灵活的数据集成。
//
// Configuration:
// 配置说明：
//
//	{
//		"keys": [                        // Keys to retrieve  要检索的键列表
//			{
//				"level": "chain",        // Cache level: "chain" or "global"  缓存级别
//				"key": "sensor_${metadata.deviceId}"  // Key with variable substitution  支持变量替换的键
//			},
//			{
//				"level": "global",
//				"key": "config_*"        // Wildcard pattern for multiple keys  多键通配符模式
//			}
//		],
//		"outputMode": 0                  // Output mode: 0=metadata, 1=merge to msg, 2=replace msg  输出模式
//	}
//
// Cache Levels:
// 缓存级别：
//
//   - "chain": Rule chain scoped cache for data sharing within the same rule chain instance
//     规则链级缓存，用于同一规则链实例内的数据共享
//   - "global": Global cache for cross-chain data sharing across all rule chain instances
//     全局缓存，用于所有规则链实例间的跨链数据共享
//
// Key Pattern Matching:
// 键模式匹配：
//
//   - Exact match: "user:123" retrieves specific key  精确匹配：检索特定键
//   - Wildcard: "user:*" retrieves all keys with prefix "user:"  通配符：检索所有前缀为 "user:" 的键
//   - Variable substitution: "data_${metadata.id}" uses runtime values  变量替换：使用运行时值
//
// Output Modes:
// 输出模式：
//
//   - 0 (CacheOutputModeMergeToMetadata): Merge results to message metadata
//     合并结果到消息元数据
//   - 1 (CacheOutputModeMergeToMsg): Merge to message payload (requires JSON data type)
//     合并到消息负荷（需要 JSON 数据类型）
//   - 2 (CacheOutputModeNewMsg): Replace message payload with cache results
//     用缓存结果替换消息负荷
//
// Usage Example:
// 使用示例：
//
//	// Retrieve user session and global config
//	// 检索用户会话和全局配置
//	{
//		"id": "cacheGet",
//		"type": "cacheGet",
//		"configuration": {
//			"keys": [
//				{"level": "chain", "key": "session_${metadata.userId}"},
//				{"level": "global", "key": "app_config_*"}
//			],
//			"outputMode": 1
//		}
//	}
type CacheGetNode struct {
	//节点配置
	Config CacheGetNodeConfiguration
	//keys模板
	keysTemplate []LevelKeyTemplate
}

func (x *CacheGetNode) Type() string { _ = "STUB: not implemented"; return "" }

func (x *CacheGetNode) New() types.Node { _ = "STUB: not implemented"; return *new(types.Node) }

// Init 初始化组件
func (x *CacheGetNode) Init(ruleConfig types.Config, configuration types.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

//初始化keys模板

func (x *CacheGetNode) OnMsg(ctx types.RuleContext, msg types.RuleMsg) {
	_ = "STUB: not implemented"
	return
}

//处理keys模板

// Destroy 销毁组件
func (x *CacheGetNode) Destroy() { _ = "STUB: not implemented"; return }

func (x *CacheGetNode) handleGet(ctx types.RuleContext, msg types.RuleMsg, keys []LevelKey) {
	_ = "STUB: not implemented"
	return
}

// cache miss，由 whenKeyNotFound 控制

// 底层报错，始终走失败链

func (x *CacheGetNode) outputResult(ctx types.RuleContext, msg types.RuleMsg, values map[string]interface{}) {
	_ = "STUB: not implemented"
	// 检查是否全部 miss
	return
}

// 走成功链，继续下面的输出逻辑

// 空值：保持原有行为，Mode 2 走失败链，其他模式走成功链

// CacheSetNodeConfiguration 缓存设置节点配置
type CacheSetNodeConfiguration struct {
	// Items 缓存项列表
	Items []CacheItem `json:"items"`
}

type CacheItem struct {
	// Level 缓存级别，chain或global
	// chain: 规则链缓存，在当前规则链命名空间下操作，用于规则链实例内不同执行上下文之间的数据共享。如果规则链实例被销毁，会自动删除该规则链命名空间下所有缓存
	// global: 全局缓存，在全局命名空间下操作，用于跨规则链间的数据共享
	Level string `json:"level"`
	//  key 键
	// 可以使用 ${metadata.key} 读取元数据中的变量或者使用 ${msg.key} 读取消息负荷中的变量进行替换
	Key string `json:"key"`
	// value 值
	// 可以使用 ${metadata.key} 读取元数据中的变量或者使用 ${msg.key} 读取消息负荷中的变量进行替换
	Value interface{} `json:"value"`
	// Ttl 过期时间
	// 示例：1h(1小时) 1h30m(1小时30分钟) 10m(10分钟) 10s(10秒)，如果为空或者0，则表示永不过期
	Ttl string `json:"ttl"`
}

type CacheItemTemplate struct {
	level         string
	keyTemplate   el.Template
	valueTemplate el.Template
	ttl           string
}

// CacheSetNode stores data in cache storage at different levels with TTL support.
// It supports multiple cache items, variable substitution, and automatic expiration.
//
// CacheSetNode 在不同级别的缓存存储中存储数据，支持 TTL。
// 支持多个缓存项、变量替换和自动过期。
//
// Configuration:
// 配置说明：
//
//	{
//		"items": [                       // Cache items to set  要设置的缓存项
//			{
//				"level": "chain",        // Cache level: "chain" or "global"  缓存级别
//				"key": "user_${metadata.userId}",     // Key with variable substitution  支持变量替换的键
//				"value": "${msg.userData}",           // Value with variable substitution  支持变量替换的值
//				"ttl": "1h30m"          // TTL format: 1h30m, 10m, 30s, empty=no expiration  TTL 格式
//			}
//		]
//	}
//
// TTL Format:
// TTL 格式：
//
//   - "1h": 1 hour  1小时
//   - "30m": 30 minutes  30分钟
//   - "10s": 10 seconds  10秒
//   - "1h30m": 1 hour 30 minutes  1小时30分钟
//   - "": No expiration (permanent storage)  无过期时间（永久存储）
//
// Variable Substitution:
// 变量替换：
//
// Both keys and values support runtime variable substitution:
// 键和值都支持运行时变量替换：
//   - ${metadata.key}: Access message metadata  访问消息元数据
//   - ${msg.key}: Access message payload fields  访问消息负荷字段
//
// Usage Example:
// 使用示例：
//
//	// Store user session with 1-hour expiration
//	// 存储用户会话，1小时过期
//	{
//		"id": "cacheSet",
//		"type": "cacheSet",
//		"configuration": {
//			"items": [
//				{
//					"level": "chain",
//					"key": "session_${metadata.userId}",
//					"value": "${msg.sessionData}",
//					"ttl": "1h"
//				},
//				{
//					"level": "global",
//					"key": "last_activity_${metadata.userId}",
//					"value": "${metadata.timestamp}",
//					"ttl": "24h"
//				}
//			]
//		}
//	}
type CacheSetNode struct {
	// 节点配置
	Config CacheSetNodeConfiguration
	// 缓存项列表模板
	itemsTemplate []CacheItemTemplate
	// 缓存项列表模板是否有变量
	hasVar bool
}

// Type 返回组件类型
func (x *CacheSetNode) Type() string { _ = "STUB: not implemented"; return "" }

func (x *CacheSetNode) New() types.Node { _ = "STUB: not implemented"; return *new(types.Node) }

// Init 初始化组件
func (x *CacheSetNode) Init(ruleConfig types.Config, configuration types.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

//初始化缓存项列表模板

func (x *CacheSetNode) OnMsg(ctx types.RuleContext, msg types.RuleMsg) {
	_ = "STUB: not implemented"
	return
}

// Destroy 销毁组件
func (x *CacheSetNode) Destroy() {
	_ = "STUB: not implemented"

	// CacheDeleteNodeConfiguration 缓存删除节点配置
	return
}

type CacheDeleteNodeConfiguration struct {
	// Keys 删除的键列表
	Keys []LevelKey `json:"keys"`
}

// CacheDeleteNode removes data from cache storage at different levels.
// It supports exact key deletion and prefix-based batch deletion with wildcard patterns.
//
// CacheDeleteNode 从不同级别的缓存存储中删除数据。
// 支持精确键删除和基于前缀的通配符批量删除。
//
// Configuration:
// 配置说明：
//
//	{
//		"keys": [                        // Keys to delete  要删除的键列表
//			{
//				"level": "chain",        // Cache level: "chain" or "global"  缓存级别
//				"key": "session_${metadata.userId}"  // Exact key with variable substitution  精确键支持变量替换
//			},
//			{
//				"level": "global",
//				"key": "temp_*"          // Wildcard pattern for batch deletion  批量删除的通配符模式
//			}
//		]
//	}
//
// Deletion Patterns:
// 删除模式：
//
//   - Exact deletion: "user:123" removes specific key  精确删除：删除特定键
//   - Batch deletion: "session:*" removes all keys with prefix "session:"  批量删除：删除所有前缀为 "session:" 的键
//   - Variable substitution: "cache_${metadata.id}" uses runtime values  变量替换：使用运行时值
//
// Usage Example:
// 使用示例：
//
//	// Clean up user session and temporary data
//	// 清理用户会话和临时数据
//	{
//		"id": "cacheDelete",
//		"type": "cacheDelete",
//		"configuration": {
//			"keys": [
//				{"level": "chain", "key": "session_${metadata.userId}"},
//				{"level": "global", "key": "temp_${metadata.requestId}_*"}
//			]
//		}
//	}
type CacheDeleteNode struct {
	//节点配置
	Config CacheDeleteNodeConfiguration
	//keys模板
	keysTemplate []LevelKeyTemplate
}

func (x *CacheDeleteNode) Type() string { _ = "STUB: not implemented"; return "" }

func (x *CacheDeleteNode) New() types.Node { _ = "STUB: not implemented"; return *new(types.Node) }

// Init 初始化组件
func (x *CacheDeleteNode) Init(ruleConfig types.Config, configuration types.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

//初始化keys模板

func (x *CacheDeleteNode) OnMsg(ctx types.RuleContext, msg types.RuleMsg) {
	_ = "STUB: not implemented"
	return
}

//处理keys模板

func (x *CacheDeleteNode) handleDelete(ctx types.RuleContext, msg types.RuleMsg, keys []LevelKey) {
	_ = "STUB: not implemented"
	return
}

// Destroy 销毁组件
func (x *CacheDeleteNode) Destroy() { _ = "STUB: not implemented"; return }
