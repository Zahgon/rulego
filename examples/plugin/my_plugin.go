package main

import (
	"github.com/rulego/rulego/api/types"
)

//go build -buildmode=plugin -o plugin.so plugin.go # Compile the plugin and generate the plugin.so file
//need to compile in mac or linux environment

// Plugins plugin entry point
var Plugins MyPlugins

type MyPlugins struct{}

func (p *MyPlugins) Init() error { _ = "STUB: not implemented"; return nil }

func (p *MyPlugins) Components() []types.Node { _ = "STUB: not implemented"; return nil }

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

type FilterNode struct {
	blacklist map[string]bool
}

func (n *FilterNode) Type() string { _ = "STUB: not implemented"; return "" }

func (n *FilterNode) New() types.Node { _ = "STUB: not implemented"; return *new(types.Node) }

func (n *FilterNode) Init(ruleConfig types.Config, configuration types.Configuration) error {
	_ = "STUB: not implemented"
	// Do some initialization work
	return nil
}

func (n *FilterNode) OnMsg(ctx types.RuleContext, msg types.RuleMsg) {
	_ = "STUB: not implemented"
	return
}

// Skip the message and do not send it to the next node

// Send the message to the next node

func (n *FilterNode) Destroy() {
	_ = "STUB: not implemented"
	// Do some cleanup work
	return
}
