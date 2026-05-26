package controller

import (
	endpointApi "github.com/rulego/rulego/api/types/endpoint"
)

var Log = &log{}

type log struct {
}

// GetDebugLogs 创建获取节点调试数据路由
func (c *log) GetDebugLogs(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

func (c *log) List(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

func (c *log) Delete(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

func (c *log) WsNodeLogRouter(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}
