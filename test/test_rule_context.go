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

package test

import (
	"context"
	"sync"

	"github.com/rulego/rulego/api/types"
)

var _ types.RuleContext = (*NodeTestRuleContext)(nil)

// NodeTestRuleContext
// 只为测试单节点，临时创建的上下文
// 无法把多个节点组成链式
// callback 回调处理结果
type NodeTestRuleContext struct {
	context  context.Context
	config   types.Config
	callback func(msg types.RuleMsg, relationType string, err error)
	self     types.Node
	selfId   string
	//所有子节点处理完成事件，只执行一次
	onAllNodeCompleted func()
	onEndFunc          types.OnEndFunc
	childrenNodes      sync.Map
	out                types.RuleMsg
	globalCache        types.Cache
	chainCache         types.Cache
	mutex              sync.RWMutex // Add mutex for thread safety
}

func (ctx *NodeTestRuleContext) GlobalCache() types.Cache {
	_ = "STUB: not implemented"
	return *new(types.Cache)
}

func (ctx *NodeTestRuleContext) ChainCache() types.Cache {
	_ = "STUB: not implemented"
	return *new(types.Cache)
}

func NewRuleContext(config types.Config, callback func(msg types.RuleMsg, relationType string, err error)) types.RuleContext {
	_ = "STUB: not implemented"
	return *new(types.RuleContext)
}

func NewRuleContextFull(config types.Config, self types.Node, childrenNodes map[string]types.Node, callback func(msg types.RuleMsg, relationType string, err error)) types.RuleContext {
	_ = "STUB: not implemented"
	return *new(types.RuleContext)
}

func (ctx *NodeTestRuleContext) TellSuccess(msg types.RuleMsg) { _ = "STUB: not implemented"; return }

func (ctx *NodeTestRuleContext) TellFailure(msg types.RuleMsg, err error) {
	_ = "STUB: not implemented"
	return
}

func (ctx *NodeTestRuleContext) TellNext(msg types.RuleMsg, relationTypes ...string) {
	_ = "STUB: not implemented"
	return
}

func (ctx *NodeTestRuleContext) TellSelf(msg types.RuleMsg, delayMs int64) {
	_ = "STUB: not implemented"
	return
}

func (ctx *NodeTestRuleContext) TellNextOrElse(msg types.RuleMsg, defaultRelationType string, relationTypes ...string) {
	_ = "STUB: not implemented"
	return
}

func (ctx *NodeTestRuleContext) NewMsg(msgType string, metaData *types.Metadata, data string) types.RuleMsg {
	_ = "STUB: not implemented"
	return *new(types.RuleMsg)
}

func (ctx *NodeTestRuleContext) GetSelfId() string { _ = "STUB: not implemented"; return "" }

func (ctx *NodeTestRuleContext) Self() types.NodeCtx {
	_ = "STUB: not implemented"
	return *new(types.NodeCtx)
}

func (ctx *NodeTestRuleContext) From() types.NodeCtx {
	_ = "STUB: not implemented"
	return *new(types.NodeCtx)
}

func (ctx *NodeTestRuleContext) RuleChain() types.NodeCtx {
	_ = "STUB: not implemented"
	return *new(types.NodeCtx)
}

func (ctx *NodeTestRuleContext) Config() types.Config {
	_ = "STUB: not implemented"
	return *new(types.Config)
}

func (ctx *NodeTestRuleContext) SubmitTack(task func()) { _ = "STUB: not implemented"; return }

func (ctx *NodeTestRuleContext) SubmitTask(task func()) { _ = "STUB: not implemented"; return }

func (ctx *NodeTestRuleContext) SetEndFunc(onEndFunc types.OnEndFunc) types.RuleContext {
	_ = "STUB: not implemented"
	return *new(types.RuleContext)
}

func (ctx *NodeTestRuleContext) GetEndFunc() types.OnEndFunc {
	_ = "STUB: not implemented"
	return *new(types.OnEndFunc)
}

func (ctx *NodeTestRuleContext) SetContext(c context.Context) types.RuleContext {
	_ = "STUB: not implemented"
	return *new(types.RuleContext)
}

func (ctx *NodeTestRuleContext) GetContext() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (ctx *NodeTestRuleContext) TellFlow(chainId string, msg types.RuleMsg, opts ...types.RuleContextOption) {
	_ = "STUB: not implemented"
	return
}

// TellNode 独立执行某个节点，通过callback获取节点执行情况，用于节点分组类节点控制执行某个节点
func (ctx *NodeTestRuleContext) TellNode(context context.Context, nodeId string, msg types.RuleMsg, skipTellNext bool, callback types.OnEndFunc, onAllNodeCompleted func()) {
	_ = "STUB: not implemented"
	return
}

// 线程安全地设置 selfId

// TellChainNode 独立执行某个节点，通过callback获取节点执行情况，用于节点分组类节点控制执行某个节点
func (ctx *NodeTestRuleContext) TellChainNode(context context.Context, chainId string, nodeId string, msg types.RuleMsg, skipTellNext bool, callback types.OnEndFunc, onAllNodeCompleted func()) {
	_ = "STUB: not implemented"
	return
}

// SetOnAllNodeCompleted 设置所有节点执行完回调
func (ctx *NodeTestRuleContext) SetOnAllNodeCompleted(onAllNodeCompleted func()) {
	_ = "STUB: not implemented"
	return
}

func (ctx *NodeTestRuleContext) DoOnEnd(msg types.RuleMsg, err error, relationType string) {
	_ = "STUB: not implemented"

	// SetCallbackFunc 设置回调函数
	return
}

