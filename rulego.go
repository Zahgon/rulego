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

// Package rulego provides a lightweight, high-performance, embedded, orchestrable component-based rule engine.
//
// # Usage
//
// Implement your business requirements by configuring components in the rule chain, and support dynamic modification. Rule chain definition format:
//
//	{
//		  "ruleChain": {
//			"id":"rule01"
//		  },
//		  "metadata": {
//		    "nodes": [
//		    ],
//			"connections": [
//			]
//		 }
//	}
//
// nodes:configure components. You can use built-in components and third-party extension components without writing any code.
//
// connections:configure the relation type between components. Determine the data flow.
//
// Example:
//
//	var ruleFile = `
//	{
//		"ruleChain": {
//		"id":"rule02",
//		"name": "test",
//		"root": true
//		},
//		"metadata": {
//		"nodes": [
//			{
//			"id": "s1",
//			"type": "jsTransform",
//			"name": "transform",
//			"debugMode": true,
//			"configuration": {
//				"jsScript": "metadata['state']='modify by js';\n msg['addField']='addValueFromJs'; return {'msg':msg,'metadata':metadata,'msgType':msgType};"
//				}
//			},
//			{
//				"id": "s2",
//				"type": "restApiCall",
//				"name": "push data",
//				"debugMode": true,
//				"configuration": {
//					"restEndpointUrlPattern": "http://127.0.0.1:9090/api/msg",
//					"requestMethod": "POST",
//				}
//			}
//		],
//		"connections": [
//			{
//				"fromId": "s1",
//				"toId": "s2",
//				"type": "Success"
//			}
//		]
//		}
//	}
//	`
//
// Create Rule Engine Instance
//
//	ruleEngine, err := rulego.New("rule01", []byte(ruleFile))
//
// Define Message Metadata
//
//	metaData := types.NewMetadata()
//	metaData.PutValue("productType", "test01")
//
// Define Message Payload And Type
//
//	msg := types.NewMsg(0, "TELEMETRY_MSG", types.JSON, metaData, "{\"temperature\":35}")
//
// Processing Message
//
//	ruleEngine.OnMsg(msg)
//
// Update Rule Chain
//
//	err := ruleEngine.ReloadSelf([]byte(ruleFile))
//
// Load All Rule Chain
//
//	err := ruleEngine.Load("./rulechains")
//
// Get Engine Instance
//
//	ruleEngine, ok := rulego.Get("rule01")
package rulego

import (
	"github.com/rulego/rulego/api/types"
	"github.com/rulego/rulego/builtin/aspect"
	"github.com/rulego/rulego/endpoint"
	"github.com/rulego/rulego/engine"
)

// Registry is the default registrar for rule engine components.
var Registry = engine.Registry

// Rules is the default instance of RuleGo with the rule engine pool set to the default pool.
var Rules = &RuleGo{pool: engine.DefaultPool}

var Endpoints = endpoint.DefaultPool

func init() {
	engine.BuiltinsAspects = append(engine.BuiltinsAspects, &aspect.EndpointAspect{
		EndpointPool: Endpoints,
	})
}

// Ensure RuleGo implements the RuleEnginePool interface.
var _ types.RuleEnginePool = (*RuleGo)(nil)

// RuleGo is a pool of rule engine instances.
type RuleGo struct {
	pool *engine.Pool
}

// NewRuleGo creates a new RuleGo instance.
func NewRuleGo() *RuleGo { _ = "STUB: not implemented"; return nil }

// Pool returns the rule engine pool.
func (g *RuleGo) Pool() *engine.Pool {
	_ = "STUB: not implemented"

	// Load loads all rule chain configurations from the specified folder and its subFolders into the rule engine instance pool.
	// The rule chain ID is taken from the ruleChain.id specified in the rule chain file.
	return nil
}

func (g *RuleGo) Load(folderPath string, opts ...types.RuleEngineOption) error {
	_ = "STUB: not implemented"
	return nil
}

// New creates a new RuleEngine and stores it in the RuleGo rule chain pool.
// If the specified id is empty (""), the ruleChain.id from the rule chain file is used.
func (g *RuleGo) New(id string, rootRuleChainSrc []byte, opts ...types.RuleEngineOption) (types.RuleEngine, error) {
	_ = "STUB: not implemented"
	return *new(types.RuleEngine), nil
}

// Get retrieves a rule engine instance by its ID.
func (g *RuleGo) Get(id string) (types.RuleEngine, bool) {
	_ = "STUB: not implemented"
	return *

	// Del removes a rule engine instance by its ID.
	new(types.RuleEngine), false
}

func (g *RuleGo) Del(id string) {
	_ = "STUB: not implemented"

	// Stop releases all rule engine instances.
	return
}

func (g *RuleGo) Stop() {
	_ = "STUB: not implemented"

	// Range iterates over all rule engine instances.
	return
}

func (g *RuleGo) Range(f func(key, value any) bool) {
	_ = "STUB: not implemented"

	// Reload reloads all rule engine instances.
	return
}

func (g *RuleGo) Reload(opts ...types.RuleEngineOption) { _ = "STUB: not implemented"; return }

// OnMsg calls all rule engine instances to process a message.
// All rule chains in the rule engine instance pool will attempt to process the message.
func (g *RuleGo) OnMsg(msg types.RuleMsg) { _ = "STUB: not implemented"; return }

// SetCallbacks sets the callbacks for the rule engine pool.
func (g *RuleGo) SetCallbacks(callbacks types.Callbacks) { _ = "STUB: not implemented"; return }

// Load loads all rule chain configurations from the specified folder and its subFolders into the rule engine instance pool.
// The rule chain ID is taken from the ruleChain.id specified in the rule chain file.
func Load(folderPath string, opts ...types.RuleEngineOption) error {
	_ = "STUB: not implemented"
	return nil
}

// New creates a new RuleEngine and stores it in the RuleGo rule chain pool.
func New(id string, rootRuleChainSrc []byte, opts ...types.RuleEngineOption) (types.RuleEngine, error) {
	_ = "STUB: not implemented"
	return *new(types.RuleEngine), nil
}

// Get retrieves a rule engine instance by its ID.
func Get(id string) (types.RuleEngine, bool) {
	_ = "STUB: not implemented"
	return *

	// Del removes a rule engine instance by its ID.
	new(types.RuleEngine), false
}

func Del(id string) {
	_ = "STUB: not implemented"

	// Stop releases all rule engine instances.
	return
}

func Stop() {
	_ = "STUB: not implemented"

	// OnMsg calls all rule engine instances to process a message.
	// All rule chains in the rule engine instance pool will attempt to process the message.
	return
}

func OnMsg(msg types.RuleMsg) {
	_ = "STUB: not implemented"

	// Reload reloads all rule engine instances.
	return
}

func Reload(opts ...types.RuleEngineOption) { _ = "STUB: not implemented"; return }

// Range iterates over all rule engine instances.
func Range(f func(key, value any) bool) {
	_ = "STUB: not implemented"

	// NewConfig creates a new Config and applies the options.
	return
}

func NewConfig(opts ...types.Option) types.Config {
	_ = "STUB: not implemented"
	return *new(types.Config)
}

// WithConfig is an option that sets the Config of the RuleEngine.
func WithConfig(config types.Config) types.RuleEngineOption {
	_ = "STUB: not implemented"
	return *new(types.RuleEngineOption)
}
