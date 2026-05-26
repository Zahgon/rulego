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
	"time"

	"github.com/rulego/rulego/api/types"

	"testing"
)

var (
	shareKey       = "shareKey"
	shareValue     = "shareValue"
	addShareKey    = "addShareKey"
	addShareValue  = "addShareValue"
	testdataFolder = "./testdata/"
	contentType    = "Content-Type"
	content        = "application/json"
)

// CreateAndInitNode 创建并初始化一个节点实例
func CreateAndInitNode(targetNodeType string, initConfig types.Configuration, registry *types.SafeComponentSlice) (types.Node, error) {
	_ = "STUB: not implemented"
	return *new(types.Node), nil
}

func InitNode(targetNodeType string, initConfig types.Configuration, registry *types.SafeComponentSlice) types.Node {
	_ = "STUB: not implemented"
	return *new(types.Node)
}

func InitNodeByConfig(config types.Config, targetNodeType string, initConfig types.Configuration, registry *types.SafeComponentSlice) types.Node {
	_ = "STUB: not implemented"
	return *new(types.Node)
}

// NodeNew 测试创建节点实例
func NodeNew(t *testing.T, targetNodeType string, targetNode types.Node, defaultConfig types.Configuration, registry *types.SafeComponentSlice) {
	_ = "STUB: not implemented"
	return
}

// NodeInit 测试初始化
func NodeInit(t *testing.T, targetNodeType string, initConfig types.Configuration, expected types.Configuration, registry *types.SafeComponentSlice) {
	_ = "STUB: not implemented"
	return
}

type NodeAndCallback struct {
	Node          types.Node
	MsgList       []Msg
	ChildrenNodes map[string]types.Node
	Callback      func(msg types.RuleMsg, relationType string, err error)
}

type Msg struct {
	Id       string
	Ts       int64
	MetaData *types.Metadata
	DataType types.DataType
	MsgType  string
	Data     string
	//发之后暂停间隔
	AfterSleep time.Duration
}

// NodeOnMsg 发送消息
func NodeOnMsg(t *testing.T, node types.Node, msgList []Msg, callback func(msg types.RuleMsg, relationType string, err error)) {
	_ = "STUB: not implemented"
	return
}

// NodeOnMsgWithChildren 发送消息
func NodeOnMsgWithChildren(t *testing.T, node types.Node, msgList []Msg, childrenNodes map[string]types.Node, callback func(msg types.RuleMsg, relationType string, err error)) {
	_ = "STUB: not implemented"
	return
}

func NodeOnMsgWithChildrenAndConfig(t *testing.T, config types.Config, node types.Node, msgList []Msg, childrenNodes map[string]types.Node, callback func(msg types.RuleMsg, relationType string, err error)) {
	_ = "STUB: not implemented"
	return
}

// UpperNode A plugin that converts the message data to uppercase
type UpperNode struct{}

func (n *UpperNode) Type() string { _ = "STUB: not implemented"; return "" }

func (n *UpperNode) New() types.Node { _ = "STUB: not implemented"; return *new(types.Node) }

func (n *UpperNode) Init(ruleConfig types.Config, configuration types.Configuration) error {
	_ = "STUB: not implemented"
	// Do some initialization work
	return nil
}

func (n *UpperNode) OnMsg(ctx types.RuleContext, msg types.RuleMsg) {
	_ = "STUB: not implemented"
	return
}

//增加新的共享数据

// Send the modified message to the next node

func (n *UpperNode) Destroy() {
	_ = "STUB: not implemented"
	// Do some cleanup work
	return
}

// TimeNode A plugin that adds a timestamp to the message metadata
type TimeNode struct{}

func (n *TimeNode) Type() string { _ = "STUB: not implemented"; return "" }

func (n *TimeNode) New() types.Node { _ = "STUB: not implemented"; return *new(types.Node) }

func (n *TimeNode) Init(ruleConfig types.Config, configuration types.Configuration) error {
	_ = "STUB: not implemented"
	// Do some initialization work
	return nil
}

func (n *TimeNode) OnMsg(ctx types.RuleContext, msg types.RuleMsg) {
	_ = "STUB: not implemented"
	return
}

// Send the modified message to the next node

func (n *TimeNode) Destroy() {
	_ = "STUB: not implemented"
	// Do some cleanup work
	return
}
