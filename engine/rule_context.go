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

package engine

import (
	"context"
	"sync"

	"github.com/rulego/rulego/api/types"
)

// Ensuring DefaultRuleContext implements types.RuleContext interface.
var _ types.RuleContext = (*DefaultRuleContext)(nil)

// GetEnv 获取环境变量和元数据
func (ctx *DefaultRuleContext) getEnv(msg types.RuleMsg, useMetadata bool, nodeIds ...string) map[string]interface{} {
	_ = "STUB: not implemented"
	// 预分配合适大小的map，减少扩容开销
	return nil
}

// 基础字段数量：id, ts, data, msgType, dataType, msg, metadata

// 估算metadata的键值对数量
// 常见metadata数量的估计值

// 设置基础字段

// 优化JSON数据处理

// 处理metadata - 使用零拷贝ForEach优化

// 使用零拷贝ForEach将metadata键值对添加到环境变量中

// continue iteration

// GetEnv 获取环境变量和元数据，支持跨节点取值
// msg: 当前消息
// useMetadata: 是否包含metadata
// 返回包含跨节点数据的上下文map，格式为 nodeId.msg.xx 和 nodeId.metadata.xx
func (ctx *DefaultRuleContext) GetEnv(msg types.RuleMsg, useMetadata bool) map[string]interface{} {
	_ = "STUB: not implemented"
	// 获取基础环境变量和metadata
	return nil
}

// 确定需要访问的节点ID列表

// 自动获取当前节点的依赖节点ID列表

// 为每个节点ID添加跨节点数据

// ContextObserver tracks the execution state of nodes in the rule chain.
type ContextObserver struct {
	// Map of executed nodes
	executedNodes sync.Map
	// Map of input messages for each node
	nodeInMsgList map[string][]types.WrapperMsg
	// Map of callbacks for node completion events
	nodeDoneEvent map[string]joinNodeCallback
	sync.RWMutex
}

// addInMsg adds an input message for a specific join node.
func (c *ContextObserver) addInMsg(joinNodeId, fromId string, msg types.RuleMsg, errStr string) bool {
	_ = "STUB: not implemented"
	return false
}

// getInMsgList retrieves the list of input messages for a specific join node.
func (c *ContextObserver) getInMsgList(joinNodeId string) []types.WrapperMsg {
	_ = "STUB: not implemented"
	return nil
}

// registerNodeDoneEvent registers a callback for when a join node completes.
func (c *ContextObserver) registerNodeDoneEvent(joinNodeId, lcaNodeId string, callback func([]types.WrapperMsg)) {
	_ = "STUB: not implemented"
	return
}

// checkNodesDone checks if all specified nodes have completed execution.
func (c *ContextObserver) checkNodesDone(nodeIds ...string) bool {
	_ = "STUB: not implemented"
	return false
}

// executedNode marks a node as executed and checks for any completed join nodes.
func (c *ContextObserver) executedNode(nodeId string) { _ = "STUB: not implemented"; return }

// checkAndTrigger checks for completed join nodes and triggers their callbacks.
func (c *ContextObserver) checkAndTrigger() { _ = "STUB: not implemented"; return }

// 获取消息列表并触发回调

// 直接执行回调，保持原有的同步行为

// joinNodeCallback represents a callback function for when a join node completes.
type joinNodeCallback struct {
	lcaNodeId  string //joinNodeId 节点最近共同祖先节点
	joinNodeId string
	callback   func([]types.WrapperMsg)
}

