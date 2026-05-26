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

package service

import (
	"context"
	"examples/server/config"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/rulego/rulego/api/types"
)

// McpService 自定义组件服务
type McpService struct {
	username          string
	config            config.Config
	ruleConfig        types.Config
	Pool              types.RuleEnginePool
	componentService  *ComponentService
	ruleEngineService *RuleEngineService
	mcpServer         *server.MCPServer
	sseServer         *server.SSEServer
}

func NewMcpService(ruleConfig types.Config, c config.Config, pool types.RuleEnginePool, componentService *ComponentService, username string) (*McpService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *McpService) GetRuleConfig() types.Config {
	_ = "STUB: not implemented"
	return *new(types.Config)
}

func (s *McpService) NewMCPServer() *server.MCPServer { _ = "STUB: not implemented"; return nil }

// 服务器名称
// 服务器版本

func (s *McpService) Callbacks() types.Callbacks {
	_ = "STUB: not implemented"
	return *new(types.Callbacks)
}

func (s *McpService) LoadTools() { _ = "STUB: not implemented"; return }

// 从组件列表添加工具

//if s.config.MCP.LoadChainsAsTool {//已经从Callbacks.OnNew 加载
//	s.LoadToolsFromChains()
//}

func (s *McpService) NewSSEServer(opts ...server.SSEOption) *server.SSEServer {
	_ = "STUB: not implemented"
	return nil
}

func (s *McpService) MCPServer() *server.MCPServer { _ = "STUB: not implemented"; return nil }

func (s *McpService) SSEServer() *server.SSEServer { _ = "STUB: not implemented"; return nil }

func (s *McpService) DeleteTools(names ...string) { _ = "STUB: not implemented"; return }

// LoadToolsFromComponents 从组件列表添加工具
func (s *McpService) LoadToolsFromComponents() { _ = "STUB: not implemented"; return }

// CheckExclude 检查组件是否需要排除
func (s *McpService) CheckExclude(name string, isComponent bool) bool {
	_ = "STUB: not implemented"
	return false
}

// AddToolsFromComponent 从组件定义添加工具
func (s *McpService) AddToolsFromComponent(name string, component types.ComponentForm) {
	_ = "STUB: not implemented"
	return
}

// 工具描述
// 添加工具
// 工具名称

// 为工具添加处理器

func (s *McpService) componentToolHandler(componentType string) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = "STUB: not implemented"
	return nil
}

// LoadToolsFromChains 从规则链列表添加工具
func (s *McpService) LoadToolsFromChains() { _ = "STUB: not implemented"; return }

func (s *McpService) AddToolsFromChain(id string, def types.RuleChain) {
	_ = "STUB: not implemented"
	return
}

//自动从所有节点中获取所有变量

// 为工具添加处理器

func (s *McpService) ruleChainToolHandler(chainId string) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = "STUB: not implemented"
	return nil
}

const (
	// ToolNameSaveRuleChain 保存规则链
	ToolNameSaveRuleChain = "saveRuleChain"
	// ToolNameListRuleChain 列出规则链
	ToolNameListRuleChain = "listRuleChain"
	// ToolNameDeleteRuleChain 删除规则链
	ToolNameDeleteRuleChain = "deleteRuleChain"
	// ToolNameExecuteRuleChain 执行规则链
	ToolNameExecuteRuleChain = "executeRuleChain"
)

// AddRuleApiTools 添加规则链工具
func (s *McpService) AddRuleApiTools() { _ = "STUB: not implemented"; return }

// AddListRuleTool 添加列出规则链工具
func (s *McpService) AddListRuleTool() { _ = "STUB: not implemented"; return }

// AddSaveRuleTool 添加保存规则链工具
func (s *McpService) AddSaveRuleTool() { _ = "STUB: not implemented"; return }

// AddDeleteRuleTool 添加删除规则链工具
func (s *McpService) AddDeleteRuleTool() { _ = "STUB: not implemented"; return }

func (s *McpService) AddExecuteRuleTool() { _ = "STUB: not implemented"; return }
