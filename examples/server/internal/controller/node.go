package controller

import (
	"github.com/rulego/rulego/api/types"
	endpointApi "github.com/rulego/rulego/api/types/endpoint"
)

var Node = &node{}

type node struct {
}

// Components 创建获取规则引擎节点组件列表路由
func (c *node) Components(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

//组件配置内置选项

// endpoints内置路由选项

//in 处理器列表

//in 处理器列表

//共享节点池

//响应endpoint和节点组件配置表单列表

//endpoint组件

//节点组件

//组件配置内置选项

// ListNodePool 获取所有共享组件
func (c *node) ListNodePool(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

// CustomNodeList 获取用户所有自定义动态组件
func (c *node) CustomNodeList(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

type ComponentList struct {
	Page  int               `json:"page"`
	Size  int               `json:"size"`
	Total int               `json:"total"`
	Items []types.RuleChain `json:"items"`
}

// CustomNodeList 获取用户所有自定义动态组件，默认从本地默认用户的自定义组件获取，如果配置了MarketBaseUrl，则从组件市场获取
// - checkMy:true，检查当前用户对应的组件是否需要升级，是否已安装
func (c *node) getCustomNodeList(getFromMarketplace bool, checkMy bool, exchange *endpointApi.Exchange) bool {
	_ = "STUB: not implemented"
	return false
}

//从组件市场获取组件

//获取当前用户已经安装的组件

//标记已安装、需要升级的组件

// CustomNodeDSL 获取动态组件DSL定义
func (c *node) CustomNodeDSL(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

func (c *node) customNodeDSL(username string, exchange *endpointApi.Exchange) bool {
	_ = "STUB: not implemented"
	return false
}

// CustomNodeInstall 安装自定义动态组件
func (c *node) CustomNodeInstall(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

// CustomNodeUpgrade 安装/升级自定义动态组件
func (c *node) CustomNodeUpgrade(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

func (c *node) customNodeInstall(username string, upgrade bool, exchange *endpointApi.Exchange) bool {
	_ = "STUB: not implemented"
	return false
}

// CustomNodeUninstall 卸载自定义动态组件
func (c *node) CustomNodeUninstall(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}
