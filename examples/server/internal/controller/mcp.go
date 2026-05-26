package controller

import (
	endpointApi "github.com/rulego/rulego/api/types/endpoint"
)

var MCP = &mcp{}

type mcp struct {
}

func (c *mcp) Handler(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

//不允许匿名访问
