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
	"errors"

	"github.com/rulego/rulego/api/types"
	"github.com/rulego/rulego/utils/schema"
)

// ErrRuleEnginePoolNil rule engine pool is nil
var ErrRuleEnginePoolNil = errors.New("rule engine pool is nil")

// ErrDSLEmpty dsl is empty
var ErrDSLEmpty = errors.New("dsl is empty")

// DynamicNode 通过子规则链动态定义节点组件
// ruleChain.id: 定义组件类型
// ruleChain.name: 定义组件label
// ruleChain.additionalInfo.category: 定义组件分类
// ruleChain.additionalInfo.icon: 定义组件图标
// ruleChain.additionalInfo.description: 定义组件描述
// ruleChain.additionalInfo.inputSchema: 使用JSON Schema 定义组件的输入参数(组件参数配置)
// ruleChain.additionalInfo.relationTypes: 定义和下一个节点允许连接关系类型
// 组件通过 ${vars.xx} 方式获取组件配置参数
// 使用示例：
// 通过dsl定义组件：
// dynamicNode := NewDynamicNode("fahrenheit", `
//
//			 {
//			 "ruleChain": {
//			   "id": "fahrenheit",
//			   "name": "华氏温度转换",
//			   "debugMode": false,
//			   "root": false,
//			   "additionalInfo": {
//			     "layoutX": 720,
//			     "layoutY": 260,
//		         "description":"this is a description",
//			     "relationTypes":["Success","Failure"],
//			     "inputSchema": {
//			       "type": "object",
//			       "properties": {
//			         "scaleFactor": {
//			           "type": "number",
//	                  "title": "换算系数",
//	                  "default": 1.8
//			         }
//			       },
//			       "required": ["scaleFactor"]
//			     }
//
//			   }
//			 },
//			 "metadata": {
//			   "firstNodeIndex": 0,
//			   "nodes": [
//			     {
//			       "id": "s2",
//			       "type": "jsTransform",
//			       "name": "摄氏温度转华氏温度",
//			       "debugMode": true,
//			       "configuration": {
//			         "jsScript": "var newMsg={'temperature': msg.temperature*vars.scaleFactor+32};\n return {'msg':newMsg,'metadata':metadata,'msgType':msgType};"
//			       }
//			     }
//			   ],
//			   "connections": [
//			     {
//			     }
//			   ]
//			 }
//			}
//
//		`)
//		注册组件
//		Registry.Register(dynamicNode)
type DynamicNode struct {
	//ComponentType 组件类型
	ComponentType string
	//Dsl 子规则链 DSL
	Dsl string
	//实例化的节点配置
	instantiatedConfig types.Configuration
	//实例化规则引擎
	ruleEngine types.RuleEngine
}

func NewDynamicNode(componentType, componentDsl string) *DynamicNode {
	_ = "STUB: not implemented"
	return nil
}

// Type 组件类型
func (x *DynamicNode) Type() string { _ = "STUB: not implemented"; return "" }

func (x *DynamicNode) New() types.Node { _ = "STUB: not implemented"; return *new(types.Node) }

// Init 初始化
func (x *DynamicNode) Init(ruleConfig types.Config, configuration types.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

//把组件配置和跟规则链vars复制到当前组件定义的vars

//动态初始化子规则链

// OnMsg 处理消息
func (x *DynamicNode) OnMsg(ctx types.RuleContext, msg types.RuleMsg) {
	_ = "STUB: not implemented"
	return
}

// Destroy 销毁
func (x *DynamicNode) Destroy() { _ = "STUB: not implemented"; return }

// Def 组件定义
func (x *DynamicNode) Def() types.ComponentForm {
	_ = "STUB: not implemented"
	return *new(types.ComponentForm)
}

// 获取关系类型

// 获取输入参数定义

// 获取字段列表并排序

// processFieldAuto 处理自动生成字段 生成规则：提取 ${vars.xx}变量
func (x *DynamicNode) processFieldAuto(def types.RuleChain) types.ComponentFormFieldList {
	_ = "STUB: not implemented"
	return *new(types.ComponentFormFieldList)
}

// 找到所有匹配的变量

// processField 处理单个字段，支持嵌套字段
func (x *DynamicNode) processField(name string, fieldMap schema.FieldSchema, parentSchema schema.JSONSchema) types.ComponentFormField {
	_ = "STUB: not implemented"
	return *new(types.ComponentFormField)
}

// 获取子字段列表并排序

func (x *DynamicNode) copyVars(targetRuleChain types.RuleChain, fromRootChain *types.RuleChain, fromNodeConfig types.Configuration) types.RuleChain {
	_ = "STUB: not implemented"
	return *new(types.RuleChain)
}
