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

// Package base provides foundational components and utilities for the RuleGo rule engine.
package base

import (
	"errors"
	"sync"

	"github.com/rulego/rulego/api/types"
)

var (
	ErrNodePoolNil   = errors.New("node pool is nil")
	ErrClientNotInit = errors.New("client not init")
)

var NodeUtils = &nodeUtils{}

type nodeUtils struct {
}

func (n *nodeUtils) GetChainCtx(configuration types.Configuration) types.ChainCtx {
	_ = "STUB: not implemented"
	return *new(types.ChainCtx)
}

func (n *nodeUtils) GetSelfDefinition(configuration types.Configuration) types.RuleNode {
	_ = "STUB: not implemented"
	return *new(types.RuleNode)
}

func (n *nodeUtils) GetVars(configuration types.Configuration) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (n *nodeUtils) GetEvn(ctx types.RuleContext, msg types.RuleMsg) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (n *nodeUtils) GetEvnAndMetadata(ctx types.RuleContext, msg types.RuleMsg) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (n *nodeUtils) IsNodePool(config types.Config, server string) bool {
	_ = "STUB: not implemented"
	return false
}

func (n *nodeUtils) GetInstanceId(config types.Config, server string) string {
	_ = "STUB: not implemented"
	return ""
}

//截取资源ID

func (n *nodeUtils) IsInitNetResource(_ types.Config, configuration types.Configuration) bool {
	_ = "STUB: not implemented"
	return false
}

func (n *nodeUtils) getEvnAndMetadata(ctx types.RuleContext, msg types.RuleMsg, useMetadata bool) map[string]interface{} {
	_ = "STUB: not implemented"
	// 直接调用ctx的GetEvnAndMetadata方法
	return nil
}

// GetDataByType 准备传递给JavaScript脚本的数据
// 根据消息的数据类型进行不同的处理：
// - JSON类型：解析为map以便JavaScript处理
// - BINARY类型：转换为字节数组，JavaScript将其视为Uint8Array
// - 其他类型：使用原始字符串数据
func (n *nodeUtils) GetDataByType(msg types.RuleMsg, readOnly bool) interface{} {
	_ = "STUB: not implemented"
	return nil

	// 根据数据类型进行不同的处理
}

// JSON类型：js会修改数据，所以这里需要重新解析

// 二进制类型：创建字节数组副本以避免并发修改问题，JavaScript会将其视为Uint8Array

// 创建副本以确保并发安全

// 其他类型：使用原始字符串数据

// TrimStrings 去除配置中所有字符串值的前后空格
// 遍历 Configuration 中的所有值，如果是字符串类型则去除前后空格
func (n *nodeUtils) TrimStrings(config types.Configuration) { _ = "STUB: not implemented"; return }

// SharedNode 共享资源组件，通过 Get 获取共享实例，多个节点可以在共享池中获取相同的实例
// 例如：mqtt 客户端、数据库客户端，也可以http server以及是可复用的节点。
type SharedNode[T any] struct {
	//节点类型
	NodeType string
	//配置
	RuleConfig types.Config
	//资源ID
	InstanceId string
	//初始化实例资源函数
	InitInstanceFunc func() (T, error)
	//清理资源的回调函数
	CloseFunc func(T) error
	////初始化资源资源，防止并发初始化
	//lock int32
	//是否从资源池获取
	isFromPool bool
	Locker     sync.RWMutex

	// 本地客户端缓存（新API使用）
	localClient       T
	clientInitialized bool
}

// Init 初始化，如果 resourcePath 为 ref:// 开头，则从网络资源池获取，否则调用 initInstanceFunc 初始化
// initNow=true，会在立刻初始化，否则在 GetInstance() 时候初始化
func (x *SharedNode[T]) Init(ruleConfig types.Config, nodeType, resourcePath string, initNow bool, initInstanceFunc func() (T, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// InitWithClose 初始化，支持自定义清理函数
func (x *SharedNode[T]) InitWithClose(ruleConfig types.Config, nodeType, resourcePath string, initNow bool, initInstanceFunc func() (T, error), closeFunc func(T) error) error {
	_ = "STUB: not implemented"
	return nil
}

//非资源池方式，初始化

// 初始化成功，缓存客户端

// IsInit 是否初始化过
func (x *SharedNode[T]) IsInit() bool { _ = "STUB: not implemented"; return false }

// GetInstance 获取共享实例
func (x *SharedNode[T]) GetInstance() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// Get 获取共享实例，并返回具体类型
		// Deprecated: 建议使用 GetSafely() 方法，该方法提供更好的并发性能和资源管理。
		// 使用 GetSafely() 时需要配合 InitWithClose() 和 Close() 方法进行完整的资源管理。
		//func (x *SharedNode[T]) Get() (T, error) {
		//	if x.InstanceId != "" {
		//		//从网络资源池获取
		//		if x.RuleConfig.NodePool == nil {
		//			return zeroValue[T](), ErrNodePoolNil
		//		}
		//		if p, err := x.RuleConfig.NodePool.GetInstance(x.InstanceId); err == nil {
		//			return p.(T), nil
		//		} else {
		//			return zeroValue[T](), err
		//		}
		//	} else if x.InitInstanceFunc != nil {
		//		//根据当前组件配置初始化一个客户端
		//		return x.InitInstanceFunc()
		//	} else {
		//		return zeroValue[T](), ErrClientNotInit
		//	}
		//}
		nil
}

// GetSafely 安全获取共享实例，如果没有实例则初始化一个
// 推荐新组件使用此方法进行资源管理。
//
// 使用说明：
// 1. 初始化时使用 InitWithClose() 方法并提供清理函数
// 2. 获取实例时使用 GetSafely() 方法
// 3. 组件销毁时调用 Close() 方法清理资源
func (x *SharedNode[T]) GetSafely() (T, error) {
	_ = "STUB: not implemented"
	return *

	// 从网络资源池获取
	new(T), nil
}

// 首先使用读锁检查客户端是否已存在

// 客户端不存在，使用写锁进行创建

// 双重检查：可能在等待写锁期间其他goroutine已经创建了客户端

// 初始化客户端

// 初始化失败，如果返回了部分初始化的客户端，尝试清理

// 初始化成功，缓存客户端

// isZeroValue 检查值是否为零值
// 使用反射来安全地比较值，避免在不可比较类型上出现运行时恐慌
func isZeroValue[T any](v T) bool {
	_ = "STUB: not implemented"
	// 使用反射来安全地检查零值
	return false
}

// Close 清理本地缓存的客户端资源
// 与 GetSafely() 和 InitWithClose() 配合使用，提供完整的资源生命周期管理
// 注意：此方法不会影响从资源池获取的客户端
func (x *SharedNode[T]) Close() error {
	_ = "STUB: not implemented"
	// 只清理本地缓存的客户端，不影响资源池中的客户端
	return nil
}

// 资源池模式，不需要清理本地客户端

// 使用用户提供的清理函数或默认的Close方法

// 尝试调用客户端的Close方法（如果有的话）

// 重置本地客户端状态

// IsFromPool 是否从资源池获取
func (x *SharedNode[T]) IsFromPool() bool { _ = "STUB: not implemented"; return false }

func (x *SharedNode[T]) Initialized() bool { _ = "STUB: not implemented"; return false }

// zeroValue 函数用于返回 T 类型的零值
func zeroValue[T any]() T { _ = "STUB: not implemented"; return *new(T) }
