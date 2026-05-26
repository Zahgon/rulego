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
	"examples/server/config"
	"examples/server/internal/dao"

	"github.com/rulego/rulego/api/types"
)

// ComponentService 自定义组件服务
type ComponentService struct {
	username     string
	config       config.Config
	ruleConfig   types.Config
	componentDao *dao.ComponentDao
	mcpService   *McpService
}

func NewComponentService(ruleConfig types.Config, c config.Config, username string) (*ComponentService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ComponentService) GetRuleConfig() types.Config {
	_ = "STUB: not implemented"
	return *new(types.Config)
}

func (s *ComponentService) LoadComponents() { _ = "STUB: not implemented"; return }

func (s *ComponentService) ComponentsRegistry() types.ComponentRegistry {
	_ = "STUB: not implemented"
	return *new(types.ComponentRegistry)
}

func (s *ComponentService) List(keywords string, size, page int) ([]types.RuleChain, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (s *ComponentService) Get(nodeType string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ComponentService) Install(id string, dsl []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ComponentService) Upgrade(id string, dsl []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ComponentService) Uninstall(nodeType string) error { _ = "STUB: not implemented"; return nil }
