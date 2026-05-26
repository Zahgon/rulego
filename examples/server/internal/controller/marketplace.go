package controller

import (
	endpointApi "github.com/rulego/rulego/api/types/endpoint"
)

// MarketplaceComponents 获取组件市场动态组件
func (c *node) MarketplaceComponents(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

//是否检查自己的组件

func (c *rule) MarketplaceChains(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}
