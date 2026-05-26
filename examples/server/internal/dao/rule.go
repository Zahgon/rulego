package dao

import (
	"examples/server/config"
	"sync"

	"github.com/rulego/rulego/api/types"
)

// IndexKeySpe key 连接符
var IndexKeySpe = ":"

type RuleDao struct {
	config   config.Config
	username string
	index    Index
	sync.RWMutex
}

// Index 定义索引结构，仅包含必要元数据
type Index struct {
	// key=chainId
	Rules map[string]RuleMeta `json:"rules"`
}

type RuleMeta struct {
	Name       string `json:"name"`
	ID         string `json:"id"`
	Root       bool   `json:"root"`
	Disabled   bool   `json:"disabled"`
	UpdateTime string `json:"updateTime"`
}

func NewRuleDao(config config.Config, username string) (*RuleDao, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load or initialize the index

func (d *RuleDao) List(username string, keywords string, root *bool, disabled *bool, size, page int) ([]types.RuleChain, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// 遍历索引中的元数据

// 根据元数据加载完整的规则链数据

// 排序逻辑

func (d *RuleDao) Get(username, chainId string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *RuleDao) GetAsRuleChain(username, chainId string) (types.RuleChain, error) {
	_ = "STUB: not implemented"
	// 根据ID加载规则链DSL数据
	return *new(types.RuleChain), nil
}

func (d *RuleDao) Save(username, chainId string, def []byte) error {
	_ = "STUB: not implemented"
	return nil
}

//创建索引

// 保存索引到文件

func (d *RuleDao) saveRuleChain(username, chainId string, def []byte) error {
	_ = "STUB: not implemented"
	return nil
}

//创建文件夹

//保存到文件

//保存规则链到文件

func (d *RuleDao) Delete(username, chainId string) error { _ = "STUB: not implemented"; return nil }

func (d *RuleDao) getIndexPath() string { _ = "STUB: not implemented"; return "" }

func (d *RuleDao) rebuildIndex() error { _ = "STUB: not implemented"; return nil }

// 构建完整的路径

// 读取目录下的所有文件

// 遍历文件

// 构建文件的完整路径

// 读取文件内容

// 解析 JSON 数据到 RuleChain 结构体

func (d *RuleDao) loadIndex(indexPath string) error { _ = "STUB: not implemented"; return nil }

func (d *RuleDao) createIndex(ruleChain types.RuleChain) { _ = "STUB: not implemented"; return }

// 更新索引

func (d *RuleDao) deleteIndex(chainId string) error { _ = "STUB: not implemented"; return nil }

func (d *RuleDao) saveIndex(indexPath string) error { _ = "STUB: not implemented"; return nil }

func (d *RuleDao) getAllIndex() []RuleMeta { _ = "STUB: not implemented"; return nil }
