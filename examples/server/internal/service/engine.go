package service

import (
	"examples/server/config"
	"examples/server/internal/dao"
	"log"
	"sync"

	"github.com/rulego/rulego"
	"github.com/rulego/rulego/api/types"
)

var UserRuleEngineServiceImpl *UserRuleEngineService

// UserRuleEngineService 用户规则引擎池
type UserRuleEngineService struct {
	Pool   map[string]*RuleEngineService
	config config.Config
	locker sync.RWMutex
}

func NewUserRuleEngineServiceImpl(c config.Config) (*UserRuleEngineService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//创建文件夹

//初始化内置用户

//检查是否有默认用户

// 创建用户
func (s *UserRuleEngineService) createUser(username string) error {
	_ = "STUB: not implemented"
	return nil
}

// Get 根据用户获取规则引擎池
func (s *UserRuleEngineService) Get(username string) (*RuleEngineService, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *UserRuleEngineService) Init(username string) (*RuleEngineService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type RuleEngineService struct {
	Pool       *rulego.RuleGo
	username   string
	config     config.Config
	ruleConfig types.Config
	logger     *log.Logger
	//基于内存的节点调试数据管理器
	//如果需要查询历史数据，请把调试日志数据存放数据库等可以持久化载体
	ruleChainDebugData *RuleChainDebugData
	onDebugObserver    map[string]*DebugObserver
	ruleDao            *dao.RuleDao
	locker             sync.RWMutex
	userSettingDao     *dao.UserSettingDao
	mainRuleEngine     types.RuleEngine
	componentService   *ComponentService
	mcpService         *McpService
}

func NewRuleEngineServiceAndInitRuleGo(c config.Config, username string) (*RuleEngineService, error) {
	_ = "STUB: not implemented"
	//隔离每个用户的自定义组价注册器
	return nil, nil
}

//初始化规则链

func NewRuleEngineService(c config.Config, ruleConfig types.Config, username string) (*RuleEngineService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//基于内存的节点调试数据管理器

func (s *RuleEngineService) GetRuleConfig() types.Config {
	_ = "STUB: not implemented"
	return *new(types.Config)
}

func (s *RuleEngineService) ComponentService() *ComponentService {
	_ = "STUB: not implemented"
	return nil
}

func (s *RuleEngineService) MCPService() *McpService { _ = "STUB: not implemented"; return nil }

func (s *RuleEngineService) ExecuteAndWait(chainId string, msg types.RuleMsg, opts ...types.RuleContextOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *RuleEngineService) Execute(chainId string, msg types.RuleMsg, opts ...types.RuleContextOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Get 获取DSL
func (s *RuleEngineService) Get(chainId string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *RuleEngineService) GetLatest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SaveAndLoad 保存或者更新DSL,并根据规则链状态部署或下架规则
func (s *RuleEngineService) SaveAndLoad(chainId string, def []byte) error {
	_ = "STUB: not implemented"
	//设置最新修改规则链
	return nil
}

//修改更新时间

//持久化规则链

//下架规则

//部署规则链

// List 获取所有规则链
func (s *RuleEngineService) List(keywords string, root *bool, disabled *bool, size, page int) ([]types.RuleChain, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Delete 删除规则链
func (s *RuleEngineService) Delete(chainId string) error { _ = "STUB: not implemented"; return nil }

// SaveBaseInfo 保存规则链基本信息
func (s *RuleEngineService) SaveBaseInfo(chainId string, baseInfo types.RuleChainBaseInfo) error {
	_ = "STUB: not implemented"

	//设置最新修改规则链
	return nil
}

//填充更新时间

//修改更新时间

// SaveConfiguration 保存规则链配置
func (s *RuleEngineService) SaveConfiguration(chainId string, key string, configuration interface{}) error {
	_ = "STUB: not implemented"

	//设置最新修改规则链
	return nil
}

//修改更新时间

// Deploy 部署规则链，创建规则链引擎实例，并发规则链状态disabled设置成启用状态
func (s *RuleEngineService) Deploy(chainId string) error { _ = "STUB: not implemented"; return nil }

// Load 加载规则链，创建规则链引擎实例，如果规则链状态=disabled则不创建
func (s *RuleEngineService) Load(chainId string) error { _ = "STUB: not implemented"; return nil }

//s.ruleConfig.Logger.Printf("chainId:%s load error: %s", chainId, err.Error())

// Undeploy 下架规则链引擎实例，并把规则链状态置为disabled
func (s *RuleEngineService) Undeploy(chainId string) error { _ = "STUB: not implemented"; return nil }

//持久化规则链

// SetMainChainId 设置主规则链
func (s *RuleEngineService) SetMainChainId(chainId string) error {
	_ = "STUB: not implemented"
	return nil
}

// saveRuleChain 持久化规则链
func (s *RuleEngineService) saveRuleChain(ruleChain types.RuleChain, whenErr error) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *RuleEngineService) GetEngine(chainId string) (types.RuleEngine, bool) {
	_ = "STUB: not implemented"
	return *new(types.RuleEngine), false
}

// OnDebug 调试日志
func (s *RuleEngineService) OnDebug(chainId, flowType string, nodeId string, msg types.RuleMsg, relationType string, err error) {
	_ = "STUB: not implemented"
	return
}

func (s *RuleEngineService) AddOnDebugObserver(chainId string, clientId string, fn func(chainId, flowType string, nodeId string, msg types.RuleMsg, relationType string, err error)) {
	_ = "STUB: not implemented"
	return
}

func (s *RuleEngineService) RemoveOnDebugObserver(clientId string) {
	_ = "STUB: not implemented"
	return
}

func (s *RuleEngineService) DebugData() *RuleChainDebugData { _ = "STUB: not implemented"; return nil }

// InitRuleGo 初始化规则链池
func (s *RuleEngineService) InitRuleGo(logger *log.Logger, workspacePath string, username string) {
	_ = "STUB: not implemented"
	return
}

//加载自定义配置

//加载lua第三方库

//把日志记录到内存管理器，用于界面显示

//节点ID

//流向OUT/IN

//消息

//关系

//Err 错误

//加载js

//加载组件插件

// 优先加载自定义组件

//加载规则链

// 加载js
func (s *RuleEngineService) loadJs(folderPath string) error {
	_ = "STUB: not implemented"
	// 创建文件夹
	return nil
}

//遍历所有文件

// 加载组件插件
func (s *RuleEngineService) loadPlugins(folderPath string) error {
	_ = "STUB: not implemented"
	//创建文件夹
	return nil
}

//遍历所有文件

// 加载规则链
func (s *RuleEngineService) loadRules(folderPath string) error {
	_ = "STUB: not implemented"
	//创建文件夹
	return nil
}

//遍历所有.json文件

// Get all file paths that match the pattern.

// Load each file and create a new rule engine instance from its contents.

//加载主规则链

// fillAdditionalInfo 填充扩展字段
func (s *RuleEngineService) fillAdditionalInfo(def *types.RuleChain) {
	_ = "STUB: not implemented"
	//修改更新时间
	return
}

type DebugObserver struct {
	chainId  string
	clientId string
	fn       func(chainId, flowType string, nodeId string, msg types.RuleMsg, relationType string, err error)
}
