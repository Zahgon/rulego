package service

import (
	"sync"

	"github.com/rulego/rulego/components/action"
)

var (
	builtins = make(map[string]interface{})
	lock     sync.RWMutex
)

func init() {
	// functions节点组件
	builtins["functions"] = map[string]interface{}{
		//函数名选项
		"functionName": action.Functions.Names(),
	}
}

// Builtins 获取内置组件配置选项
func Builtins() map[string]interface{} { _ = "STUB: not implemented"; return nil }

// RegisterBuiltin 注册内置组件配置选项
// value 可以是静态值，也可以是 func() interface{} 类型的函数，用于实时获取数据
func RegisterBuiltin(name string, value interface{}) { _ = "STUB: not implemented"; return }
