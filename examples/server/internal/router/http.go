package router

import (
	"examples/server/config"

	"github.com/rulego/rulego/api/types"
	endpointApi "github.com/rulego/rulego/api/types/endpoint"
	"github.com/rulego/rulego/node_pool"
)

const (
	// base HTTP paths.
	apiVersion  = "v1"
	apiBasePath = "/api/" + apiVersion
	moduleFlows = "rules"
	// moduleDcs 动态组件
	moduleDynamicComponents = "dynamic-components"
	// moduleSharedNodes 共享组件
	moduleSharedNodes = "shared-nodes"
	moduleLocales     = "locales"
	moduleLogs        = "logs"
	moduleMarketplace = "marketplace"
	ContentTypeKey    = "Content-Type"
	JsonContextType   = "application/json"
)

// SystemRulegoConfig 系统rulego配置
var SystemRulegoConfig types.Config

// SystemNodePool 系统内部节点池
var SystemNodePool *node_pool.NodePool

func InitRulegoConfig() { _ = "STUB: not implemented"; return }

// NewRestServe rest服务 接收端点
func NewRestServe(config config.Config) (endpointApi.HttpEndpoint, error) {
	_ = "STUB: not implemented"
	//初始化日志
	return *new(endpointApi.HttpEndpoint), nil
}

//添加全局拦截器

//重定向UI界面

//创建获取所有规则引擎组件列表路由

//获取所有共享组件

//获取组件市场组件列表

//获取组件市场规则链列表

//获取用户所有自定义动态组件列表

//获取自定义动态组件DSL

//安装/升级自定义动态组件

//卸装自定义动态组件

//获取所有规则链列表

//获取最新修改的规则链DSL 实际是：/api/v1/rules/get/latest

//获取规则链DSL

//新增/修改规则链DSL

//删除规则链

//保存规则链附加信息

//保存规则链配置信息

//执行规则链,并得到规则链处理结果

//处理数据上报请求，并转发到规则引擎，不等待规则引擎处理结果

//部署或者下线规则链

//获取节点调试日志列表

//获取规则链运行日志列表

//获取规则链运行日志详情

//创建用户登录路由

// 加载静态文件映射

//把默认HTTP服务设置成共享节点

//把默认HTTP服务添加到系统节点池

// LoadServeFiles 加载静态文件映射
func LoadServeFiles(c config.Config, restEndpoint endpointApi.HttpEndpoint) {
	_ = "STUB: not implemented"
	return
}
