package controller

import (
	"github.com/rulego/rulego/api/types"
	endpointApi "github.com/rulego/rulego/api/types/endpoint"
)

var Rule = &rule{}

type rule struct {
}

// Get 创建获取指定规则链路由
func (c *rule) Get(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

// GetLatest 获取最近修改的规则链
func (c *rule) GetLatest(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

// Save 创建保存/更新指定规则链路由
func (c *rule) Save(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

// List 创建获取所有规则链路由
func (c *rule) List(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

func (c *rule) list(getFromMarketplace bool, exchange *endpointApi.Exchange) bool {
	_ = "STUB: not implemented"
	return false
}

//从组件市场获取规则链

// Delete 创建删除指定规则链路由
func (c *rule) Delete(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

// SaveBaseInfo 保存规则链扩展信息
func (c *rule) SaveBaseInfo(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

// SaveConfiguration 保存规则链配置
func (c *rule) SaveConfiguration(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

func (c *rule) transformMsg(router endpointApi.Router, exchange *endpointApi.Exchange) bool {
	_ = "STUB: not implemented"
	return false
}

//获取消息类型

//把http header放入消息元数据

//if msg.Metadata.GetValue(constants.KeySetWorkDir)=="true"{
//	username := msg.Metadata.GetValue(constants.KeyUsername)
//	//设置工作目录
//	var paths = []string{config.C.DataDir, constants.DirWorkflows, username, constants.DirWorkflowsRule}
//	msg.Metadata.PutValue(constants.KeyWorkDir, path.Join(paths...)
//}

// Execute 处理请求，并转发到规则引擎，同步等待规则链执行结果返回给调用方
// .To("chain:${id}") 这段逻辑相当于：
//
//	engine,err:=pool.Get(chainId)
//	engine.OnMsgAndWait(msg)
func (c *rule) Execute(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

//错误

//把处理结果响应给客户端，http endpoint 必须增加 Wait()，否则无法正常响应

// PostMsg 处理请求，并转发到规则引擎
// .To("chain:${id}") 这段逻辑相当于：
//
//	engine,err:=pool.Get(chainId)
//	engine.OnMsg(msg)
func (c *rule) PostMsg(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

func (c *rule) addWithOnRuleChainCompleted() types.RuleContextOption {
	_ = "STUB: not implemented"
	return *new(types.RuleContextOption)
}

// Operate 部署/下架规则链
func (c *rule) Operate(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}
