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

// Package websocket provides a WebSocket endpoint implementation for the RuleGo framework.
// It allows creating WebSocket servers that can receive and process incoming WebSocket messages,
// routing them to appropriate rule chains or components for further processing.
//
// Key components in this package include:
// - Endpoint (alias Websocket): Implements the WebSocket server and message handling
// - RequestMessage: Represents an incoming WebSocket message
// - ResponseMessage: Represents the WebSocket message to be sent back
//
// The WebSocket endpoint supports dynamic routing configuration, allowing users to
// define message patterns and their corresponding rule chain or component destinations.
// It also provides flexibility in handling different WebSocket message types and formats.
//
// This package integrates with the broader RuleGo ecosystem, enabling seamless
// data flow from WebSocket messages to rule processing and back to WebSocket responses.
package websocket

import (
	"net/http"
	"net/textproto"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/rulego/rulego/api/types"
	"github.com/rulego/rulego/api/types/endpoint"
	"github.com/rulego/rulego/endpoint/rest"
)

// Type 组件类型
const Type = types.EndpointTypePrefix + "ws"

// Endpoint 别名
type Endpoint = Websocket

// RequestMessage websocket请求消息
type RequestMessage struct {
	//ws消息类型 TextMessage=1/BinaryMessage=2
	messageType int
	request     *http.Request
	body        []byte
	//路径参数
	Params httprouter.Params
	msg    *types.RuleMsg
	err    error
}

func (r *RequestMessage) Body() []byte { _ = "STUB: not implemented"; return nil }

func (r *RequestMessage) Headers() textproto.MIMEHeader {
	_ = "STUB: not implemented"
	return *new(textproto.MIMEHeader)
}

func (r RequestMessage) From() string { _ = "STUB: not implemented"; return "" }

func (r *RequestMessage) GetParam(key string) string { _ = "STUB: not implemented"; return "" }

func (r *RequestMessage) SetMsg(msg *types.RuleMsg) { _ = "STUB: not implemented"; return }

func (r *RequestMessage) GetMsg() *types.RuleMsg {
	_ = "STUB: not implemented"

	// 默认指定是JSON格式，如果不是该类型，请在process函数中修改
	return nil
}

func (r *RequestMessage) SetStatusCode(statusCode int) { _ = "STUB: not implemented"; return }

func (r *RequestMessage) SetBody(body []byte) { _ = "STUB: not implemented"; return }

func (r *RequestMessage) SetError(err error) { _ = "STUB: not implemented"; return }

func (r *RequestMessage) GetError() error { _ = "STUB: not implemented"; return nil }

func (r *RequestMessage) Request() *http.Request {
	_ = "STUB: not implemented"

	// ResponseMessage websocket响应消息
	return nil
}

type ResponseMessage struct {
	headers textproto.MIMEHeader
	//ws消息类型 TextMessage/BinaryMessage
	messageType int
	log         func(format string, v ...interface{})
	request     *http.Request
	conn        *websocket.Conn
	body        []byte
	to          string
	msg         *types.RuleMsg
	err         error
	locker      sync.RWMutex
}

func (r *ResponseMessage) Body() []byte { _ = "STUB: not implemented"; return nil }

func (r *ResponseMessage) Headers() textproto.MIMEHeader {
	_ = "STUB: not implemented"
	return *new(textproto.MIMEHeader)
}

func (r *ResponseMessage) From() string { _ = "STUB: not implemented"; return "" }

func (r *ResponseMessage) GetParam(key string) string { _ = "STUB: not implemented"; return "" }

func (r *ResponseMessage) SetMsg(msg *types.RuleMsg) { _ = "STUB: not implemented"; return }

func (r *ResponseMessage) GetMsg() *types.RuleMsg { _ = "STUB: not implemented"; return nil }

// SetStatusCode 不提供设置状态码
func (r *ResponseMessage) SetStatusCode(statusCode int) { _ = "STUB: not implemented"; return }

func (r *ResponseMessage) SetBody(body []byte) {
	_ = "STUB: not implemented"
	// 在设置body和写入WebSocket之前加锁
	return
}

func (r *ResponseMessage) SetError(err error) { _ = "STUB: not implemented"; return }

func (r *ResponseMessage) GetError() error { _ = "STUB: not implemented"; return nil }

// Config Websocket 服务配置
type Config = rest.Config

// Websocket 接收端端点
type Websocket struct {
	*rest.Rest
	//配置
	Config   Config
	Upgrader websocket.Upgrader
}

// Type 组件类型
func (ws *Websocket) Type() string { _ = "STUB: not implemented"; return "" }

func (ws *Websocket) New() types.Node { _ = "STUB: not implemented"; return *new(types.Node) }

// Init 初始化
func (ws *Websocket) Init(ruleConfig types.Config, configuration types.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

// 允许所有跨域请求

func (ws *Websocket) Id() string { _ = "STUB: not implemented"; return "" }

func (ws *Websocket) AddRouter(router endpoint.Router, params ...interface{}) (id string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ws *Websocket) RemoveRouter(routerId string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (ws *Websocket) Printf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (ws *Websocket) Start() error { _ = "STUB: not implemented"; return nil }

// 允许所有跨域请求

// addRouter 注册1个或者多个路由
func (ws *Websocket) addRouter(routers ...endpoint.Router) *Websocket {
	_ = "STUB: not implemented"
	return nil
}

//存储路由

//添加到http路由器

func (ws *Websocket) handler(router endpoint.Router) httprouter.Handle {
	_ = "STUB: not implemented"
	return *new(httprouter.Handle)
}

//捕捉异常

//ws.Printf("recv:", string(message))

//把路径参数放到msg元数据中

//把url?参数放到msg元数据中
