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

package engine

import (
	"github.com/rulego/rulego/api/types"
)

// JsonParser Json
type JsonParser struct {
}

// DecodeRuleChain 通过json解析规则链结构体
func (p *JsonParser) DecodeRuleChain(rootRuleChain []byte) (types.RuleChain, error) {
	_ = "STUB: not implemented"
	return *new(types.RuleChain), nil
}

// DecodeRuleNode 通过json解析节点结构体
func (p *JsonParser) DecodeRuleNode(rootRuleChain []byte) (types.RuleNode, error) {
	_ = "STUB: not implemented"
	return *new(types.RuleNode), nil
}

func (p *JsonParser) EncodeRuleChain(def interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//格式化Json

func (p *JsonParser) EncodeRuleNode(def interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//格式化Json
