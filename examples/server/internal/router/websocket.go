package router

import (
	"examples/server/config"

	endpointApi "github.com/rulego/rulego/api/types/endpoint"
	"github.com/rulego/rulego/endpoint"
)

// NewWebsocketServe Websocket服务 接收端点
func NewWebsocketServe(c config.Config, httpEndpoint endpointApi.HttpEndpoint) (endpoint.Endpoint, error) {
	_ = "STUB: not implemented"

	// 使用Registry创建websocket端点，使用HTTP端点作为底层端点
	return *new(endpoint.Endpoint), nil
}

//写入报错