// DefaultRuleContext is the default context for message processing in the rule engine.
type DefaultRuleContext struct {
	// Context for sharing semaphores and data across different components.
	context context.Context
	// Configuration settings for the rule engine.
	config types.Config
	// Context of the root rule chain.
	ruleChainCtx *RuleChainCtx
	// Context of the previous node.
	from types.NodeCtx
	// Context of the current node.
	self types.NodeCtx
	// Indicates if this is the first node in the chain.
	isFirst bool
	// Goroutine pool for concurrent execution.
	pool types.Pool
	// Callback function for when the rule chain branch processing ends.
	onEnd types.OnEndFunc
	// Count of child nodes that have not yet completed execution.
	waitingCount int32
	// Parent rule context.
	parentRuleCtx *DefaultRuleContext
	// Event that triggers once when all child nodes have completed, executed only once.
	onAllNodeCompleted func()
	// Indicates if the onAllNodeCompleted function has been executed.
	onAllNodeCompletedDone int32
	// Pool for sub-rule chains.
	ruleChainPool types.RuleEnginePool
	// Indicates whether to skip executing child nodes, default is false.
	skipTellNext bool
	// List of aspects.
	aspects types.AspectList
	// List of around aspects.
	aroundAspects []types.AroundAspect
	// List of before aspects.
	beforeAspects []types.BeforeAspect
	// List of after aspects.
	afterAspects []types.AfterAspect
	// Runtime snapshot for debugging and logging.
	runSnapshot *RunSnapshot
	// Observer for join nodes - 延迟初始化
	observer *ContextObserver
	// first node relationType
	relationTypes []string
	// OUT msg
	out types.RuleMsg
	// IN or OUT err
	err        error
	chainCache types.Cache
	// nodeOutputCache 节点输出缓存，用于跨节点取值
	nodeOutputCache *NodeOutputCache
	//该链是否有结束节点
	hasEndNode bool
	// restoreNodeInfo 恢复执行节点信息
	restoreNodeInfo *RestoreNodeInfo
}

// RestoreNodeInfo 恢复执行节点信息
type RestoreNodeInfo struct {
	// NodeRequests 恢复执行节点请求列表
	NodeRequests []types.NodeRequest
}

func (ctx *DefaultRuleContext) GlobalCache() types.Cache {
	_ = "STUB: not implemented"
	return *new(types.Cache)
}

func (ctx *DefaultRuleContext) ChainCache() types.Cache {
	_ = "STUB: not implemented"
	return *

	// GetNodeOutputCache 获取节点输出缓存
	// GetNodeOutputCache returns the node output cache
	new(types.Cache)
}

func (ctx *DefaultRuleContext) GetNodeOutputCache() *NodeOutputCache {
	_ = "STUB: not implemented"
	return nil

	// NewRuleContext creates a new instance of the default rule engine message processing context.
}

func NewRuleContext(context context.Context, config types.Config, ruleChainCtx *RuleChainCtx, from types.NodeCtx, self types.NodeCtx, pool types.Pool, onEnd types.OnEndFunc, ruleChainPool types.RuleEnginePool) *DefaultRuleContext {
	_ = "STUB: not implemented"

	// Initialize aspects list.
	return nil
}

// If no aspects are defined, use built-in aspects.

// Get node-specific aspects.

// Return a new DefaultRuleContext populated with the provided parameters and aspects.

// RunSnapshot holds the state and logs for a rule chain execution.
type RunSnapshot struct {
	// Unique identifier for the message being processed.
	msgId string
	// Context of the rule chain being executed.
	chainCtx *RuleChainCtx
	// Timestamp marking the start of execution.
	startTs int64
	// Callback function for when the rule chain execution is completed.
	onRuleChainCompletedFunc func(ctx types.RuleContext, snapshot types.RuleChainRunSnapshot)
	// Callback function for when a node execution is completed.
	onNodeCompletedFunc func(ctx types.RuleContext, nodeRunLog types.RuleNodeRunLog)
	// Logs for each node's execution.
	logs map[string]*types.RuleNodeRunLog
	// Custom debug callback function.
	onDebugCustomFunc func(ruleChainId string, flowType string, nodeId string, msg types.RuleMsg, relationType string, err error)
	// Lock for synchronizing access to logs.
	lock sync.RWMutex
}

// NewRunSnapshot creates a new instance of RunSnapshot with the given parameters.
func NewRunSnapshot(msgId string, chainCtx *RuleChainCtx, startTs int64) *RunSnapshot {
	_ = "STUB: not implemented"
	return nil
}

// Initialize the logs map.

// needCollectRunSnapshot determines if there is a need to collect a snapshot of the rule chain execution.
func (r *RunSnapshot) needCollectRunSnapshot() bool { _ = "STUB: not implemented"; return false }

// collectRunSnapshot collects a snapshot of the rule node's execution state.
func (r *RunSnapshot) collectRunSnapshot(ctx types.RuleContext, flowType string, nodeId string, msg types.RuleMsg, relationType string, err error) {
	_ = "STUB: not implemented"
	return
}

// If the flow type is 'In', update the log with the incoming message and timestamp.

// If the flow type is 'Out', update the log with the outgoing message, relation type, and timestamp.

// If the flow type is 'Log', append the log item to the node's log items.

// onDebugCustom invokes the custom debug function with the provided parameters.
func (r *RunSnapshot) onDebugCustom(ruleChainId string, flowType string, nodeId string, msg types.RuleMsg, relationType string, err error) {
	_ = "STUB: not implemented"
	return
}

