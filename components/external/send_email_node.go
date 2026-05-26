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

package external

import (
	"net/smtp"
	"time"

	"github.com/rulego/rulego/api/types"
	"github.com/rulego/rulego/utils/el"
)

// 分隔符
const splitUserSep = ","

func init() {
	Registry.Add(&SendEmailNode{})
}

// Email 邮件消息体
type Email struct {
	//From 发件人邮箱
	From string `json:"from"`
	//To 收件人邮箱，多个与`,`隔开
	To string `json:"to"`
	//Cc 抄送人邮箱，多个与`,`隔开
	Cc string `json:"cc"`
	//Bcc 密送人邮箱，多个与`,`隔开
	Bcc string `json:"bcc"`
	//Subject 邮件主题，可以使用 ${metadata.key} 读取元数据中的变量或者使用 ${msg.key} 读取消息负荷中的变量进行替换
	Subject string `json:"subject"`
	//Body 邮件模板，可以使用 ${metadata.key} 读取元数据中的变量或者使用 ${msg.key} 读取消息负荷中的变量进行替换
	Body string `json:"body"`
}

// EmailTemplates 邮件模板结构体，统一管理所有邮件字段的模板
type EmailTemplates struct {
	// fromTemplate 发件人模板
	fromTemplate el.Template
	// toTemplate 收件人模板
	toTemplate el.Template
	// ccTemplate 抄送人模板
	ccTemplate el.Template
	// bccTemplate 密送人模板
	bccTemplate el.Template
	// subjectTemplate 主题模板
	subjectTemplate el.Template
	// bodyTemplate 正文模板
	bodyTemplate el.Template
	// hasVar 标识模板是否包含变量
	hasVar bool
}

// initTemplates 初始化邮件模板
// Initialize email templates
// initTemplates 初始化所有邮件字段的模板
func (x *SendEmailNode) initTemplates() error {
	_ = "STUB: not implemented"

	// 创建发件人模板
	return nil
}

// 创建收件人模板

// 创建抄送人模板

// 创建密送人模板

// 创建主题模板

// 创建正文模板

// 检查是否包含变量

// createEmailMsg 创建邮件消息内容
func (x *SendEmailNode) createEmailMsg(ctx types.RuleContext, ruleMsg types.RuleMsg) ([]byte, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 执行模板渲染

// 将所有的收件人、抄送和密送合并为一个切片

// 创建一个邮件消息，符合RFC 822标准

func (x *SendEmailNode) SendEmail(ctx types.RuleContext, ruleMsg types.RuleMsg, addr string, auth smtp.Auth, connectTimeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// 获取渲染后的发件人地址

// 调用SendMail函数发送邮件

func (x *SendEmailNode) SendEmailWithTls(ctx types.RuleContext, ruleMsg types.RuleMsg, addr string, auth smtp.Auth, connectTimeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// TLS

// Auth

// To && From
// 获取渲染后的发件人地址

// Data

// SendEmailConfiguration 配置
type SendEmailConfiguration struct {
	//SmtpHost Smtp主机地址
	SmtpHost string `json:"smtpHost"`
	//SmtpPort Smtp端口
	SmtpPort int `json:"smtpPort"`
	//Username 用户名
	Username string `json:"username"`
	//Password 授权码
	Password string `json:"password"`
	//EnableTls 是否是使用tls方式
	EnableTls bool `json:"enableTls"`
	//Email 邮件内容配置
	Email Email `json:"email"`
	//ConnectTimeout 连接超时，单位秒
	ConnectTimeout int
}

// SendEmailNode 通过SMTP服务器发送邮消息
// 如果请求成功，发送消息到`Success`链, 否则发到`Failure`链，
type SendEmailNode struct {
	//节点配置
	Config                 SendEmailConfiguration
	ConnectTimeoutDuration time.Duration
	smtpAddr               string
	smtpAuth               smtp.Auth
	// templates 邮件模板管理器
	templates EmailTemplates
}

// Type 组件类型
func (x *SendEmailNode) Type() string { _ = "STUB: not implemented"; return "" }

func (x *SendEmailNode) New() types.Node { _ = "STUB: not implemented"; return *new(types.Node) }

// Init 初始化
func (x *SendEmailNode) Init(ruleConfig types.Config, configuration types.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

// 初始化邮件模板

// 创建一个PLAIN认证

// OnMsg 处理消息
func (x *SendEmailNode) OnMsg(ctx types.RuleContext, msg types.RuleMsg) {
	_ = "STUB: not implemented"
	return
}

// Destroy 销毁
func (x *SendEmailNode) Destroy() { _ = "STUB: not implemented"; return }