func (ctx *NodeTestRuleContext) SetCallbackFunc(functionName string, f interface{}) {
	_ = "STUB: not implemented"

	// GetCallbackFunc 获取回调函数
	return
}

func (ctx *NodeTestRuleContext) GetCallbackFunc(functionName string) interface{} {
	_ = "STUB: not implemented"

	// OnDebug 调用配置的OnDebug回调函数
	return nil
}

func (ctx *NodeTestRuleContext) OnDebug(ruleChainId string, flowType string, nodeId string, msg types.RuleMsg, relationType string, err error) {
	_ = "STUB: not implemented"
	return
}

func (ctx *NodeTestRuleContext) SetExecuteNodes(nodes ...types.NodeRequest) {
	_ = "STUB: not implemented"
	return
}

func (ctx *NodeTestRuleContext) TellCollect(msg types.RuleMsg, callback func(msgList []types.WrapperMsg)) bool {
	_ = "STUB: not implemented"
	return false
}

func (ctx *NodeTestRuleContext) GetOut() types.RuleMsg {
	_ = "STUB: not implemented"
	return *new(types.RuleMsg)
}

func (ctx *NodeTestRuleContext) GetRelationTypes() []string {
	_ = "STUB: not implemented"

	// setOut safely sets the out field
	return nil
}

func (ctx *NodeTestRuleContext) setOut(msg types.RuleMsg) { _ = "STUB: not implemented"; return }

func (ctx *NodeTestRuleContext) GetErr() error { _ = "STUB: not implemented"; return nil }

func (ctx *NodeTestRuleContext) TellStream(msg types.RuleMsg) { _ = "STUB: not implemented"; return }

// GetEnv 获取环境变量和元数据
func (ctx *NodeTestRuleContext) GetEnv(msg types.RuleMsg, useMetadata bool) map[string]interface{} {
	_ = "STUB: not implemented"
	// 创建环境变量map
	return nil
}

// 设置基础环境变量

// 使用 GetJsonData() 避免重复JSON解析

// 解析失败，使用原始数据

// 如果不是 JSON 类型，直接使用原始数据

// 优化 metadata 处理

// 遍历metadata，将键值对添加到环境变量中 - use zero-copy ForEach

// continue iteration

// GetNodeRuleMsg 获取节点的完整消息信息（测试上下文中暂不支持跨节点取值）
// GetNodeRuleMsg retrieves the complete RuleMsg of a node (not supported in test context)
func (ctx *NodeTestRuleContext) GetNodeRuleMsg(nodeId string) (types.RuleMsg, bool) {
	_ = "STUB: not implemented"
	return *new(types.RuleMsg), false
}

// ExtendedTestRuleContext 扩展的测试上下文，支持结果收集和节点处理器设置
// 可以替代 SimpleTestContext 和 MockRuleContext
type ExtendedTestRuleContext struct {
	*NodeTestRuleContext
	nodeHandlers map[string]func(msg types.RuleMsg) (string, error)
	results      []string
	resultsChan  chan TestResult
	handlerMutex sync.RWMutex
}

// TestResult 测试结果结构
type TestResult struct {
	RelationType string
	Err          error
}

// NewExtendedTestRuleContext 创建扩展的测试上下文
// 用于替代 SimpleTestContext 和 MockRuleContext
func NewExtendedTestRuleContext(config types.Config, callback func(msg types.RuleMsg, relationType string, err error)) *ExtendedTestRuleContext {
	_ = "STUB: not implemented"
	return nil
}

// NewExtendedTestRuleContextWithChannel 创建带结果通道的扩展测试上下文
// 主要用于替代 SimpleTestContext
func NewExtendedTestRuleContextWithChannel() *ExtendedTestRuleContext {
	_ = "STUB: not implemented"
	return nil
}

// SetNodeHandler 设置节点处理器，用于模拟节点行为
// 替代 MockRuleContext 的 SetNodeHandler 方法
func (ctx *ExtendedTestRuleContext) SetNodeHandler(nodeId string, handler func(msg types.RuleMsg) (string, error)) {
	_ = "STUB: not implemented"
	return
}

// GetResults 获取收集的结果
// 替代 MockRuleContext 的 GetResults 方法
func (ctx *ExtendedTestRuleContext) GetResults() []string { _ = "STUB: not implemented"; return nil }

// GetResultsChannel 获取结果通道
// 用于替代 SimpleTestContext 的 results 通道
func (ctx *ExtendedTestRuleContext) GetResultsChannel() <-chan TestResult {
	_ = "STUB: not implemented"
	return nil

	// TellNode 重写 TellNode 方法以支持节点处理器
}

func (ctx *ExtendedTestRuleContext) TellNode(context context.Context, nodeId string, msg types.RuleMsg, skipTellNext bool, callback types.OnEndFunc, onAllNodeCompleted func()) {
	_ = "STUB: not implemented"
	return
}

// 使用自定义处理器（模拟节点行为）

// 使用原有的 TellNode 逻辑

// TellNext 重写以支持结果收集
func (ctx *ExtendedTestRuleContext) TellNext(msg types.RuleMsg, relationTypes ...string) {
	_ = "STUB: not implemented"
	// 调用原有逻辑
	return
}

// 收集结果

// 发送到结果通道

// TellSuccess 重写以支持结果收集
func (ctx *ExtendedTestRuleContext) TellSuccess(msg types.RuleMsg) {
	_ = "STUB: not implemented"
	// 调用原有逻辑
	return
}

// 收集结果

// 发送到结果通道

// TellFailure 重写以支持结果收集
func (ctx *ExtendedTestRuleContext) TellFailure(msg types.RuleMsg, err error) {
	_ = "STUB: not implemented"
	// 调用原有逻辑
	return
}

// 收集结果

// 发送到结果通道
