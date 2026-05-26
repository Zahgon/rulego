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

package filter

//规则链节点配置示例：
//{
//        "id": "s2",
//        "type": "jsSwitch",
//        "name": "脚本路由",
//        "debugMode": false,
//        "configuration": {
//          "jsScript": "return ['one','two'];"
//        }
//      }
import (
	"errors"

	"github.com/rulego/rulego/api/types"
)

// JsSwitchReturnFormatErr JavaScript脚本必须返回数组
var JsSwitchReturnFormatErr = errors.New("return the value is not an array")

// init 注册JsSwitchNode组件
func init() {
	Registry.Add(&JsSwitchNode{})
}

// JsSwitchNodeConfiguration JsSwitchNode配置结构
type JsSwitchNodeConfiguration struct {
	// JsScript JavaScript脚本，用于确定消息路由路径
	// 函数参数：msg, metadata, msgType, dataType
	// 必须返回字符串数组，表示路由关系类型
	//
	// 内置变量：
	//   - $ctx: 上下文对象，提供缓存操作
	//   - global: 全局配置属性
	//   - vars: 规则链变量
	//   - UDF函数: 用户自定义函数
	//
	// 示例: "return ['route1', 'route2'];"
	JsScript string
}

// JsSwitchNode 使用JavaScript确定消息路由路径的开关节点
type JsSwitchNode struct {
	// Config 节点配置
	Config JsSwitchNodeConfiguration

	// jsEngine JavaScript执行引擎
	jsEngine types.JsEngine

	// defaultRelationType 默认关系类型
	defaultRelationType string
}

// Type 返回组件类型
func (x *JsSwitchNode) Type() string {
	_ = "STUB: not implemented"

	// New 创建新实例
	return ""
}

func (x *JsSwitchNode) New() types.Node { _ = "STUB: not implemented"; return *new(types.Node) }

// Init 初始化节点
func (x *JsSwitchNode) Init(ruleConfig types.Config, configuration types.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

// OnMsg 处理消息，执行JavaScript脚本确定路由路径
func (x *JsSwitchNode) OnMsg(ctx types.RuleContext, msg types.RuleMsg) {
	_ = "STUB: not implemented"
	// 准备传递给JS脚本的数据
	return
}

// Destroy 清理资源
func (x *JsSwitchNode) Destroy() { _ = "STUB: not implemented"; return }
