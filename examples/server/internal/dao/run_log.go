package dao

import (
	"examples/server/config"
	"path/filepath"

	"github.com/rulego/rulego/api/types"
)

type EventDao struct {
	*FileStorage
	config config.Config
}

func NewEventDao(config config.Config) (*EventDao, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SaveRunLog 保存工作流运行日志快照
func (s *EventDao) SaveRunLog(username string, ctx types.RuleContext, snapshot types.RuleChainRunSnapshot) error {
	_ = "STUB: not implemented"
	return nil
}

//创建文件夹

//保存到文件

//v, _ := json.Format(byteV)
//保存规则链到文件

func (s *EventDao) Delete(username string, chainId, id string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *EventDao) DeleteByChainId(username string, chainId string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *EventDao) visit(files *[]string) filepath.WalkFunc {
	_ = "STUB: not implemented"
	return *new(filepath.WalkFunc)
}

func (s *EventDao) List(username string, chainId string, current, size int) ([]types.RuleChainRunSnapshot, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

//加载所有

// 获取目录下所有运行日志文件

// 按文件时间戳排序

// 计算分页的起始索引

// 遍历文件，每个文件对应一条 RuleChainRunSnapshot 记录

func (s *EventDao) Get(username, chainId, snapshotId string) (types.RuleChainRunSnapshot, error) {
	_ = "STUB: not implemented"
	return *new(types.RuleChainRunSnapshot), nil
}