// createRuleChainRunLog creates a log of the entire rule chain's execution.
func (r *RunSnapshot) createRuleChainRunLog(endTs int64) types.RuleChainRunSnapshot {
	_ = "STUB: not implemented"
	return *new(types.RuleChainRunSnapshot)
}

// onRuleChainCompleted is called when the rule chain execution is completed.
func (r *RunSnapshot) onRuleChainCompleted(ctx types.RuleContext) {
	_ = "STUB: not implemented"
	return
}

// NewNextNodeRuleContext creates a new instance of RuleContext for the next node in the rule engine.
// 预定义常用关系类型的单例 slice，避免重复分配
// Pre-defined singleton slices for common relation types to avoid repeated allocations
var (
	successRelationTypes = []string{types.Success}
	failureRelationTypes = []string{types.Failure}
	trueRelationTypes    = []string{types.True}
	falseRelationTypes   = []string{types.False}
)

// NewNextNodeRuleContext 创建下一个节点的规则上下文
// NewNextNodeRuleContext creates a rule context for the next node
func (ctx *DefaultRuleContext) NewNextNodeRuleContext(nextNode types.NodeCtx) *DefaultRuleContext {
	_ = "STUB: not implemented"
	// Create a new context directly instead of using object pool to avoid data races
	// 但是复用不可变的共享状态以减少内存开销
	return nil
}

// 共享配置，不可变
// 共享规则链上下文

// 共享协程池

// 共享规则链池
// 直接复用context，避免调用GetContext()

// 共享切面列表，它们在运行时不会改变

// 共享运行时状态

// 子context共享observer
// 共享observer实例

// 共享缓存
// 共享节点输出缓存

func (ctx *DefaultRuleContext) TellSuccess(msg types.RuleMsg) { _ = "STUB: not implemented"; return }

func (ctx *DefaultRuleContext) TellFailure(msg types.RuleMsg, err error) {
	_ = "STUB: not implemented"
	return
}

func (ctx *DefaultRuleContext) TellNext(msg types.RuleMsg, relationTypes ...string) {
	_ = "STUB: not implemented"
	return
}

func (ctx *DefaultRuleContext) TellSelf(msg types.RuleMsg, delayMs int64) {
	_ = "STUB: not implemented"
	return
}

func (ctx *DefaultRuleContext) TellNextOrElse(msg types.RuleMsg, defaultRelationType string, relationTypes ...string) {
	_ = "STUB: not implemented"
	return
}

func (ctx *DefaultRuleContext) TellCollect(msg types.RuleMsg, callback func(msgList []types.WrapperMsg)) bool {
	_ = "STUB: not implemented"
	return false
}

//通知当前节点至共同祖先这分支链已经执行完。

//通知当前节点至共同祖先这分支链已经执行完。

// 获取LCA节点

// 注意：不再调用 executedNode(fromId)
// LCA节点应该通过 childDoneWithoutCallback 流程在 waitingCount 归零时被标记为完成
// 而不是在这里立即标记。否则当 fork 节点直接连接到 join 节点时，
// fork 会在所有子节点完成前就被标记为完成，导致 join 过早触发回调。

func (ctx *DefaultRuleContext) NewMsg(msgType string, metaData *types.Metadata, data string) types.RuleMsg {
	_ = "STUB: not implemented"
	return *new(types.RuleMsg)
}

func (ctx *DefaultRuleContext) GetSelfId() string { _ = "STUB: not implemented"; return "" }

func (ctx *DefaultRuleContext) Self() types.NodeCtx {
	_ = "STUB: not implemented"
	return *new(types.NodeCtx)
}

func (ctx *DefaultRuleContext) From() types.NodeCtx {
	_ = "STUB: not implemented"
	return *new(types.NodeCtx)
}

func (ctx *DefaultRuleContext) RuleChain() types.NodeCtx {
	_ = "STUB: not implemented"
	return *new(types.NodeCtx)
}

func (ctx *DefaultRuleContext) Config() types.Config {
	_ = "STUB: not implemented"
	return *new(types.Config)
}

func (ctx *DefaultRuleContext) SetEndFunc(onEndFunc types.OnEndFunc) types.RuleContext {
	_ = "STUB: not implemented"
	return *new(types.RuleContext)
}

