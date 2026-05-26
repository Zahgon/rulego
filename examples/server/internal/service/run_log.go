package service

import (
	"examples/server/config"
	"examples/server/internal/dao"

	"github.com/rulego/rulego/api/types"
)

var EventServiceImpl *EventService

type EventService struct {
	EventDao *dao.EventDao
	config   config.Config
}

func NewEventService(config config.Config) (*EventService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SaveRunLog 保存工作流运行日志快照
func (s *EventService) SaveRunLog(username string, ctx types.RuleContext, snapshot types.RuleChainRunSnapshot) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *EventService) Delete(username, chainId, id string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *EventService) DeleteByChainId(username, chainId string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *EventService) List(username, chainId string, current, size int) ([]types.RuleChainRunSnapshot, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (s *EventService) Get(username, chainId, snapshotId string) (types.RuleChainRunSnapshot, error) {
	_ = "STUB: not implemented"
	return *new(types.RuleChainRunSnapshot), nil
}
