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
	"github.com/rulego/rulego/components/filter"
)

// init 注册 InclusiveNode 组件
// Init registers the InclusiveNode component with the default registry.
func init() {
	filter.Registry.Add(&InclusiveNode{})
}

// InclusiveNode 基于表达式评估，向所有匹配的分支同时路由的过滤组件
// InclusiveNode embeds SwitchNode to reuse initialization and configuration, and overrides routing behavior.
type InclusiveNode struct {
	SwitchNode
}

// Type 返回组件类型标识
// Type returns the component type identifier.
func (x *InclusiveNode) Type() string {
	_ = "STUB: not implemented"

	// New 创建组件实例
	// New creates a new component instance with default demo cases.
	return ""
}

func (x *InclusiveNode) New() types.Node { _ = "STUB: not implemented"; return *new(types.Node) }

// Init 初始化组件，编译所有 case 表达式
// Init initializes the component by compiling all case expressions into programs.
// Init 复用 SwitchNode 的初始化逻辑
// Init reuses SwitchNode.Init for compiling case expressions.
func (x *InclusiveNode) Init(ruleConfig types.Config, configuration types.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

// OnMsg 处理消息：评估所有 case，将消息路由到所有匹配的关系；若无匹配则路由到默认关系
// OnMsg processes the incoming message by evaluating all case expressions.
// It routes to each relation whose expression evaluates to true. If none matches, routes to Default.
func (x *InclusiveNode) OnMsg(ctx types.RuleContext, msg types.RuleMsg) {
	_ = "STUB: not implemented"
	return
}

// Destroy 清理资源
// Destroy cleans up resources.
func (x *InclusiveNode) Destroy() { _ = "STUB: not implemented"; return }