func (ctx *DefaultRuleContext) GetEndFunc() types.OnEndFunc {
	_ = "STUB: not implemented"
	return *new(types.OnEndFunc)
}

func (ctx *DefaultRuleContext) SetContext(c context.Context) types.RuleContext {
	_ = "STUB: not implemented"
	return *new(types.RuleContext)
}

func (ctx *DefaultRuleContext) GetContext() context.Context {
	_ = "STUB: not implemented"
	return *

	// Deprecated: Use Flow SubmitTask instead.
	new(context.Context)
}

func (ctx *DefaultRuleContext) SubmitTack(task func()) { _ = "STUB: not implemented"; return }

func (ctx *DefaultRuleContext) SubmitTask(task func()) { _ = "STUB: not implemented"; return }

// 在提交任务前捕获需要的值，避免并发访问

// 如果工作池提交失败，回退到直接创建goroutine
// 这确保任务不会丢失，避免计数器不匹配导致的死锁

// TellFlow 执行子规则链，ruleChainId 规则链ID
// onEndFunc 子规则链链分支执行完的回调，并返回该链执行结果，如果同时触发多个分支链，则会调用多次
// onAllNodeCompleted 所以节点执行完触发，无结果返回
// 如果找不到规则链，并把消息通过`Failure`关系发送到下一个节点
func (ctx *DefaultRuleContext) TellFlow(ruleChainId string, msg types.RuleMsg, opts ...types.RuleContextOption) {
	_ = "STUB: not implemented"
	return
}

// TellNode 从指定节点开始执行，如果 skipTellNext=true 则只执行当前节点，不通知下一个节点。
// onEnd 查看获得最终执行结果
// onAllNodeCompleted 所以节点执行完触发，无结果返回
func (ctx *DefaultRuleContext) TellNode(chanCtx context.Context, nodeId string, msg types.RuleMsg, skipTellNext bool, onEnd types.OnEndFunc, onAllNodeCompleted func()) {
	_ = "STUB: not implemented"
	return
}

//Whether to only execute the current node

//如果只执行一个节点，则肯定没有结束节点（它本身就是结束节点）

func (ctx *DefaultRuleContext) TellChainNode(chanCtx context.Context, ruleChainId, nodeId string, msg types.RuleMsg, skipTellNext bool, onEnd types.OnEndFunc, onAllNodeCompleted func()) {
	_ = "STUB: not implemented"
	// Tell current chain node
	return
}

// Tell other chain node

func (ctx *DefaultRuleContext) tellOtherChainNode(chanCtx context.Context, ruleChainId, nodeId string, msg types.RuleMsg, skipTellNext bool, onEnd types.OnEndFunc, onAllNodeCompleted func()) {
	_ = "STUB: not implemented"
	return
}

// SetRuleChainPool 设置子规则链池
func (ctx *DefaultRuleContext) SetRuleChainPool(ruleChainPool types.RuleEnginePool) {
	_ = "STUB: not implemented"
	return
}

// GetRuleChainPool 获取子规则链池
func (ctx *DefaultRuleContext) GetRuleChainPool() types.RuleEnginePool {
	_ = "STUB: not implemented"
	return *new(types.RuleEnginePool)
}

// SetOnAllNodeCompleted 设置所有节点执行完回调
func (ctx *DefaultRuleContext) SetOnAllNodeCompleted(onAllNodeCompleted func()) {
	_ = "STUB: not implemented"
	return
}

func (ctx *DefaultRuleContext) HasEndNode() bool { _ = "STUB: not implemented"; return false }

// DoOnEnd  结束规则链分支执行，触发 OnEnd 回调函数
func (ctx *DefaultRuleContext) DoOnEnd(msg types.RuleMsg, err error, relationType string) {
	_ = "STUB: not implemented"
	return
}

// 拷贝msg

// 确保Metadata不为nil，避免空指针异常

// 如果配置了结束节点，只有结束节点才能触发回调；如果没有配置结束节点，所有节点都可以触发

// 是否触发回调

//全局回调
//通过`Config.OnEnd`设置

// types.withOnEnd 设置的回调

// 是否触发回调

//全局回调
//通过`Config.OnEnd`设置

// types.withOnEnd 设置的回调

// 执行AfterAop

func (ctx *DefaultRuleContext) SetCallbackFunc(functionName string, f interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ctx *DefaultRuleContext) GetCallbackFunc(functionName string) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *DefaultRuleContext) OnDebug(ruleChainId string, flowType string, nodeId string, msg types.RuleMsg, relationType string, err error) {
	_ = "STUB: not implemented"
	// 在方法开始时就缓存runSnapshot引用，避免并发竞态条件
	return
}

// 智能拷贝优化：只有在真正需要时才拷贝消息

// 只有在真正需要拷贝时才创建副本

// 在提交异步任务前捕获需要的值，避免并发访问

//异步记录日志

//记录快照

// SetExecuteNodes 设置执行节点
// 可以设置单个或多个节点，用于恢复执行或指定起始节点
func (ctx *DefaultRuleContext) SetExecuteNodes(nodes ...types.NodeRequest) {
	_ = "STUB: not implemented"
	return

	// 检查是否是搜索模式或者包含关系类型
	// 如果 RelationTypes 不为 nil，则视为查找子节点模式
	// If RelationTypes is not nil, it is considered as finding child nodes mode.
}

// 执行当前节点模式

// 清空 restoreNodeInfo，确保使用 TellNext 路径

// GetRelationTypes 获取当前输入节点执行关系
func (ctx *DefaultRuleContext) GetRelationTypes() []string { _ = "STUB: not implemented"; return nil }

func (ctx *DefaultRuleContext) GetOut() types.RuleMsg {
	_ = "STUB: not implemented"
	return *new(types.RuleMsg)
}

func (ctx *DefaultRuleContext) GetErr() error {
	_ = "STUB: not implemented"

	// IsDebugMode 是否调试模式，优先使用规则链指定的调试模式
	return nil
}

func (ctx *DefaultRuleContext) IsDebugMode() bool { _ = "STUB: not implemented"; return false }

// 增加一个待执行子节点
func (ctx *DefaultRuleContext) childReady(msg types.RuleMsg, relationType string) {
	_ = "STUB: not implemented"
	return
}

// 减少一个待执行子节点
// 如果返回数量0，表示该分支链条已经都执行完成，递归父节点，直到所有节点都处理完，则触发onAllNodeCompleted事件。
func (ctx *DefaultRuleContext) childDone() { _ = "STUB: not implemented"; return }

// 在进行任何异步操作前捕获需要的值，避免并发问题

//该节点已经执行完成，通知父节点

// 只有在observer存在时才记录节点执行完成（通常是join节点场景）

//记录当前节点执行完成

//完成回调

// childDoneWithoutCallback 通知父节点有一个子节点执行完成，但不触发 onAllNodeCompleted 回调事件
//
// 与 childDone() 方法的区别：
// 1. childDone(): 子节点执行完成时会触发 onAllNodeCompleted 回调，适用于正常节点执行完成场景
// 2. childDoneWithoutCallback(): 子节点执行完成时不触发回调，专用于聚合多条分支链数据场景
//
// 使用场景：
// - 配合 TellCollect 方法使用，查询该节点的父节点共同祖先是否所有分支到当前聚合节点都已经执行完成
// - 用于多分支汇聚场景的状态跟踪，避免过早触发完成事件
func (ctx *DefaultRuleContext) childDoneWithoutCallback() { _ = "STUB: not implemented"; return }

//if atomic.CompareAndSwapInt32(&ctx.onAllNodeCompletedDone, 0, 1) {

// 在进行任何异步操作前捕获需要的值，避免并发问题

//有子节点执行完成，通知父节点

// 只有在observer存在时才记录节点执行完成（通常是join节点场景）

//记录当前节点执行完成

// getNextNodes 获取当前节点指定关系的子节点
func (ctx *DefaultRuleContext) getNextNodes(relationType string) ([]types.NodeCtx, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// tellSelf 执行自身节点
func (ctx *DefaultRuleContext) tellSelf(msg types.RuleMsg, err error, relationTypes ...string) {
	_ = "STUB: not implemented"
	return
}

// 异步执行需要拷贝确保线程安全
// 注意：不能简单根据节点类型优化，因为其他并发分支可能修改消息

// tellNext 通知执行子节点，如果是当前第一个节点则执行当前节点
func (ctx *DefaultRuleContext) tell(msg types.RuleMsg, err error, relationTypes ...string) {
	_ = "STUB: not implemented"
	return
}

// tellNext 通知执行子节点，如果是当前第一个节点则执行当前节点
// 如果找不到relationTypes对应的节点，而且defaultRelationType非默认值，则通过defaultRelationType查找节点
func (ctx *DefaultRuleContext) tellOrElse(msg types.RuleMsg, err error, defaultRelationType string, relationTypes ...string) {
	_ = "STUB: not implemented"
	return
}

//找不到子节点，则执行结束回调

//执行After aop

//根据relationType查找子节点列表

//根据默认关系查找节点

// 只有多个子节点或者并行多个关系时才需要拷贝

//增加一个待执行的子节点

//除1个节点和并行多个关系外的其他节点创建拷贝

//唯一节点可以直接使用原消息

//通知执行子节点

//为了保证流块的顺序

//调用DoOnEnd 会调用 childDone()对waitingCount会减1，所以childReady和childDone成对出现

//找不到子节点，则执行结束回调

// 执行环绕aop
// 返回值true: 继续执行下一个节点，否则不执行
func (ctx *DefaultRuleContext) executeAroundAop(msg types.RuleMsg, relationType string) bool {
	_ = "STUB: not implemented"
	// before aop
	return false
}

//是否已经执行了tellNext逻辑
//如果 AroundAspect 已经执行了tellNext逻辑，则引擎不再执行tellNext逻辑

// 执行After aop
func (ctx *DefaultRuleContext) executeAfterAop(msg types.RuleMsg, err error, relationType string) types.RuleMsg {
	_ = "STUB: not implemented"
	// after aop
	return *new(types.RuleMsg)
}

// 执行下一个节点
func (ctx *DefaultRuleContext) tellNext(msg types.RuleMsg, nextNode types.NodeCtx, relationType string) {
	_ = "STUB: not implemented"

	//捕捉异常
	return
}

//执行After aop

// 统一检查上下文是否已取消（优雅停机）
// Unified check for context cancellation (graceful shutdown)

// 上下文已取消，停止处理并通知失败
// Context cancelled, stop processing and notify failure
// 使用DoOnEnd确保正确触发结束回调和活跃消息计数减少
// DoOnEnd内部会调用childDone()，所以这里不需要再次调用

// 上下文正常，继续处理
// Context is normal, continue processing

// 在执行下一个节点之前，存储当前节点的输出到缓存
// Store current node output to cache before executing next node

//环绕aop

// 如果AroundAspect阻止了执行，需要调用childDone来平衡之前的childReady

// AroundAop 已经执行节点OnMsg逻辑，不在执行下面的逻辑

// setRelationType 优化关系类型的赋值，使用预定义的单例或复用已分配的 slice
// setRelationType optimizes relation type assignment using predefined singletons or reusing allocated slices
func (ctx *DefaultRuleContext) setRelationType(nextCtx *DefaultRuleContext, relationType string) {
	_ = "STUB: not implemented"
	// 对于常用的关系类型，使用预定义的单例 slice 以避免内存分配
	// For common relation types, use predefined singleton slices to avoid memory allocation
	return
}

// 对于自定义关系类型，复用已分配的 slice
// For custom relation types, reuse the pre-allocated slice

// GetNodeRuleMsg retrieves the complete RuleMsg of a specific executed node by nodeId
// IMPORTANT: Node dependency must be established beforehand to successfully retrieve data
//
// Dependency establishment methods:
// 1. Using FetchNodeOutputNode component (automatic)
// 2. Manually calling chainCtx.AddNodeDependency(currentNodeId, targetNodeId)
// 3. Node configuration contains references to other nodes. e.g. ${nodeId.msg.xx} (auto-detected)
func (ctx *DefaultRuleContext) GetNodeRuleMsg(nodeId string) (types.RuleMsg, bool) {
	_ = "STUB: not implemented"
	// 从节点输出缓存中获取目标节点的RuleMsg
	// 只有建立了依赖关系的节点输出才会被缓存
	// Retrieve target node's RuleMsg from node output cache
	// Only outputs from nodes with established dependencies are cached
	return *new(types.RuleMsg), false
}

// Target node output not found, possible reasons:
// 1. Node not yet executed
// 2. Dependency not established
// 3. Node execution failed

// StoreNodeOutput 存储节点输出到缓存中，用于跨节点取值
// 只有在以下情况下才会进行缓存：
// 1. 配置中启用了节点输出缓存 (EnableNodeOutputCache = true)
// 2. 或者已检测到跨节点取值用法 (通过EnableCrossNodeAccess()启用)
// 参数:
//   - nodeId: 节点ID
//   - msg: 规则消息
func (ctx *DefaultRuleContext) StoreNodeOutput(nodeId string, msg types.RuleMsg) {
	_ = "STUB: not implemented"
	return
}
